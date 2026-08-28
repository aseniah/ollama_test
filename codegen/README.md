# Code Generation Benchmark

Benchmarks local and cloud AI models for code generation quality across multiple languages. The goal is to evaluate which models are best suited for AI coding assistant use (VS Code, OpenCode, etc.).

Each test presents a model with a programming task. The model generates code, the code is executed, and a per-test verifier checks the output for correctness. Results capture both generation metrics (latency, tokens) and execution outcomes (pass/fail with failure categorization).

## Tests

| # | Test | What it checks |
|---|------|----------------|
| 001 | CSV to JSON | File I/O, data transformation |
| 002 | Word frequency | String processing, sorting, stdout formatting |
| 003 | Fibonacci | Algorithm, CLI arg parsing |
| 004 | JSON filter | JSON parsing, conditional logic |
| 005 | Unit test writer | Format-following, test structure |
| 006 | Bug fix | Code reading, debugging |
| 007 | Beatles interview | CSV parsing, date math, edge cases (hardest) |
| 008 | Prime numbers | Algorithm, CLI arg parsing |

Each test runs against all four languages: **Python, TypeScript, Go, C#**.

## Setup

**Python 3.12+** is required (no external dependencies for the runner itself).

**Language runtimes** are required to execute and score generated code:
- Python: built-in
- TypeScript: `npm install -g tsx`
- Go: [go.dev/dl](https://go.dev/dl/)
- C#: `dotnet tool install -g dotnet-script`

**A local harness** must be running for local model benchmarks:
- **Ollama** (default): `ollama serve`
- **LM Studio** (`--harness lmstudio`): with `autostart = true` (default) the benchmark runs `lms server start` itself and `lms server stop` at the end. The `lms` CLI (bundled with LM Studio) must be on PATH or at `~/.lmstudio/bin/lms`. The models you list must already be downloaded in LM Studio.

**Apple on-device model** (`--apple` flag) requires [apfel](https://github.com/Arthur-Ficial/apfel) and macOS 26+.

## Configuration

`settings.toml` controls everything about a run:

- `[harness.ollama].models` / `[harness.lmstudio].models` — a model list **per harness** (Ollama tags and LM Studio model ids differ, so results stay separate). Set `enabled = true` on the ones you want; the rest stay in the file. `think = true/false` picks the mode; `infer_timeout` / `exec_timeout` override `[defaults]` per model. LM Studio ids are whatever `GET http://localhost:1234/v1/models` reports (usually the lowercased key, e.g. `qwen3.8-27b-mlx-4bit`).
- `[harness]` — `default` local harness (`ollama` or `lmstudio`); each `[harness.*]` block has a `base_url`. `[harness.lmstudio]` and `[harness.apple]` take `autostart = true/false`.
- `[defaults]` — `infer_timeout`, `exec_timeout`, `languages`.
- `anthropic_models = [...]` — Claude aliases → model ids for `run_claude_test.py`.

Between models the previous one is unloaded to free memory (Ollama `keep_alive: 0`, LM Studio `lms unload`). For LM Studio also enable **Auto-Evict** (or a short JIT TTL) in its settings as a backstop.

## Running — local models

```sh
python3 benchmark.py                      # enabled models, default harness (ollama)
python3 benchmark.py 3                     # 3 runs, results averaged at end
python3 benchmark.py --harness lmstudio    # same models against LM Studio
python3 benchmark.py --apple               # also run the Apple on-device model
```

Results are written to `results/v002/{harness}/{model}/results.jsonl`.

## Running — Claude API models

This benchmark uses [Claude Code](https://claude.ai/code) as the orchestrator. Set `enabled = true` on the Claude aliases in `settings.toml` (`anthropic_models`), then start a Claude Code session and say:

> Run the Claude codegen benchmark for [sonnet/haiku/opus]. Follow the orchestration instructions in `CLAUDE.md` exactly.

Results are written to `results/v002/anthropic/{alias}/results.jsonl` (e.g. `results/v002/anthropic/opus/results.jsonl`).

## Latest findings

See [`results/FINDINGS_v002.md`](results/FINDINGS_v002.md) for the full v2 analysis and model rankings.

> **Note on timing:** Generation time (`ms`) is not captured for Claude API runs — it is stored as `null` in results. Claude API latency includes network round-trips and subagent overhead, making it incomparable to Ollama on-device inference times. Pass/fail and partial scores are the primary signals for Claude API comparisons.

## Reviewing results

Results are analyzed via targeted `jq` queries against the per-harness+model JSONL files (note the two-level `results/v{NNN}/*/*/results.jsonl` glob). See `CLAUDE.md` for query examples and failure category definitions.

## Files

| Path | Description |
|---|---|
| `settings.toml` | Model list, harness config, timeouts, languages |
| `settings.py` | Loader/validator for `settings.toml` |
| `backends.py` | Ollama / LM Studio / Apple inference backends |
| `benchmark.py` | Main local test runner |
| `run_claude_test.py` | Claude API test helper — executes and verifies generated code |
| `migrate_results.py` | One-time migration of old result dirs into the harness layout |
| `tests/NNN_name/test/prompt.md` | Task prompt for each test |
| `tests/NNN_name/test/input/` | Input files for tests that need them |
| `tests/NNN_name/grading/verify.py` | Correctness verifier for each test |
| `tests/test_*.py` | Unit tests (`python3 -m unittest discover -s tests -t . -p 'test_*.py'`) |
| `results/v{NNN}/{harness}/{model}/results.jsonl` | Per-harness+model benchmark results |
| `results/v{NNN}/{harness}/{model}/{timestamp}/` | Per-run artifacts (solution, stdout, stderr) |
| `results/FINDINGS_v{NNN}.md` | Human-readable analysis and conclusions |
| `results/findings_instructions.md` | Guide for generating versioned findings reports |
| `CHANGELOG.md` | History of prompt and verifier changes by version |
