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

### Runner + language runtimes

**Python 3.12+** — the runner has no external dependencies (stdlib only: `tomllib`, `urllib`, `unittest`).

Language runtimes are required to execute and score generated code:

```sh
# TypeScript (run with tsx)
npm install -g tsx
# Go
brew install go              # or https://go.dev/dl
# C# (.csx scripts via dotnet-script)
brew install dotnet
dotnet tool install -g dotnet-script
```

Python is built in. `bash lint.sh` from `codegen/` checks the runner (ruff + pyright); `python3 -m unittest discover -s tests -t . -p 'test_*.py'` runs the unit tests.

### Ollama harness (default)

```sh
brew install ollama          # CLI + server
ollama serve &               # start the API server on :11434 (or run the menubar app)
ollama pull qwen3.8:27b      # pull each model you want to benchmark
ollama list                  # see what's installed — these tags go in settings.toml
```

Model ids in `settings.toml` are the exact `ollama pull` tags (e.g. `qwen3.8:27b`, `qwen3-coder:30b`).

### LM Studio harness (`--harness lmstudio`)

```sh
brew install --cask lm-studio
~/.lmstudio/bin/lms bootstrap   # put `lms` on PATH (restart your shell after)
lms get "qwen3.8 27b" --mlx     # search + download (or use the app's Discover tab)
lms ls                          # list installed models — the LLM column is the id for settings.toml
```

The benchmark starts (`lms server start`) and stops (`lms server stop`) the LM Studio server itself when `[harness.lmstudio].autostart = true` (the default) — you don't need the app open. Model ids in `settings.toml` are what `lms ls` / `GET http://localhost:1234/v1/models` report (e.g. `qwen3.8-27b-mlx`, `qwen/qwen3-coder-30b`).

### Apple on-device (`--apple`)

```sh
brew install apfel           # https://apfel.franzai.com — OpenAI-compatible server for Apple Intelligence
```

Requires **macOS 26+** with Apple Intelligence enabled. There is no model to download — apfel exposes the built-in foundation model. The benchmark starts and stops apfel on port 11435 automatically.

### Anthropic / Claude

No install. Generation is orchestrated by a [Claude Code](https://claude.ai/code) session (the `codegen-worker` subagent produces the code; `run_claude_test.py` executes and scores it). You need Claude Code and an API key or subscription. Set the model aliases in `settings.toml` under `anthropic_models`.

## Configuration

`settings.toml` controls everything about a run:

- `[harness.ollama].models` / `[harness.lmstudio].models` — a model list **per harness** (Ollama tags and LM Studio ids differ, and the same weights have different quantizations across the two, so results are kept separate). Set `enabled = true` on the ones you want; the rest stay in the file. `think = true/false` picks the mode; `infer_timeout` / `exec_timeout` override `[defaults]` per model.
- `[harness]` — `default` local harness (`ollama` or `lmstudio`); each `[harness.*]` block has a `base_url`. `[harness.lmstudio]` and `[harness.apple]` take `autostart = true/false`.
- `[defaults]` — `infer_timeout`, `exec_timeout`, `max_tokens` (generation cap → Ollama `num_predict` / OpenAI `max_tokens`), `languages`. Any of these can be overridden per model.
- `anthropic_models` — Claude alias → model id, for `run_claude_test.py`.

> **LM Studio verbosity:** LM Studio applies its own generation defaults and chat template, so a model can be several times more verbose there than the same weights under Ollama (e.g. ~2000 tokens vs ~400 for a simple task). The LM Studio model entries use a longer `infer_timeout` for this reason; `max_tokens` caps genuine runaway output.

Each model is warmed up (loaded) before its first test, then unloaded before the next model to free memory (Ollama `keep_alive`, LM Studio `lms unload`, with LM Studio's Auto-Evict as a backstop).

## Running

All commands run from `codegen/`. Results go to `results/v002/{harness}/{model}/results.jsonl`; per-run artifacts (solution, stdout, stderr) go alongside under `{timestamp}/{lang}/{test}/`.

### Ollama

```sh
python3 benchmark.py                    # all enabled ollama models, 1 run
python3 benchmark.py 3                  # 3 runs, aggregated summary at the end
```

Needs `ollama serve` running. Edit `[harness.ollama].models` to choose models.

### LM Studio

```sh
python3 benchmark.py --harness lmstudio
python3 benchmark.py 3 --harness lmstudio
```

Edit `[harness.lmstudio].models`. The server is started/stopped automatically (autostart on); if you set `autostart = false`, run `lms server start` yourself first.

### Apple

```sh
python3 benchmark.py --apple            # runs the enabled ollama models AND the Apple model
python3 benchmark.py --harness lmstudio --apple
```

`--apple` adds the Apple foundation model on top of whichever local harness is selected. apfel is started/stopped automatically.

### Anthropic / Claude

Set `enabled = true` on the aliases you want in `settings.toml` (`anthropic_models`), start a Claude Code session in `codegen/`, and say:

> Run the Claude codegen benchmark for haiku. Follow the orchestration instructions in `CLAUDE.md` exactly.

Claude Code dispatches the 32 generation cells, then calls `run_claude_test.py` per cell. Results land in `results/v002/anthropic/{alias}/results.jsonl`.

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
