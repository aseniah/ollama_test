# iTerm2 Shell Command Benchmark

## Purpose

Benchmark local AI model backends for the iTerm2 AI plugin — everyday shell command generation on macOS. A good backend must:

- Return raw shell commands — no markdown, backticks, or fences
- Use BSD/macOS userland correctly (not GNU flags)
- Respond fast enough for interactive use (~1–2s acceptable)

If a model fails a prompt category due to content guardrails, context limits, or other constraints, that is a meaningful data point and is recorded as-is.

## Structure

- `benchmark.py` — main test runner (Ollama-focused)
- `apfel_backend.py` — Apple on-device model backend
- `results/vNNN.jsonl` — one record per model/variant/prompt/run
- `findings/FINDINGS.md` — human-readable analysis and conclusions
- `findings/APFEL-FINDINGS.md` — Apple model specific findings

## Running

```sh
python3 benchmark.py              # Ollama models only
python3 benchmark.py --apple      # Include Apple on-device model (requires apfel, macOS 26+)
python3 benchmark.py 3 --apple    # 3 runs
```

Ollama must be running. With `--apple`, apfel is started automatically on port 11435 and stopped when the run completes (unless it was already running).

## Prompt Variants

Each model is tested against 3 system prompt variants (A/B/C) and 8 shell tasks. Results are appended to the current version's `.jsonl` file — re-runs accumulate for averaging.

## Lint

```sh
bash lint.sh  # run from shell/
```
