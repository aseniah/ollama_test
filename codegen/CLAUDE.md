# Code Generation Benchmark

## Purpose

Benchmark local and cloud AI model backends for code generation quality across multiple languages.

## Structure

- `settings.toml` — model list, harness config, timeouts, languages. Edit this to change a run.
- `settings.py` — loader/validator for `settings.toml`
- `backends.py` — `OllamaBackend`, `LMStudioBackend`, `AppleBackend`
- `benchmark.py` — main test runner (Ollama / LM Studio / apfel)
- `run_claude_test.py` — Claude API test runner (harness = `anthropic`, via subagents)
- `machine.py` — detects the host (chip/cores/RAM), maps it to a slug via `machines/<slug>.toml`
- `machines/<slug>.toml` — committed registry, one per known benchmark host
- `migrate_results.py` — one-time migration of old `results/v{NNN}/{model}/` dirs into the harness layout
- `migrate_results_machine.py` — one-time migration of `results/v{NNN}/{harness}/` into the machine layout
- `results/FINDINGS_v{NNN}.md` — human-readable analysis and conclusions
- `results/findings_instructions.md` — guide for generating versioned findings reports
- `results/v{NNN}/{machine}/{harness}/{alias}/results.jsonl` — results per machine+harness+model; `NNN` = `PROMPT_VERSION`
- `results/v{NNN}/{machine}/{harness}/{alias}/{timestamp}/` — per-run: `meta.json` (machine snapshot) + `{lang}/{test}/` artifacts (solution, stdout, stderr)

Claude API results are hardware-independent — they live under the `api` pseudo-machine
(`results/v{NNN}/api/anthropic/{alias}/`).

