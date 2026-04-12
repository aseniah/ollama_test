# Changelog

## v2 (2026-04-09)

### Impact on results

The bug fix prompt rewrite (006) should meaningfully improve scores for that test — v1 runs were largely polluted by models returning the FizzBuzz sequence as plain text rather than code, which is a harness misunderstanding rather than a capability failure. Expect most models to do better here in v2.

The vacuous check fixes (001, 004, 008) will modestly lower scores for models that produce empty output due to compile failures. Previously those models received partial credit for checks that passed vacuously — now they won't. The effect is largest for languages with high compile failure rates (Go, C#).

The C# LANG_NOTE update gives models explicit guidance on CSV and JSON parsing without NuGet. This should reduce compile failures on tests that involve those formats, raising C# scores across the board.

### Test changes

- `006_bug_fix/prompt.md`: removed "output the corrected sequence" phrasing that caused models to return the FizzBuzz sequence as plain text instead of returning code
- `008_prime_numbers/verify.py`: `one_excluded` and `no_composite` now require non-empty output — previously both passed vacuously on empty stdout from compile errors
- `001_csv_to_json/verify.py`: `age_is_int` and `score_is_float` now require non-empty output — same vacuous pass pattern as above
- `004_json_filter/verify.py`: `bob_excluded` and `carol_excluded` now require non-empty output — same vacuous pass pattern as above

### Prompt changes

- `benchmark.py` variant C system prompt: added "Write a {lang_name} program that solves the task described by the user." Tests 001, 002, 004, and 007 use imperative task specs ("Read the file... output...") that weaker models could answer directly rather than generating code. The new line closes that ambiguity for all tests. Applies to both Ollama and Claude runners since both resolve `_VARIANT_TEMPLATES["C"]` from `benchmark.py` at runtime.

### Harness changes

- `benchmark.py` / `run_claude_test.py`: `verify_error` flag in exception handler corrected to `True` (was `False`)
- `benchmark.py` LANG_NOTE C#: added CSV parsing guidance (`File.ReadAllLines()` / `string.Split(',')`); "no NuGet packages" moved earlier in the note
- `run_claude_test.py`: removed `--gen-ms` argument; `ms` is now stored as `null` for all Claude API runs. Generation time is not captured — Claude API latency includes network and subagent overhead, making it incomparable to Ollama inference times. Pass/fail is the primary signal for Claude comparisons.
- `benchmark.py` / `run_claude_test.py`: updated all paths for the new test directory structure (`test/prompt.md`, `test/input/`, `grading/verify.py`)

### Test directory restructure

Each test directory now has `test/` (prompt.md + input/) and `grading/` (verify.py) subdirectories. README.md stays in the test root. This prevents Claude subagents from reading grading artifacts during generation — previously subagents had filesystem access to verify.py, which could influence code generation and invalidate first-shot measurement.

### Claude orchestration

Overhauled to enforce genuine first-shot measurement. The orchestrator now assembles fully-resolved prompts before dispatching — reading `benchmark.py` once, substituting `{language}` and `{source_code}`, and for test 007 embedding both `input.csv` and `expected_format.json` contents directly in the briefing. Subagents receive only the resolved system and user prompts, require no tool access, and return only raw source code. The orchestrator extracts the code and calls `run_claude_test.py` itself. Runs proceed in waves of 4 (one test × all 4 languages simultaneously) to prevent cross-language contamination within a test. Previously, subagents had full filesystem access and ran the verifier themselves, which allowed iteration on failures and invalidated first-shot measurement.

### Input file inlining

`benchmark.py` `build_messages` now appends all data input files (CSV, JSON, TXT — anything not a language source extension) directly into the user prompt, in `--- input/filename ---\n{contents}` format. Previously only the Claude orchestrator did this (for test 007 only), while Ollama models received only the bare `prompt.md` text and had to write code that read the files without knowing their schema upfront. This created an inconsistency: Claude models had a cheat sheet for field names and structure; Ollama models did not. The change brings both paths to parity. Language source files (`.py`, `.ts`, `.go`, `.cs`, `.csx`) are skipped since they are already inlined via `{source_code}` substitution for tests 005 and 006. Tests 003 and 008 have no input files and are unaffected. The Claude orchestrator instructions (CLAUDE.md) are updated to apply the same rule to all tests rather than special-casing 007. All v2 Ollama results should be re-run.

### Codegen worker agent

Added `.claude/agents/codegen-worker.md` — a project-scoped Claude Code sub-agent definition for code generation cells. The agent has no tools, preventing any filesystem access during generation. This closes the remaining validity gap from the v2 orchestration overhaul: general-purpose subagents carried the full Claude Code system prompt and tool availability even when those tools went unused. The `codegen-worker` agent strips that scaffolding so generation context is limited to the resolved system and user prompts, matching Ollama's inference conditions as closely as possible within the Claude Pro (no direct API) constraint.

## v1 (2026-04-05)

Initial release.

- 8 tests across 4 languages: Python, TypeScript, Go, C#
- Prompt variant C: natural style with explicit stdout hint
- Ollama backend with preload/warmup/unload lifecycle
- Apple on-device model support via apfel backend
- Claude API runner (Sonnet, Haiku, Opus) via subagent tool use
- Artifact writing per run: solution, stdout, stderr, result.json
- Results appended to `results/v{NNN}/{model}/results.jsonl` (all backends)
