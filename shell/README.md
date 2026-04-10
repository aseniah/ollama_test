# iTerm2 Shell Command Benchmark

Benchmarks local AI models for use as an [iTerm2](https://iterm2.com/) AI plugin backend. The plugin passes natural-language requests to a model and expects back a raw shell command — no markdown, no explanation, just the command.

A good backend must:
- Return a raw shell command with no fences, backticks, or prose
- Use BSD/macOS userland correctly (not GNU coreutils flags)
- Respond fast enough for interactive use (~1–2 seconds is acceptable)

## How it works

The benchmark runs each model through 8 shell tasks (git log, process listing, tar with exclusions, recursive grep, etc.) against 3 system prompt variants (A/B/C). Each combination is scored for correctness (is the output a clean shell command?) and latency (ms). Results are appended to a versioned JSONL file so re-runs accumulate for averaging.

Models are tested via the [Ollama](https://ollama.com/) REST API. Apple's on-device model (`apfel`) can also be included.

## Setup

**Ollama** must be running with at least one model pulled:
```sh
ollama serve
ollama pull qwen2.5-coder:7b
```

**Python 3.12+** is required (no external dependencies — stdlib only).

For Apple on-device model support, `apfel` is required (macOS 26+ only). It is started and stopped automatically by the benchmark when `--apple` is passed.

## Running

```sh
cd shell
python3 benchmark.py              # Ollama models, single run
python3 benchmark.py 3            # 3 runs (results averaged at end)
python3 benchmark.py --apple      # Include Apple on-device model
python3 benchmark.py 3 --apple    # 3 runs with Apple model
```

Results are written to `results/vNNN.jsonl` where `NNN` is the current prompt version. Running multiple times appends to the same file.

## Configuring models

Edit the `MODELS` list near the top of `benchmark.py`. Models are tested sequentially; each is loaded, tested across all prompts and variants, then unloaded before the next begins.

## Output

After each run the benchmark prints:
- A summary table per model (latency + clean-command rate, by prompt variant)
- The latest raw response for each model/variant/task combination

## Files

| Path | Description |
|---|---|
| `benchmark.py` | Main test runner |
| `results/vNNN.jsonl` | One record per model/variant/prompt/run |
| `findings/FINDINGS.md` | Human-readable analysis and conclusions |
| `findings/APFEL-FINDINGS.md` | Apple model specific findings |
