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

**Ollama** must be running for local model benchmarks: `ollama serve`

**Apple on-device model** (`--apple` flag) requires [apfel](https://github.com/Arthur-Ficial/apfel) and macOS 26+.

## Running — Ollama models

Edit the `MODELS` list near the top of `benchmark.py` to control which models run (several larger models are commented out). Then:

```sh
python3 benchmark.py              # single run
python3 benchmark.py 3            # 3 runs, results averaged at end
python3 benchmark.py --apple      # include Apple on-device model
```

Results are written to `results/v001/{model}/results.jsonl`.

## Running — Claude API models

This benchmark uses [Claude Code](https://claude.ai/code) as the orchestrator. Edit the `MODELS` dict in `run_claude_test.py` to uncomment the Claude models you want to test, then start a Claude Code session and say:

> Run the Claude codegen benchmark. For each active model in `run_claude_test.py`, dispatch all test × language combinations in parallel. For each combination, spawn a subagent using that Claude model whose **only job is to return the raw source code** — the subagent may read input files to understand the task but must not run `run_claude_test.py` or any verifier. After the subagent returns, extract the raw code from its response (strip any explanation or markdown fences), save it to a temp file, then call `run_claude_test.py` to execute, verify, and record the result. The subagent must never see verifier output — it generates once and returns.

Results are written to `results/v002/{model}/results.jsonl` (e.g. `results/v002/opus/results.jsonl`).

> **Note on timing:** Generation time (`ms`) is not captured for Claude API runs — it is stored as `null` in results. Claude API latency includes network round-trips and subagent overhead, making it incomparable to Ollama on-device inference times. Pass/fail and partial scores are the primary signals for Claude API comparisons.

## Reviewing results

Results are analyzed via targeted `jq` queries against the per-model JSONL files. See `CLAUDE.md` for query examples and failure category definitions.

## Files

| Path | Description |
|---|---|
| `benchmark.py` | Main Ollama test runner |
| `run_claude_test.py` | Claude API test helper — executes and verifies generated code |
| `tests/NNN_name/prompt.md` | Task prompt for each test |
| `tests/NNN_name/verify.py` | Correctness verifier for each test |
| `results/v{NNN}/{model}/results.jsonl` | Per-model benchmark results (Ollama + Claude) |
| `results/v{NNN}/{model}/{timestamp}/` | Per-run artifacts (solution, stdout, stderr) |
| `findings/FINDINGS_v{NNN}.md` | Human-readable analysis and conclusions |
| `findings/findings_instructions.md` | Guide for generating versioned findings reports |
| `CHANGELOG.md` | History of prompt and verifier changes by version |