Every model entry has an `alias` (our label — drives the result dir name and the `model`
field) and a `model_id` (what's sent to the harness; defaults to `alias`). For Ollama they're
equal; for LM Studio the alias is runtime-explicit (`qwen3.8-27b-mlx-4bit`) and `model_id` is
LM Studio's own id (`qwen/qwen3.8-27b@4bit`). Records carry both.

Model lists live in `settings.toml`, one per local harness (`[harness.ollama].models`,
`[harness.lmstudio].models`) — Ollama tags and LM Studio model ids differ (and the same
weights carry different quantizations across the two), so the lists are independent and their
results stay separate. Set `enabled = true` on the models you want in a
run; the rest stay in the file. The benchmark goal is evaluating models for local AI coding
assistant use.

Between models the benchmark unloads the previous one to free memory: Ollama via
`keep_alive: 0`, LM Studio via `lms unload` (falls back to LM Studio's JIT auto-evict if the
`lms` CLI is unavailable or the id does not match).

## Running

Model list, harness, and timeouts live in `settings.toml`. Pick a local harness with
`--harness` (default: `[harness].default`).

```sh
python3 benchmark.py                      # enabled models, default harness (ollama)
python3 benchmark.py 3                     # 3 runs
python3 benchmark.py --harness lmstudio    # same models against LM Studio
python3 benchmark.py --apple               # also run the Apple on-device model
```

Ollama must be running. LM Studio and apfel are started automatically when their harness is
selected and `autostart = true` (the default in `settings.toml`) — via `lms server start` and
apfel respectively — and stopped at the end, unless they were already running. Set
`autostart = false` under `[harness.lmstudio]` to require a manually-started server instead.
The `lms` CLI (bundled with LM Studio) must be on PATH or at `~/.lmstudio/bin/lms` for LM
Studio autostart.

### Machine identity

`benchmark.py` resolves the host machine first thing. Results are hardware-bound (token
rates, latency), so they are stored per machine under `results/v{NNN}/{machine}/`. On a known
host (fingerprint matches a `machines/*.toml`) it prints `Machine: <slug>` and continues. On a
new host it prompts for a slug and writes `machines/<slug>.toml`; pass `--machine <slug>` to
skip the prompt (required in non-interactive contexts). `python3 machine.py` prints the
detected specs and resolved slug; `python3 machine.py --register <slug>` registers without a
run. The fingerprint is derived from hardware class only (model id, chip, cores, RAM) so OS
and toolchain updates don't orphan a machine.

## Running Claude Tests (sonnet / haiku / opus)

### Orchestrator steps

1. Generate run_id: `python3 -c "import datetime; print(datetime.datetime.now().strftime('%Y%m%d_%H%M%S'))"` (local time, no UTC)
2. Read `benchmark.py` to extract `LANGUAGES`, `LANG_EXT`, `LANG_NAME`, `LANG_NOTE`, and `_VARIANT_TEMPLATES["C"]`
3. Read `run_claude_test.py` for `PROMPT_VERSION`
4. Discover tests by listing subdirectories of `tests/` (skip `_`-prefixed dirs)
5. For each test, pre-read all files the subagent will need (see "Per-cell prompt assembly" below)
6. Build the full 32-cell queue (8 tests × 4 languages)
7. **Run in waves of 4, grouped by test**: dispatch all 4 languages for one test simultaneously, wait for all 4 to return, then move to the next test. This ensures no subagent for a given test can observe solutions from another language.
8. For each completed subagent: strip any markdown fences from the returned code, write to `/tmp/{test_id}_{language}_solution{ext}`, then run:
   ```
   python3 run_claude_test.py {language} {test_id} /tmp/{test_id}_{language}_solution{ext} {run_id} --model {model}
   ```
   from the `codegen/` directory
9. Collect and display all results when all waves complete

### Per-cell prompt assembly (orchestrator does this, not the subagent)

For each (test_id, language) cell, build two strings to pass to the subagent:

**System prompt:** `_VARIANT_TEMPLATES["C"].format(language=LANG_NAME[language], lang_note=LANG_NOTE[language]).strip()`

**User prompt:** start with `tests/{test_id}/test/prompt.md`, then apply substitutions:
- Always: `{language}` → `LANG_NAME[language]`
- Test 005 only: `{source_code}` → contents of `tests/005_unit_test_writer/test/input/source{ext}` (fall back to `.cs` if no match)
- Test 006 only: `{source_code}` → contents of `tests/006_bug_fix/test/input/buggy{ext}` (fall back to `.cs` if no match)
- All tests with data input files (001, 002, 004, 007): after all substitutions, append each file in `tests/{test_id}/test/input/` that is NOT a language source file (i.e. skip `.py`, `.ts`, `.go`, `.cs`, `.csx`), in sorted order:
  ```
  --- input/{filename} ---
  {file contents}
  ```

### Each subagent's job

Dispatch each cell using `subagent_type: "codegen-worker"`. This agent has no tools, preventing any filesystem access during generation. Do not use the default general-purpose agent — it carries tool access and Claude Code scaffolding that would give the model context beyond the resolved prompts.

The subagent receives the fully-resolved system prompt and user prompt as strings — no file access required. It has one job: **return only raw source code**. No explanation, no markdown fences, no tool calls.

## Analyzing Results

Results are per-machine+harness+model JSONL files at
`results/v{NNN}/{machine}/{harness}/{alias}/results.jsonl`. Every record carries `machine`
(a slug, or `api` for Claude), `harness` (`ollama` / `lmstudio` / `apple` / `anthropic`),
`model` (the alias — group on this), and `model_id` (the raw id sent to the harness). Use
targeted jq queries — don't load all records at once. Note the three-level glob (`*/*/*`):
machine, harness, alias.

```sh
# Pass rate per model across a version (all machines + harnesses)
jq -s 'group_by(.model) | map({model: .[0].model, passed: map(select(.passed))|length, total: length})' results/v002/*/*/*/results.jsonl

# One machine + harness only
jq -s 'group_by(.model) | map({model: .[0].model, passed: map(select(.passed))|length, total: length})' results/v002/m3-max-48gb/ollama/*/results.jsonl

# Same model, ollama vs lmstudio (one machine)
jq -s 'group_by(.harness) | map({harness: .[0].harness, passed: map(select(.passed))|length, total: length})' \
  results/v002/m3-max-48gb/ollama/qwen3.8_27b_nothink/results.jsonl results/v002/m3-max-48gb/lmstudio/qwen3.8_27b_nothink/results.jsonl

# Same model + harness across machines (score should match; tok/s will not)
jq -s 'group_by(.machine) | map({machine: .[0].machine, tok_per_sec: (map(.tok_per_sec)|add/length), passed: map(select(.passed))|length, total: length})' \
  results/v002/*/ollama/qwen3-coder_30b/results.jsonl

# Compare thinking vs non-thinking for a specific model
jq -s 'group_by(.thinking) | map({thinking: .[0].thinking, passed: map(select(.passed))|length, total: length})' \
  results/v002/m3-max-48gb/ollama/qwen3.5_4b_think/results.jsonl results/v002/m3-max-48gb/ollama/qwen3.5_4b_nothink/results.jsonl

# Pass rate by language across all models
jq -s 'group_by(.language) | map({language: .[0].language, passed: map(select(.passed))|length, total: length})' results/v002/*/*/*/results.jsonl

# Mean tok/s per model
jq -s 'group_by(.model + (.thinking|tostring)) | map({model: .[0].model, thinking: .[0].thinking, tok_per_sec: (map(.tok_per_sec) | add / length)})' results/v002/*/*/*/results.jsonl

# Failures for one model with stderr snippet
jq 'select(.passed == false) | {test, language, stderr: .stderr[:120]}' results/v002/api/anthropic/haiku/results.jsonl

# Partial scores (checks passed / total) per model
jq -s 'group_by(.model) | map({model: .[0].model, checks_passed: map(.checks | to_entries | map(select(.value)) | length) | add, checks_total: map(.checks | length) | add})' results/v002/*/*/*/results.jsonl
```

### Failure categories

Derived from `ran`, `exit_code`, and `stderr` fields in each record:

| Category | Condition |
|----------|-----------|
| `PASS` | `passed == true` |
| `LOGIC` | ran, exit 0, checks failed |
| `RUNTIME` | ran, exit ≠ 0, stderr contains `Traceback` / `Error:` / `Exception` |
| `COMPILE` | ran, exit ≠ 0, stderr matches `error CS\d+` / `# command-line-arguments` / `undefined:` / `syntax error` / `Transform failed` |
| `BAD_PKG` | ran, exit ≠ 0, stderr contains `no required module` / `go.mod file not found` |
| `ERR_EXIT` | ran, exit ≠ 0, none of the above |
| `TIMEOUT` | not ran, stderr contains `TIMEOUT` |
| `CTX_OVF` | not ran, stderr contains `context_length_exceeded` or `context window` |
| `NO_RUN` | not ran, none of the above |

### Partial scoring

Each record has a `checks` object (`{check_name: bool}`). Partial score = count of `true` values / total keys. Useful for ranking models that mostly pass vs. those that mostly fail.

## Prompt Variant History

Each JSONL record has a `prompt_variant` field (`"A"`, `"B"`, `"C"`, ...) and a `prompt_v` field
(`PROMPT_VERSION` in `benchmark.py`). These together identify the exact prompt used.

| Version | Variant | Description |
|---------|---------|-------------|
| v1–v2   | A       | Strict — "Return ONLY the raw code. No markdown fences..." |
| v1–v2   | B       | Natural — "You are a helpful {lang} developer." |
| v3+     | C       | Natural + stdout hint — "You are a helpful {lang} developer. Write all output to stdout." |

A/B were retired after v2. B consistently outperformed A (both models, all languages). The
backtick suppression in A was a carry-over from the shell benchmark and hurt code quality by
making models produce overly minimal code.

## Known Failure Patterns

These are model quality issues, not harness bugs — don't re-investigate them:

- **Go COMPILE** (`# command-line-arguments`): Models consistently omit required imports.
  Affects both models across most tests.
- **C# arg parsing**: Models use `args[0]` (undefined) or `Environment.GetCommandLineArgs()[1]`
  (returns the script filename, not the first arg). The correct approach for `.csx` is `Args[0]`
  (built-in `IList<string>`). Hint is in `LANG_NOTE` but models often ignore it.
- **C# NuGet / missing types**: Models reach for library types (`ObjectReader`, `CsvHelper`,
  `Newtonsoft.Json`) that aren't available in dotnet-script without NuGet. LANG_NOTE now
  explicitly lists the correct built-in approaches for CSV and JSON.
