# AI Model Benchmarks

Two benchmarks for evaluating local AI model backends on a Mac. Both support [Ollama](https://ollama.com/) models and Apple's on-device model via [apfel](https://github.com/Arthur-Ficial/apfel) (macOS 26+).

## Benchmarks

### `shell/` — Shell command generation

Tests models as an [iTerm2](https://iterm2.com/) AI plugin backend. Each model is given natural-language requests and must return raw shell commands — no markdown, no explanation, just the command. Scored for correctness (clean command output, BSD/macOS userland) and latency.

See [`shell/README.md`](shell/README.md) for setup and usage.

### `codegen/` — Code generation

Tests models on programming tasks across Python, TypeScript, Go, and C#. Each test presents a task prompt, the model generates code, the code is executed, and a per-test verifier scores the output. Results capture both pass/fail and partial credit per check.

See [`codegen/README.md`](codegen/README.md) for setup and usage.

## Shared

- `apfel_backend.py` — Apple on-device model backend, shared by both benchmarks
- `.gitignore` — covers both benchmark directories

## Requirements

- Python 3.12+
- Ollama (for local model runs): `ollama serve`
- Language runtimes for codegen: `tsx` (TypeScript), Go, `dotnet-script` (C#)
- apfel + macOS 26+ (optional, for Apple on-device model)
