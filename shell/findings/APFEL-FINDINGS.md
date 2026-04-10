# Apple On-Device Model — iTerm2 Viability Report

iTerm2's AI plugin lets you describe what you want in plain English and get back a shell command. It supports any OpenAI-compatible backend, which raises an obvious question: could Apple's on-device language model fill that role? Apple Silicon Macs running macOS 26 Tahoe with Apple Intelligence enabled include a built-in language model accessible via the FoundationModels framework. [apfel](https://github.com/Arthur-Ficial/apfel) wraps it in an OpenAI-compatible HTTP server — no model downloads, no API keys, no GPU configuration. For users who can't or don't want to run a local inference server, it looks like a natural fit.

This report documents whether the Apple on-device model is actually viable for that use case. It was tested against 8 shell command prompts across 3 system prompt variants and 3 runs on an M3 Max (48GB). For context, the same prompts were run against three Ollama-hosted models on the same machine. Apple's model was tested twice: once with default settings and once with apfel's `--permissive` flag, which is documented as relaxing content guardrails.

---

## The Appeal

Before getting to results, it's worth understanding why this is attractive:

- **Zero setup** — no Ollama, no model download, no GPU configuration. Works on any Apple Silicon Mac with Apple Intelligence enabled and macOS 26 Tahoe.
- **Broad device coverage** — works on any Apple Silicon Mac with Apple Intelligence enabled, including 8GB machines where running a dedicated Ollama model isn't practical (8GB MacBook Air, MacBook Neo — which has no higher-memory option).
- **Fully private** — no API keys, no network calls, no cost.

The hope was that these advantages could make it a practical default or fallback for users who cannot run Ollama.

---

## Verdict

**Not recommended** as an iTerm2 backend. Content guardrails block legitimate shell commands and cannot be worked around with the `--permissive` flag. Command accuracy is poor independent of guardrail issues. Speed is below dedicated coder models despite shorter outputs. No prompt variant resolves all issues simultaneously.

The apple model handles simple, well-known commands adequately when it doesn't hit a guardrail. It may be worth considering as a last-resort fallback for users who cannot run any local model — but it should not be the primary recommendation.

---

## What Was Tested

### Test prompts

Eight common shell tasks were used, chosen to cover a range of complexity and to expose known problem areas (BSD vs GNU userland differences, process monitoring, file operations):

1. Show the last 10 git commits with hash and message, one per line
2. Show the top 5 processes by memory usage
3. Create a tar.gz of the current directory, excluding `.git` and `node_modules`
4. Search recursively for the string TODO in all `.js` files
5. Show disk usage of each subdirectory, sorted by size
6. Find all executable files in the directory tree
7. Replace all occurrences of `foo` with `bar` in-place in `file.txt`
8. Print the size in bytes of `data.txt`

### System prompt variants

Three system prompt variants were tested per model:

- **Variant A** — Full prompt with explicit BSD/macOS flag corrections and worked examples
- **Variant B** — Compressed version of A, fewer words, same key instructions
- **Variant C** — Minimal prompt, no explicit flag guidance

All variants instruct the model to return only a raw shell command — no markdown, backticks, or fences.

### Pass criteria

A result is counted as clean if the response is a raw shell command with no markdown formatting. Guardrail blocks count as failures. Command correctness is evaluated separately from formatting.

---

## Apple Model Results

Tests run on M3 Max (48GB), 3 runs per variant. 72 total results (8 prompts × 3 variants × 3 runs).

| Variant | Clean output | Guardrail blocks | Notes |
|---|---|---|---|
| A (full prompt) | 15/24 | 4 (top 5 by memory 3/3, dir sizes 1/3) | Formatting failures on grep and find |
| B (compressed) | 20/24 | 3 (top 5 by memory 3/3) | Formatting failure on find |
| C (minimal) | 23/24 | 0 | Guardrails avoided; answers are wrong |

**Overall: 58/72 clean.**

---

## Issues

### Content guardrails

Apple's on-device content filtering blocked legitimate shell commands across both test runs (with and without `--permissive`):

- **"Show the top 5 processes by memory usage"** was blocked in every Variant A and B run across both test sessions. The trigger appears to be process monitoring combined with a system prompt that includes sorting examples. Variant C avoids the block but returns wrong answers (`top -bn1 | tail -n5`, which uses GNU flags and doesn't work on macOS).
- **"Show disk usage of each subdirectory"** was blocked intermittently in Variant A during the `--permissive` run.

apfel documents a `--permissive` flag intended to relax guardrails. Testing with this flag produced no improvement — Variant B's block rate increased from 2/3 to 3/3, and Variant A acquired an additional intermittent block. The overall clean score was unchanged at 58/72.

Process monitoring is a routine sysadmin task. A backend that consistently refuses it is unsuitable for general terminal use.

### Command accuracy

The model produces wrong commands for several prompts regardless of variant. These are not formatting issues — the commands themselves are incorrect:

| Prompt | Expected | Apple model returns | Problem |
|---|---|---|---|
| Find executables | `find . -type f -perm +111` | `find . -exec chmod +x {} \;` | Confuses "find executable" with "make executable" |
| Print file size in bytes | `stat -f %z data.txt` | `ls -l data.txt \| wc -l` | Counts lines (always 1), not bytes |
| Sed in-place | `sed -i '' 's/foo/bar/g' file.txt` | `sed -i 's/foo/bar/g' file.txt` | Missing `''` — GNU syntax, fails on BSD |
| Tar with exclusions | `tar -czf out.tar.gz --exclude=.git --exclude=node_modules .` | Various broken pipelines | Systematically wrong flags |
| Dir sizes sorted | `du -sh * \| sort -h` | Wrong tool or flags entirely | Inconsistent across runs |

The find-executables confusion is consistent across all 3 variants and all 3 runs — it is a model capability problem, not a prompting problem.

### Output formatting

Variants A and B return markdown-fenced responses despite explicit instructions to return raw commands. This is most consistent on `find` and `grep` prompts. Variant C avoids this, but as noted above, its answers are often wrong.

No variant simultaneously avoids guardrails, returns correct commands, and returns clean output.

### Speed

The apple model averages ~522ms at ~15 tok/s. Its responses are short (4–12 tokens), so wall-clock times look closer to Ollama models than they are. Longer or more complex commands widen the gap considerably.

| Model | Avg response | tok/s |
|---|---|---|
| apple-foundationmodel | ~522ms | ~15 |
| qwen2.5-coder:7b | ~303ms | ~78 |
| qwen3-coder:30b | ~315ms | ~86 |

---

## Model Comparison

All models tested on M3 Max (48GB) with the same 8-prompt test set, 3 variants, 3 runs each.

| Model | Avg response | tok/s | Clean | Notes |
|---|---|---|---|---|
| `qwen3.5:4b-nvfp4` | ~242ms | ~86 | 72/72 | Fastest; `</think>` token leakage in practice |
| `qwen2.5-coder:7b` | ~303ms | ~78 | 72/72 | Reliable; strong BSD accuracy; 4.7GB footprint |
| `qwen3-coder:30b` | ~315ms | ~86 | 72/72 | Best accuracy; requires ~18–20GB RAM |
| `apple-foundationmodel` | ~522ms | ~15 | 58/72 | `--permissive` tested; guardrails unchanged |
| `qwen3.5:4b` | ~674ms | ~51 | 72/72 | `</think>` leakage; 2× slower than top models |

`qwen2.5-coder:7b` is the recommended default: 4.7GB footprint, instant load, reliable BSD accuracy, and 303ms average response. `qwen3-coder:30b` has marginally better accuracy but requires ~18–20GB RAM and a slow cold load on first use.