- **TypeScript stdin**: Fixed in v3 by adding `process.argv` hint to `LANG_NOTE`. Prior runs
  show TIMEOUT on 008_prime_numbers because models used readline/stdin instead.
- **005_unit_test_writer**: Fails frequently. `has_pass_lines` requires the exact `PASS: <desc>`
  format specified in the prompt. Models often use `PASSED:`, test frameworks, or single-line
  summaries. This is intentional — format-following is part of the test.
- **007_beatles_interview**: The hardest test. Most models fail entirely. Common issues: wrong
  CSV column name assumptions, incorrect age calculations, no output. Expected to remain mostly
  failing for smaller models.
- **Apple CTX_OVF**: Apple on-device model has a 4096-token context window. Long tests (007
  beatles in Go) overflow it when context accumulates across a run. Not fixable without a
  session-reset API in apfel.

## Maintenance

**007_beatles_interview** — when a living Beatle (currently Paul McCartney or Ringo Starr) dies,
three files must be updated:
- `tests/007_beatles_interview/input/input.csv`: fill in the `Died` column for that person
- `tests/007_beatles_interview/input/expected_format.json`: update their `Age` to their age at death
- `tests/007_beatles_interview/verify.py`: update the hardcoded age check for that person to their age at death

## Lint

```sh
bash lint.sh  # run from codegen/
```
