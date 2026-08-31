# Code Generation Benchmark — Findings

**Date:** 2026-04-12 to 2026-04-19  
**Machine:** MacBook Pro M3 Max (48GB unified memory)  
**Scope:** 26 model configurations on the Ollama harness, 8 tests × 4 languages, 3 runs each (Claude models: 1 run each). `qwen3-coder:30b` and `qwen3.8:27b` also ran on LM Studio — `qwen3.8:27b` across GGUF, MLX 4-bit, and MLX 8-bit, nothink and think, most 3–5 runs — see Harness Comparison. All runs are on one machine (`m3-max-48gb`, MacBook Pro / M3 Max / 48 GB).

---

## Introduction

This benchmark evaluates how well various AI models generate correct, executable code across four languages (Python, TypeScript, Go, C#) and eight tasks ranging from simple algorithms to complex CSV transformation. The goal is to understand which local Ollama models are viable coding assistant alternatives to cloud-hosted Claude models, and whether tradeoffs like quantization or extended thinking modes are worth it in practice.

Each model generates a solution in response to a task prompt. The solution is executed, and a per-test verifier checks correctness. Results capture both a binary pass/fail and a partial-credit check score.

---

## TL;DR

**Claude models hit the ceiling:** Opus, Sonnet, and Haiku all achieved 100% score and pass rate — perfect first-shot performance across all 32 test/language cells. All three variants are in a class of their own.

**`qwen3-coder:30b` is the top local model — score 94 at ~4s/task.** It passes every test in every language except 007 (Beatles interview), at the fastest inference speed in the benchmark and 100% run-to-run consistency. The MoE architecture (30B total, ~3.3B active) is what makes top-tier quality and interactive latency coexist.

**`gemma4:31b` think (score 93) owns the hardest test:** 11 out of 12 full passes on 007 — the only local configuration that approaches Claude's 4/4. Its overall score trails qwen3-coder slightly, but on multi-step reasoning it is far ahead of every other local model.

**`qwen3.8:27b` is the strongest dense Qwen — 91 nothink / 90 think.** Nothink ties `gemma4:31b` nothink for the best non-thinking local score. Think mode posts 7/12 on 007 — the second-most full passes of any local configuration, behind only `gemma4:31b` think.

**Thinking is a per-model question, not a global one.** The score delta ranges from +20 (`qwen3.5:4b`) to −1 (`qwen3.8:27b`). For `qwen3.8:27b` and both `qwen3.6` models, nothink is the right default; thinking is a targeted tool for 007-class tasks, not a general upgrade.

**`qwen3.6:35b` nothink (score 89) at ~11s** is the best sub-15s dense option — 4 points below qwen3.8:27b nothink but at less than half the time.

**nvfp4 quantization is lossless at 27B nothink**: ±0 points for 32% faster generation. At 27B think the speed gain nearly vanishes (−2 points, 5% faster); at 4B think it costs 13 points for 58% faster.

**On LM Studio, run Qwen in think mode.** The `<think></think>` nothink workaround (needed for LM Studio's MLX thinking bug) costs ~4 points — GGUF nothink goes 84 → 88 with it removed, and MLX nothink can't remove it, so it sits at 82. MLX **think** mode sidesteps the whole problem: **94 / 93% pass**, the benchmark's best `qwen3.8:27b`. A small ~3-point LM-Studio-vs-Ollama residual remains at matched quant (88 vs 91). `qwen3-coder:30b` takes a genuine ~6-point MLX-4bit hit — use GGUF for it.

**Apple's FoundationModel scores 39 / 34% pass rate** — low, and weak on everything except Python, but it runs with essentially no thermal signature: no fan noise, no visible GPU load, unlike every Ollama model.

---

## Overall Rankings

Each thinking/non-thinking variant is listed separately since they represent meaningfully different configurations. Claude models don't have a thinking field in results (they use standard API mode).

Two metrics are reported for each model:

- **Score** — a nuanced grade of overall capability. Each of the 32 test/language cells is scored as the fraction of its correctness assertions that passed (e.g. a solution that passes 3 of 4 checks contributes 0.75 for that cell), then averaged equally across all cells. Every test counts equally regardless of how many assertions it happens to have — a near-miss on a hard task and a clean pass on an easy one both move the score by the same maximum amount.
- **Pass rate** — binary reliability. What fraction of tests did the solution fully pass end-to-end? This answers the practical question: how often will it actually work?

When score exceeds pass rate, the model is earning partial credit on failures — near-misses that clear some checks. When pass rate exceeds score, failures are clean but a few weak spots pull the average down. `qwen3-coder:30b` (score 94 / pass 88%) is the classic first case: it passes everything cleanly except 007, and on 007 it still clears roughly half the checks per cell, so the score sits well above the pass rate.

| Rank | Model | Mode | Score | Pass Rate | Avg Time/task |
|------|-------|------|-------|-----------|--------------|
| 1 | `claude-haiku-4-5` | — | **100** | 100% (32/32) | — |
| 1 | `claude-opus-4-6` | — | **100** | 100% (32/32) | — |
| 1 | `claude-sonnet-4-6` | — | **100** | 100% (32/32) | — |
| 4 | `qwen3-coder:30b` | — | **94** | 88% (84/96) | ~4s |
| 5 | `gemma4:31b` | think | **93** | 93% (89/96) | ~79s |
| 6 | `gemma4:31b` | nothink | **91** | 85% (82/96) | ~24s |
| 6 | `qwen3.5:27b` | think | **91** | 84% (81/96) | ~109s |
| 6 | `qwen3.8:27b` | nothink | **91** | 84% (81/96) | ~23s |
| 9 | `qwen3.6:35b` | think | **90** | 82% (79/96) | ~89s |
| 9 | `qwen3.8:27b` | think | **90** | 90% (86/96) | ~54s |
| 11 | `qwen3.5:27b-nvfp4` | think | **89** | 86% (83/96) | ~103s |
| 11 | `qwen3.6:35b` | nothink | **89** | 81% (78/96) | ~11s |
| 13 | `gemma4:26b` | think | **88** | 88% (84/96) | ~43s |
| 13 | `qwen3.5:35b-a3b-coding-nvfp4` | think | **88** | 83% (80/96) | ~13s |
| 15 | `qwen3.6:35b-a3b-coding-nvfp4` | think | **87** | 78% (75/96) | ~43s |
| 16 | `gemma4:26b` | nothink | **86** | 79% (76/96) | ~9s |
| 16 | `qwen3.5:27b` | nothink | **86** | 80% (77/96) | ~31s |
| 16 | `qwen3.5:27b-nvfp4` | nothink | **86** | 80% (77/96) | ~21s |
| 19 | `qwen3.6:35b-a3b-coding-nvfp4` | nothink | **85** | 80% (77/96) | ~5s |
| 20 | `qwen3.5:35b-a3b-coding-nvfp4` | nothink | **83** | 79% (76/96) | ~5s |
| 21 | `qwen3.5:4b` | think | **68** | 66% (63/96) | ~45s |
| 22 | `qwen2.5-coder:7b` | — | **59** | 53% (51/96) | ~5s |
| 23 | `qwen3.5:4b-nvfp4` | think | **55** | 49% (47/96) | ~19s |
| 24 | `qwen3.5:4b` | nothink | **48** | 45% (43/96) | ~8s |
| 25 | `qwen3.5:4b-nvfp4` | nothink | **44** | 41% (39/96) | ~8s |
| 26 | `apple-foundationmodel` | — | **39** | 34% (33/96) | ~7s |

> **Avg time** is average wall clock time per test/language combination from generation start to result, over cells that actually produced output (timeouts excluded from the average but counted as failures elsewhere). Codegen tasks require 100–500+ output tokens; thinking-mode tasks often exceed 1000 tokens. Claude timing was not instrumented; times will include API round-trip latency and are not comparable to Ollama on-device measurements.

---

## Model Architectures

**Mixture of Experts (MoE):** Rather than activating all parameters for every token, MoE models route each token through a small subset of expert sub-networks. This delivers large-model knowledge at small-model inference cost.

- **`qwen3-coder:30b`** — MoE: 30B total parameters, ~3.3B active per token. This directly explains its combination of the highest local score in the benchmark and ~4s average inference time.
- **`qwen3.5:35b-a3b-coding-nvfp4`** — MoE: 35B total parameters, ~3B active per token ("a3b" = active 3B). Despite having more total parameters than qwen3-coder, it runs at the same ~5s nothink speed for the same reason — only ~3B parameters fire per token. Weight loading still requires memory proportional to total parameters.
- **`qwen3.6:35b-a3b-coding-nvfp4`** — MoE: same architecture as qwen3.5:35b-a3b (35B total, ~3B active, nvfp4 quantized, code-specialized). The 3.6 generation brings modest score improvement (+2 nothink) at the same ~5s inference speed.

**Code-specialized models:**

- **`qwen2.5-coder:7b`** — Fine-tuned specifically on code corpora. Despite specialization, score 59 / 53% pass rate is the weakest of the larger Ollama models, showing that code-specific training at 7B doesn't compensate for the parameter ceiling on multi-step codegen tasks.
- **`qwen3-coder:30b`** — Also code-specialized, at 30B MoE scale. The combination works: it's the fastest competitive local model and the top-scoring one.
- **`qwen3.5:35b-a3b-coding-nvfp4`** / **`qwen3.6:35b-a3b-coding-nvfp4`** — Code-specialized MoE variants. The nvfp4 quantization compresses the 35B weights for efficient inference on Apple Silicon.

**General-purpose models evaluated for coding:**

- **`gemma4:26b`** / **`gemma4:31b`** — Google's Gemma 4 family, not code-specialized, yet strong and exceptionally language-balanced. The 31B variant is a dense model (not MoE) that runs slower than the 26B but at significantly higher quality — particularly on 007-level reasoning tasks.
- **`qwen3.5` family** — General-purpose Qwen 3.5 generation with optional thinking mode.
- **`qwen3.6:35b`** — Dense 35B model (not MoE) from the Qwen 3.6 generation. General-purpose with optional thinking. Despite being larger than qwen3.5:27b, it runs faster (~11s vs ~31s nothink), reflecting architectural and inference efficiency gains in the 3.6 generation.
- **`qwen3.8:27b`** — Dense 27.3B model from the Qwen 3.8 generation (internal architecture family `qwen35`), with a 256K context window and a small (~460M) vision encoder attached — it is nominally multimodal, though only its code generation is exercised here. Not MoE: it runs at ~18 tok/s, comparable to `gemma4:31b` dense. The generational jump over `qwen3.5:27b` is large — +5 score nothink at similar latency, and the highest-scoring dense Qwen in the benchmark.

---

## Thinking vs. Non-Thinking

| Model | Nothink | Think | Score delta | Time cost |
|-------|---------|-------|-------------|-----------|
| `qwen3.5:4b` | 48 | 68 | **+20** | ~8s → ~45s (5.8×) |
| `qwen3.5:4b-nvfp4` | 44 | 55 | **+11** | ~8s → ~19s (2.4×) |
| `qwen3.5:27b` | 86 | 91 | **+5** | ~31s → ~109s (3.6×) |
| `qwen3.5:35b-a3b-coding-nvfp4` | 83 | 88 | **+5** | ~5s → ~13s (2.7×) |
| `qwen3.5:27b-nvfp4` | 86 | 89 | **+3** | ~21s → ~103s (4.9×) |
| `gemma4:31b` | 91 | 93 | **+2** | ~24s → ~79s (3.2×) |
| `gemma4:26b` | 86 | 88 | **+2** | ~9s → ~43s (4.7×) |
| `qwen3.6:35b-a3b-coding-nvfp4` | 85 | 87 | **+2** | ~5s → ~43s (9.2×) |
| `qwen3.6:35b` | 89 | 90 | **+1** | ~11s → ~89s (8.4×) |
| `qwen3.8:27b` | 91 | 90 | **−1** | ~23s → ~54s (2.4×) |

**Key observations:**

- For `qwen3.5:4b`, thinking mode gains 20 points — a tier shift. At 4B parameters, the extra reasoning pass overcomes a significant portion of the base model's limitations.
- The gain compresses as model size increases within a generation. In the qwen3.5 dense line, 4B gains +20 and 27B gains +5.
- `gemma4` shows a modest +2 at both sizes. For the 26B that makes nothink the clear default. For the 31B the score delta is also small, but the think mode's 007 performance (11/12 passes vs 5/12 nothink) makes it compelling for complex tasks regardless.
- **`qwen3.8:27b` is the only configuration where thinking nets negative** (−1). The two effects are real and opposite: on 007 specifically, thinking takes it from 0 to 7 full passes; on the other seven tests it regresses by roughly 2 passes. The aggregate lands just below nothink. Thinking here is a targeted instrument for 007-class multi-step tasks, not a general quality lever.
- **`qwen3.6` models show the worst thinking ROI in the benchmark.** `qwen3.6:35b` gains +1 at 8.4× the time; `qwen3.6:35b-a3b-coding-nvfp4` gains +2 at 9.2×. Both are largely "thinking-saturated" in nothink mode — extended reasoning adds little the base inference doesn't already produce.
- `qwen3.5:35b-a3b-coding-nvfp4` is the most time-efficient thinking upgrade: +5 points for 2.7× cost starting from an already fast ~5s baseline. The resulting ~13s think latency is fully interactive.

**Recommendation:** Use thinking for `qwen3.5:4b` — the quality gain justifies the wait. Skip thinking for `gemma4:26b`, both `qwen3.6` models, and `qwen3.8:27b` in general use — but reach for `qwen3.8:27b` think specifically when a task is 007-class (multi-step schema inference plus date arithmetic). For `gemma4:31b`, prefer think mode if 007-level tasks appear in your workload; otherwise nothink is fine. For `qwen3.5:35b-a3b`, think mode is cheap enough (~13s) to use routinely. For `qwen3.5:27b`, thinking is practical only for batch or background work.

---

## Quantization: nvfp4 vs Standard

The nvfp4 format is an **MLX-native quantization** for Apple Silicon's unified memory architecture. Unlike GGUF (designed for CPU/GPU split inference), MLX models operate natively in Apple's unified memory space.

| Model | Standard | nvfp4 | Score delta | Speed delta |
|-------|----------|-------|-------------|-------------|
| `qwen3.5:27b` nothink | 86 | 86 | **±0** | ~31s → ~21s (32% faster) |
| `qwen3.5:27b` think | 91 | 89 | **−2** | ~109s → ~103s (5% faster) |
| `qwen3.5:4b` nothink | 48 | 44 | **−4** | ~8s → ~8s (~same) |
| `qwen3.5:4b` think | 68 | 55 | **−13** | ~45s → ~19s (58% faster) |

**At 27B nothink**, nvfp4 is a free win: no measurable score cost for 32% faster generation and lower memory. Use it whenever memory is a consideration.

**At 27B think**, the speed advantage nearly vanishes (5% faster) while costing 2 score points. The thinking-mode token volume is high enough that the quantization speed benefit gets absorbed. Standard weights are preferable when using thinking mode.

**At 4B nothink**, nvfp4 costs 4 points at essentially the same speed — not compelling.

**At 4B think**, nvfp4 is 58% faster with a 13-point score drop. That is a steep exchange — at a standard-weights base of 68, the quantized 55 falls back into "weak" territory.

---

## Harness Comparison: Ollama vs LM Studio

`qwen3-coder:30b` and `qwen3.8:27b` both ran on Ollama (GGUF Q4_K_M) and LM Studio (MLX 4-bit, MLX 8-bit, GGUF Q4_K_M) to separate harness effect from quantization effect. Identical decoding parameters (`temperature 1`, `top_k 20`, `top_p 0.95`, `min_p 0`, `repeat_penalty 1`) were sent to every run. The LM Studio rows below are multi-run (an overnight batch), so the numbers are no longer n=1 noise.

### `qwen3.8:27b`

| Harness | Quant | Mode | nothink hack | Runs | Score | Pass | tok/s |
|---|---|---|---|---|---|---|---|
| Ollama | GGUF Q4_K_M | nothink | — (native) | 3 | 91 | 84% | ~18 |
| Ollama | GGUF Q4_K_M | think | — | 3 | 90 | 89% | ~17 |
| LM Studio | GGUF Q4_K_M | nothink | **no** | 3 | 88 | 83% | ~20 |
| LM Studio | GGUF Q4_K_M | nothink | yes | 4 | 84 | 78% | ~20 |
| LM Studio | MLX 4-bit | nothink | yes (forced) | 5 | 82 | 78% | ~23 |
| LM Studio | MLX 8-bit | nothink | yes (forced) | 1 | 81 | 78% | ~13 |
| LM Studio | MLX 4-bit | **think** | — | 3 | **94** | **93%** | ~23 |

### The nothink workaround costs ~4 points; a smaller harness residual remains

Dropping the `<think></think>` prefix hack from the GGUF run (it honors `think=false` natively and never needed it) moved the score **84 → 88** and pass rate **78% → 83%**, with 0 reasoning tokens in every cell. So the workaround does hurt — about 4 points — confirming the direction, though less than the ~8 the earlier n=1 rows suggested.

The hack exists because LM Studio's MLX engine ignores `enable_thinking: false` for Qwen 3.5/3.6/3.8 (bug tracker #1559 / #1870 / #1933, unfixed on 0.4.22 / mlx-llm 1.11.0): left alone, a `think=false` MLX run reasons ~3,200 tokens/cell, blows the token cap, and truncates. Appending `<think>\n\n</think>\n\n` as a trailing assistant message makes LM Studio treat it as a response prefix and the model resumes after an empty block — but prefilling that block and forcing continuation from it bloats and degrades the output (hacked output runs ~40% longer: eval_count ~556 vs ~401).

**MLX nothink cannot drop the hack** — the bug is unavoidable there — so it eats the full ~4-point tax and lands at 82.

A **~3-point residual** still separates LM Studio GGUF no-hack (88) from Ollama GGUF (91) at the same quant. Per-run the ranges barely overlap (LM Studio 89 / 89 / 86, Ollama 93 / 89 / 89), so it is probably a small real difference — LM Studio's community Q4_K_M build vs Ollama's, or LM Studio's chat-template handling — sitting near the benchmark's noise floor. It is not the harness "getting the model wrong"; it is a minor fidelity gap, an order of magnitude smaller than picking the wrong mode.

### Thinking is the answer for Qwen on LM Studio

MLX 4-bit **think** mode sidesteps both the bug and the hack — real reasoning, no prefill — and scores **94 / 93% pass** at ~23 tok/s: the best `qwen3.8:27b` result anywhere in the benchmark, above Ollama's own think (90) and nothink (91). The `max_tokens = 8192` bump for the reasoning volume held — no truncation — and the run-to-run spread (94 / 91 / 97) is far tighter than the nothink runs (86–89 no-hack, 76–88 hacked).

The model's real level is ~88–94 across every harness and quant; only hacked / bug-forced nothink underperforms. If you run Qwen 3.x on LM Studio, **use think mode**.

### `qwen3-coder:30b` — a real MLX quant cost

| Metric | Ollama GGUF (3 runs) | LM Studio MLX 4-bit (5 runs) |
|---|---|---|
| Score | 94 | 88 |
| Pass rate | 87% (84/96) | 83% (134/160) |
| Avg time/task | ~4.0s | ~3.6s |
| tok/s | ~81 | ~96 |

qwen3-coder has no `<think>` mechanism (`nothink_prefix = ""`, no hack), so this 6-point gap is not the workaround — it is a genuine MLX 4-bit quantization cost for this model, unlike qwen3.8 which reaches parity in think mode. Use **GGUF for qwen3-coder**; MLX buys ~18% faster tokens (~96 vs ~81) at a real quality cost.

### MLX 8-bit — still a poor trade

n=1 (disabled after this batch), score 81, and two of its cell losses were artifacts (a 001 TS load-stall timeout, a 004 Python `<tool_call>` leak). The real cost is speed: **~13 tok/s, half of 4-bit**, because 8-bit doubles memory-bandwidth traffic and the M3 Max is bandwidth-bound. The ~29 GB footprint (vs ~15 GB) also raises load-stall risk. No measurable quality upside over 4-bit.

### No overnight degradation

The ~5-hour unattended batch showed zero drift: tok/s held to ±2% across every run of every model, zero timeouts. LM Studio's autostart / per-model unload cycle is stable across a long multi-model session.

### Takeaway

- **The nothink workaround costs ~4 points** (GGUF 84 → 88 with it removed). MLX nothink can't remove it — the thinking bug forces it — so MLX nothink stays depressed at 82.
- **A ~3-point LM-Studio-vs-Ollama residual** remains at the same quant, no-hack (88 vs 91) — likely the GGUF build or chat template, near the noise floor. The harness is close to neutral, not exactly.
- **For Qwen 3.x on LM Studio, run think mode** — `qwen3.8:27b` MLX hits 94 / 93%, sidesteps bug and hack, best in the benchmark.
- **qwen3-coder:30b: use GGUF.** MLX 4-bit costs a real ~6 points and think mode isn't an option (no reasoning path).
- **MLX 8-bit: skip it** on bandwidth-bound Apple Silicon.

---

## Language Breakdown

Pass rates aggregated across all models and runs.

| Language | Pass Rate | Notes |
|----------|-----------|-------|
| Python | **85%** | Highest overall; most models are well-calibrated here |
| TypeScript | **81%** | Strong, with more variance across models |
| Go | **72%** | Compiler strictness reliably surfaces missing or unused imports |
| C# | **64%** | `.csx` scripting environment requires specific idioms |

Go and C# consistently underperform Python. Go failures cluster on compile errors from missing or unused imports — Go's strict compiler rejects code that Python or TypeScript would run. C# issues are primarily the `.csx` script argument access pattern: models default to `args[0]` or `Environment.GetCommandLineArgs()[1]`, both wrong; the correct form in dotnet-script is `Args[0]`. JSON handling in `.csx` mode (004 JSON filter) is also a persistent pain point, with models reaching for full .NET class patterns that don't compile in script mode.

**Per model and mode, language pass rates:**

Models with thinking variants are split into separate rows. For models without thinking (Claude, qwen3-coder, qwen2.5-coder, Apple), one row is shown.

| Model | Mode | Python | TypeScript | Go | C# |
|-------|------|--------|-----------|----|----|
| `claude-opus-4-6` | — | 100% | 100% | 100% | 100% |
| `claude-sonnet-4-6` | — | 100% | 100% | 100% | 100% |
| `claude-haiku-4-5` | — | 100% | 100% | 100% | 100% |
| `gemma4:31b` | think | **100%** | **100%** | 96% | 75% |
| `gemma4:31b` | nothink | 92% | 96% | 88% | 67% |
| `qwen3-coder:30b` | — | 88% | 88% | 88% | 88% |
| `qwen3.8:27b` | think | 96% | **100%** | 96% | 67% |
| `qwen3.8:27b` | nothink | 88% | 88% | 83% | 79% |
| `qwen3.6:35b` | think | 88% | 88% | 79% | 75% |
| `qwen3.6:35b` | nothink | 79% | 88% | 88% | 71% |
| `qwen3.5:35b-a3b-coding-nvfp4` | think | 88% | 88% | 83% | 75% |
| `qwen3.5:35b-a3b-coding-nvfp4` | nothink | 88% | 83% | 83% | 62% |
| `qwen3.6:35b-a3b-coding-nvfp4` | nothink | 88% | 88% | 83% | 62% |
| `qwen3.6:35b-a3b-coding-nvfp4` | think | 62%† | 88% | 83% | 79% |
| `qwen3.5:27b` | think | 88% | 88% | 88% | 75% |
| `qwen3.5:27b` | nothink | 88% | 79% | 83% | 71% |
| `qwen3.5:27b-nvfp4` | think | 92% | 96% | 88% | 71% |
| `qwen3.5:27b-nvfp4` | nothink | 88% | 92% | 83% | 58% |
| `gemma4:26b` | think | 88% | **100%** | 71% | 92% |
| `gemma4:26b` | nothink | 83% | 83% | 67% | 83% |
| `qwen2.5-coder:7b` | — | 83% | 67% | 29% | 33% |
| `qwen3.5:4b` | think | 88% | 67% | 50% | **58%** |
| `qwen3.5:4b` | nothink | 71% | 50% | 33% | 25% |
| `qwen3.5:4b-nvfp4` | think | 75% | 58% | 33% | 29% |
| `qwen3.5:4b-nvfp4` | nothink | 71% | 42% | 29% | 21% |
| `apple-foundationmodel` | — | **75%** | 38% | 12% | 12% |

> † Python rate for `qwen3.6:35b-a3b-coding-nvfp4` think is suppressed by timeout failures in Python cells across the 3 runs. The underlying Python capability is likely comparable to the nothink variant.

Notable patterns:
- `gemma4:31b` think achieves **100% in both Python and TypeScript** — matching Claude's perfect mark in those languages. Go (96%) is near-perfect; C# (75%) is the only meaningful weakness.
- `qwen3-coder:30b` achieves perfectly uniform 88% across all four languages — the most balanced non-Claude model by language.
- `qwen3.8:27b` think reaches **100% TypeScript** and 96% in both Python and Go, but drops to 67% in C# — the widest single-language gap of any top-10 configuration. Nothink is more even (88/88/83/79) and is the only configuration where C# is *not* the weakest language.
- `qwen3.6:35b` nothink shows an unusual pattern: 88% in TypeScript and Go, but only 79% in Python. Think mode recovers Python to 88% but Go drops to 79%.
- Thinking mode is especially impactful on **C#** for smaller models: `qwen3.5:4b` jumps from 25% to 58% (+33pp). C# requires very specific `.csx` scripting idioms that smaller models recall more reliably with extended reasoning.
- `gemma4:26b` think reaches **100% TypeScript** — matched by `gemma4:31b` think, making the Gemma 4 family distinctively strong on TypeScript in think mode.
- `qwen3.5:27b-nvfp4` nothink has an unusual TypeScript peak (92%) that exceeds its Python rate (88%), and think mode pushes it to 96%. This appears consistent across runs.
- `apple-foundationmodel` shows an extreme cliff: 75% Python, then 38% TypeScript, 12% Go, 12% C#.

---

## Test Difficulty

Score and pass rate across all models for each test. Score reflects the per-cell check fraction averaged over all cells; pass rate is binary end-to-end success.

### 006 — Bug fix — score **98** / pass 98%

> A FizzBuzz implementation with a range bug. Fix the bug; return only the corrected source code.

The highest-scoring test and the most binary: score equals pass rate almost exactly, meaning failures are near-total misses with minimal partial credit. The fix requires identifying that the loop starts at 0 instead of 1 and terminates at `< n` instead of `<= n`. Nearly every model gets this right across all languages. The few failures are Go compile errors from models that produced syntactically invalid code, plus one C# solution that leaked a markdown backtick into the output.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 100% | 99% | 96% | 97% |

---

### 008 — Prime numbers — score **89** / pass 89%

> Given integer N as a CLI argument, print all prime numbers up to N, one per line.

Score and pass rate are equal — this test has little partial credit opportunity; failures tend to be complete. Failures split roughly evenly between Go/C# compile errors and logic issues. Among logic failures, models occasionally produce a header line or include 1 in the output. 1 is not a prime number by definition (primes must be greater than 1; this has been the standard mathematical definition since the late 19th century) — the prompt doesn't need to say so explicitly, and models that include it are making a mathematical error. Python is nearly universal at 99%; Go and C# trail due to compile failures.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 99% | 93% | 86% | 76% |

---

### 005 — Unit test writer — score **88** / pass 81%

> Given a function implementation, write unit tests that print `PASS: <description>` or `FAIL: <description>` to stdout for each assertion.

The score/pass gap (+7) is the largest in the benchmark after 007. Two failure modes dominate. First: models reach for testing frameworks (pytest, jest, `testing.T`, NUnit) instead of the plain-output format — the tests run cleanly but produce no `PASS:` output, earning only the `ran_clean` check. Second: models compute wrong expected values for the discount tiers, causing `FAIL:` lines when the correct function disagrees with their expectations. Since `ran_clean` is nearly universal (the function under test is valid code that compiles), nearly every failure earns 1/3 checks regardless of output quality — inflating the score above the pass rate.

Language scores are relatively even: this is a format comprehension problem, not a language-specific compilation challenge.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 88% | 85% | 75% | 76% |

---

### 001 — CSV to JSON — score **82** / pass 81%

> Read a CSV file and output it as a JSON array with type coercion: age as integer, score as float.

Score slightly exceeds pass rate: some failures produce structurally valid JSON with correct row count but fail type coercion, earning 2/5 checks. The type requirement is the specific trip point — models parse the CSV correctly and produce valid JSON but output all values as strings, failing `age_is_int` and `score_is_float`. Python (93%) handles type inference naturally. C# (72%) and Go (68%) trail on compile errors and stricter type coercion requirements.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 93% | 89% | 68% | 72% |

---

### 002 — Word frequency — score **81** / pass 81%

> Read a text file, count word frequency (case-insensitive), print top 10 sorted by frequency descending, alphabetically for ties.

The word counting and case folding are rarely wrong. The trip point is tie-breaking: the test data has words tied at the same frequency that must appear alphabetically when frequency is equal. Models that sort only by frequency fail the `ties_sorted` check and typically fail all five checks together — it's an all-or-nothing miss when it fails. Python (96%) handles this naturally; Go (72%) and C# (72%) trail on compile failures and string-handling idiom mismatches.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 96% | 82% | 72% | 72% |

---

### 004 — JSON filter — score **79** / pass 79%

> Read a JSON array of objects; output only records where `active == true` and `score >= threshold`, sorted by score descending.

Score and pass rate are equal — failures earn no partial credit. Python (97%) is near-perfect: list comprehensions and `json` module make this near-trivial. TypeScript (94%) and Go (88%) perform well. C# (38%) is the outlier: reading a JSON array, deserializing to typed objects, filtering with LINQ, and re-serializing in `.csx` script mode requires `JsonDocument`, `JsonNode`, or `dynamic` patterns; most models default to full-project approaches (`.Select`/`.Where` on `JsonNode`, typed `ObjectReader`) that fail to compile. This is the single largest language gap in the benchmark.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 97% | 94% | 88% | 38% |

---

### 003 — Fibonacci — score **79** / pass 79%

> Accept integer N as CLI argument. Print all Fibonacci numbers up to N, one per line. Sequence starts: 1, 1, 2, 3, 5...

Score equals pass rate — both checks fail together when anything is wrong, so no partial credit exists. Compile errors dominate the failure pool: Go and C# together account for roughly half of all failures. Among logic failures, the common errors are starting at 0 instead of 1 (mathematical convention vs. the prompt's explicit spec), outputting only one leading 1 (`1, 2, 3, 5...`), or truncating due to C#'s wrong arg parsing pattern.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 92% | 83% | 76% | 65% |

---

### 007 — Beatles interview — score **44** / pass 16%

> Read a CSV of Beatles members. Using a provided example JSON to infer the transformation, produce a JSON array with each member's first name, last name, birthday, age (at death for deceased, current age for living), and non-null relatives.

The benchmark's hardest test. It requires: parsing a CSV with mixed present/absent death dates, inferring the transformation from an example JSON, computing ages correctly via date arithmetic (handling living vs. deceased differently), and matching the exact JSON structure. 47 records fully passed across 288 attempts — dominated by Claude and gemma4 configurations, with `qwen3.8:27b` think as the strongest non-gemma local result.

Score (44) is more than 2.5× pass rate (16%), revealing that many models produce structurally valid output while failing the computed fields. The per-model breakdown shows a clear capability hierarchy:

| Model | Mode | Score | Passes |
|-------|------|-------|--------|
| `claude-haiku-4-5` | — | **100** | 4/4 |
| `claude-opus-4-6` | — | **100** | 4/4 |
| `claude-sonnet-4-6` | — | **100** | 4/4 |
| `gemma4:31b` | think | **92** | 11/12 |
| `gemma4:31b` | nothink | **83** | 5/12 |
| `qwen3.6:35b-a3b-coding-nvfp4` | think | **73** | 1/12 |
| `qwen3.6:35b` | think | **61** | 1/12 |
| `qwen3.8:27b` | think | **58** | 7/12 |
| `gemma4:26b` | nothink | **54** | 0/12 |
| `qwen3.5:27b` | think | **52** | 0/12 |
| `qwen3.6:35b` | nothink | **52** | 0/12 |
| `qwen3.8:27b` | nothink | **51** | 0/12 |
| `gemma4:26b` | think | **50** | 6/12 |
| `qwen3-coder:30b` | — | **50** | 0/12 |
| `qwen3.5:27b` | nothink | **43** | 0/12 |
| `qwen3.5:27b-nvfp4` | nothink | **42** | 1/12 |
| `qwen3.5:27b-nvfp4` | think | **42** | 3/12 |
| `qwen3.6:35b-a3b-coding-nvfp4` | nothink | **42** | 0/12 |
| `qwen3.5:35b-a3b-coding-nvfp4` | think | **40** | 0/12 |
| `qwen3.5:35b-a3b-coding-nvfp4` | nothink | **30** | 0/12 |
| `qwen2.5-coder:7b` | — | **16** | 0/12 |
| `qwen3.5:4b-nvfp4` | think | **8** | 0/12 |
| `qwen3.5:4b` | think | **5** | 0/12 |
| `qwen3.5:4b` | nothink | **4** | 0/12 |
| `apple-foundationmodel` | — | **0** | 0/12 |
| `qwen3.5:4b-nvfp4` | nothink | **0** | 0/12 |

`gemma4:31b` think (11/12) is the standout — nearly matching Claude and far ahead of every other local model. Its nothink variant also passes outright 5/12 times (score 83). The 31B scale appears to cross a reasoning threshold the 26B model could not reliably reach.

`qwen3.8:27b` think is the second-strongest local result by full passes (7/12), despite a lower score (58) than the two 1-pass qwen3.6 think configs — because its failures are all-or-nothing. In think mode, every 007 check lands at exactly 7/12: the seven cells that compile pass all 11 checks, and the five that fail (all three C# cells plus one Go and one Python) fail every check. Thinking fixes the age-at-death reasoning completely; the residual failures are compilation, not logic. Its nothink variant tells the opposite story — it gets structurally close (8/12 on `valid_json`, most checks 6–8/12) but scores 0/12 on `age_john` and 1/12 on `age_george`: it computes current age for the deceased members rather than age at death, and never earns a full pass.

The MoE models (`qwen3-coder:30b`, `qwen3.5:35b-a3b`) score 30–50 with zero full passes — structurally correct output with wrong computed fields. `qwen3.6:35b-a3b-coding-nvfp4` think (score 73, 1/12) is the first MoE model to clear a full 007 pass and posts the highest 007 score outside gemma4. `qwen3.6:35b` think also reaches 1/12. The most common failure for mid-tier models: wrong age calculation for John Lennon and George Harrison specifically, because models compute current age from birthdate without checking the death date field.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 17% | 25% | 11% | 12% |

---

## Per-Model Profiles

### claude-opus-4-6 — score 100 / pass 100%

Perfect on both metrics across all 32 test/language combinations. No compile errors, no logic failures, no runtime errors. The reference ceiling.

### claude-sonnet-4-6 — score 100 / pass 100%

Matches Opus on every metric. No failures in any test or language. The two models are interchangeable on this benchmark.

### claude-haiku-4-5 — score 100 / pass 100%

Clean sweep — all 32 cells passing. A notable result: Haiku is significantly cheaper and faster than Sonnet or Opus, yet matches them perfectly here. Well-defined codegen tasks with clear correctness criteria fall well within Haiku's capability range.

### qwen3-coder:30b — score 94 / pass 88%

**The top local model.** MoE architecture (30B total, ~3.3B active) delivers ~81 tok/s and ~4s average per task — the fastest in the benchmark — while posting the highest local score. All 12 failures are on 007; it passes every other test in every language without exception, and its language balance is perfectly uniform at 88%. On 007 it produces valid JSON with correct structure but fails age calculations (especially John Lennon and George Harrison), earning ~50% partial credit per cell but never a full pass. Failure modes are almost entirely logic-based (9 logic, 3 compile). Run-to-run consistency is 100% — every test/language cell resolves the same way across all 3 runs, the only local model to achieve that. Requires ~18–20GB RAM; comfortable on 48GB, tight on 36GB.

### gemma4:31b — score 93 think / 91 nothink — pass 93% / 85%

**The reasoning leader.** Think mode (score 93 / 93% pass) is second only to qwen3-coder among local configs, and it is the clear #1 on the hardest task: **11 out of 12 full passes on 007** (score 92 for that test), versus 6/12 for gemma4:26b think and 3 or fewer for every other local model bar qwen3.8:27b think (7/12). The 31B scale crosses a reasoning threshold the 26B approached but could not sustain.

Nothink mode (score 91 / 85% pass) is also strong, with 5/12 007 passes. Language balance in think mode: Python 100%, TypeScript 100%, Go 96%, C# 75%. Nothink: Python 92%, TypeScript 96%, Go 88%, C# 67%. Non-007 failures concentrate on 003 fibonacci and 008 prime numbers (Go/C# compile errors).

The downside is speed: ~17–18 tok/s dense inference at 31B parameters. Nothink averages ~24s/task; think ~79s. Think mode is the recommended configuration if 007-level multi-step reasoning tasks appear in your workload; the latency places it in the batch/background tier for interactive use. Requires ~18–20GB RAM.

### qwen3.8:27b — score 91 nothink / 90 think — pass 84% / 90%

**The strongest dense Qwen.** Nothink (score 91) ties `gemma4:31b` nothink for the best non-thinking local score, and it does so at similar latency (~23s/task, ~18 tok/s). Against its predecessor `qwen3.5:27b` nothink it is +5 score at comparable speed — a clean generational gain. Non-007 performance is excellent: 7 of 8 tests at 92%+ pass rate, C# the only language under 83%.

Think mode is a targeted trade rather than an upgrade. The aggregate score drops 1 point, but the composition changes sharply: on the seven easy tests it regresses by ~2 passes, while on 007 it goes from 0 to 7 full passes — the second-most of any local configuration. 007 in think mode is binary: cells that compile pass all 11 checks, cells that don't (all 3 C#, plus one Go and one Python) pass none. The C# regression is visible elsewhere too — think-mode C# drops to 67% (from 79% nothink) as the extended reasoning leads it toward `.csx`-incompatible library patterns.

Failure breakdown: nothink 7 compile / 8 logic; think 7 compile / 2 logic / 1 runtime. Consistency: 91% nothink, 88% think. Requires ~16–18GB RAM (dense 27.3B weights). Use nothink as the default; switch to think only for 007-class tasks.

### qwen3.6:35b — score 90 think / 89 nothink — pass 82% / 81%

Score 89 nothink at ~11s trails only `qwen3.8:27b` (91) and `gemma4:31b` (91) among nothink local configs — and at less than half their time. Despite being a dense 35B model it runs considerably faster than the smaller `qwen3.5:27b` (~31s nothink), reflecting efficiency gains in the 3.6 generation.

Think mode (+1 point, 8.4× time cost to ~89s) is among the worst thinking ROI in the benchmark. The nothink model is already nearly fully capable. Use nothink.

Language balance nothink: TypeScript 88%, Go 88%, C# 71%, Python 79% — the low Python rate is unexpected relative to the overall score. 007 Beatles: 0 full passes nothink (score 52), 1 full pass think (score 61). Consistency: 94% nothink, 81% think. Requires ~18–20GB RAM (full 35B dense weights).

### gemma4:26b — score 88 think / 86 nothink — pass 88% / 79%

Strong at the 26B weight class and the fastest dense model at this capability tier (~9s nothink, ~43s think). Think mode reaches score 88 / 88% pass and enables 007 passes (6/12) — the most of any local model except gemma4:31b and qwen3.8:27b think. Nothink (score 86 / 79%) remains competitive. Language balance differs by mode: nothink sits near 83% for Python, TypeScript, and C# with Go lagging at 67%; think pushes TypeScript to 100% and C# to 92% but Go only reaches 71%. Failure modes are well-distributed (nothink: 9 compile, 10 logic, 1 runtime). Requires ~14–16GB RAM.

### qwen3.5:35b-a3b-coding-nvfp4 — score 88 think / 83 nothink — pass 83% / 79%

**Speed leader among non-trivial local models.** MoE architecture (35B total, ~3B active) delivers ~82 tok/s. Nothink averages ~5s/task, think ~13s. At those latencies the scores are exceptional: 83 nothink and 88 think.

Think mode is unusually cheap: +5 score for 2.7× time cost from a fast baseline, producing a ~13s fully-interactive thinking experience — no other model delivers a thinking result this quickly at this quality.

Language balance nothink: Python 88%, TypeScript 83%, Go 83%, C# 62%. C# is the persistent weak point at both settings. 007 Beatles: 0 passes in both modes (nothink 30, think 40). Like qwen3-coder, the sparse activation pattern appears to cap 007-level multi-step reasoning. Consistency: 94% nothink, 91% think. Requires ~18–20GB RAM for weight loading despite the ~3B active-per-token inference cost.

### qwen3.6:35b-a3b-coding-nvfp4 — score 87 think / 85 nothink — pass 78% / 80%

The MoE successor to `qwen3.5:35b-a3b`. Same architecture: 35B total, ~3B active, code-specialized, nvfp4. Inference speed unchanged at ~5s nothink.

Score 85 nothink (+2 over the qwen3.5 predecessor) with language balance Python 88%, TypeScript 88%, Go 83%, C# 62%. Think mode (+2 over nothink, 9.2× time cost to ~43s) is poor ROI. It also produced 8 timeouts across 96 attempts (concentrated in Python cells), where the extended reasoning chain exceeded the benchmark threshold — this suppresses the measured think score and signals instability. Nothink is the right mode.

007 Beatles think: score 73 (1/12 passes) — the highest 007 score of any MoE model and the first MoE full 007 pass. Nothink 007: score 42, 0 passes. Consistency: 91% nothink, 78% think (the think decline is driven by timeouts). Requires ~18–20GB RAM.

### qwen3.5:27b — score 91 think / 86 nothink — pass 84% / 80%

Thinking mode is the story: +5 score, +4 pass rate, at ~3.6× the time (~31s → ~109s). In think mode it scores 91 — tied for the best local score outside qwen3-coder and gemma4:31b think — but the ~109s latency confines it to batch/background use. Nothink (86) is competitive but unremarkable. Think-mode failures: a handful of C# compile/runtime errors, two timeouts, and 007. Consistency is 100% in think mode, 81% nothink. Requires ~16–18GB RAM.

### qwen3.5:27b-nvfp4 — score 89 think / 86 nothink — pass 86% / 80%

The nvfp4 variant is lossless in nothink (86, matching standard weights) at 32% faster generation — the mode where nvfp4 clearly earns its place. In think mode it trades 2 score points for only 5% faster generation, a poor exchange. A consistent quirk: TypeScript is this model's strongest language — 92% nothink (highest among nothink local configs) and 96% think. Consistency: 91% think, 84% nothink. Requires ~14–16GB RAM.

### qwen2.5-coder:7b — score 59 / pass 53%

The weakest of the explicitly Ollama-targeted models. Go (29%) and C# (33%) are severe weak points — nearly all compile errors from missing or undefined imports and wrong `.csx` idioms. Python (83%) is its strongest area by far. The main advantage is minimal footprint: ~5GB, ~4.9s per task, runs on any machine. Codegen tasks requiring more than one logical step consistently expose the 7B parameter ceiling. Consistency 81%.

### qwen3.5:4b — score 68 think / 48 nothink — pass 66% / 45%

Thinking mode transforms this model: +20 points, +21 pass rate points. Nothink is weak across the board (Python 71%, TypeScript 50%, Go 33%, C# 25%). Think shifts the picture: Python 88%, TypeScript 67%, Go 50%, C# 58%. C# benefits most from the reasoning pass (25% → 58%), likely because `.csx` idioms require deliberate recall. Compile errors dominate the failure pool. Consistency is low (59–66%) — high run-to-run variance compared to larger models. Requires ~3–4GB RAM.

### qwen3.5:4b-nvfp4 — score 55 think / 44 nothink — pass 49% / 41%

The nvfp4 variant is most impactful in think mode: 58% faster than standard 4b think at the cost of 13 score points. Nothink is essentially standard speed at 4 points lower — not recommended. Think mode at score 55 clears standard nothink (48) but trails standard think (68) badly, so the quantization only makes sense when the ~19s think latency (vs ~45s standard) is the deciding factor. The compile error rate is the highest of any model, making Go and C# near-unusable. Consistency 59–72%. Requires ~3GB RAM.

---

## Apple FoundationModel

**Score: 39 / Pass rate: 34% (33/96)**

The Apple on-device model ran via [apfel](https://github.com/Arthur-Ficial/apfel), an OpenAI-compatible wrapper for Apple's FoundationModels framework on macOS 26 Tahoe. It was tested on the same 8 tests × 4 languages, 3 runs each.

### Accuracy breakdown

| Language | Score | Pass Rate |
|----------|-------|-----------|
| Python | **79** | 75% |
| TypeScript | **43** | 38% |
| Go | **18** | 12% |
| C# | **17** | 12% |

| Test | Score | Pass Rate |
|------|-------|-----------|
| 006 — Bug fix | **92** | 92% |
| 008 — Prime numbers | **58** | 58% |
| 005 — Unit test writer | **44** | 8% |
| 001 — CSV to JSON | **43** | 42% |
| 002 — Word frequency | **25** | 25% |
| 003 — Fibonacci | **25** | 25% |
| 004 — JSON filter | **25** | 25% |
| 007 — Beatles interview | **0** | 0% |

The model performs adequately in Python (score 79 / 75% pass) — competitive with qwen2.5-coder:7b in that language specifically — and collapses in TypeScript (43 / 38%), Go (18 / 12%), and C# (17 / 12%). Bug fix stands out at score 92 / 92%: the model handles the FizzBuzz correction well across languages, outperforming most Ollama models on that specific test. The unit test writer score (44) is far above its pass rate (8%) — a `ran_clean` partial credit pattern, where code compiles but produces no matching `PASS:` output.

### Resource profile

The most distinctive characteristic of this model is what it doesn't do. During every Ollama model run, GPU utilization was high, fans ran at maximum, and the machine was audibly working. During Apple model runs, neither the GPU nor CPU showed meaningful load — Activity Monitor showed essentially idle behavior on both, and the fans stayed completely off throughout all inference runs.

The observable effect:

- **Battery life** — No GPU saturation means dramatically lower power draw. Sustained Ollama inference on a MacBook Pro drains the battery at GPU-level rates; this model doesn't.
- **Thermals** — Critical for fanless machines. An 8GB MacBook Air will thermal-throttle under sustained Ollama inference; this model has no thermal footprint.
- **Device coverage** — Runs on any Apple Silicon Mac with Apple Intelligence enabled, including 8GB devices where even qwen3.5:4b is impractical.

### Verdict

Not competitive with mid-tier Ollama models for multi-language code generation. Scores below 20 in Go and C# disqualify it as a general coding backend. As a Python-only zero-configuration assistant on constrained hardware — specifically 8GB MacBook Airs where Ollama isn't viable — it has a niche. The thermal profile is real and matters in those contexts.

---

## Speed vs. Quality

Claude generation times were not instrumented; those rows are omitted from this table.

| Model | Mode | Avg time/task | tok/s | Score | Pass Rate |
|-------|------|--------------|-------|-------|-----------|
| `qwen3-coder:30b` | — | ~4s | ~81 | **94** | 88% |
| `qwen3.6:35b-a3b-coding-nvfp4` | nothink | ~5s | ~86 | **85** | 80% |
| `qwen3.5:35b-a3b-coding-nvfp4` | nothink | ~5s | ~82 | **83** | 79% |
| `qwen2.5-coder:7b` | — | ~5s | ~73 | 59 | 53% |
| `apple-foundationmodel` | — | ~7s | ~64 | 39 | 34% |
| `qwen3.5:4b` | nothink | ~8s | ~49 | 48 | 45% |
| `qwen3.5:4b-nvfp4` | nothink | ~8s | ~86 | 44 | 41% |
| `gemma4:26b` | nothink | ~9s | ~74 | **86** | 79% |
| `qwen3.6:35b` | nothink | ~11s | ~37 | **89** | 81% |
| `qwen3.5:35b-a3b-coding-nvfp4` | think | ~13s | ~82 | **88** | 83% |
| `qwen3.5:4b-nvfp4` | think | ~19s | ~85 | 55 | 49% |
| `qwen3.5:27b-nvfp4` | nothink | ~21s | ~20 | 86 | 80% |
| `qwen3.8:27b` | nothink | ~23s | ~18 | **91** | 84% |
| `gemma4:31b` | nothink | ~24s | ~18 | **91** | 85% |
| `qwen3.5:27b` | nothink | ~31s | ~14 | 86 | 80% |
| `gemma4:26b` | think | ~43s | ~72 | 88 | 88% |
| `qwen3.6:35b-a3b-coding-nvfp4` | think | ~43s | ~84 | **87** | 78% |
| `qwen3.5:4b` | think | ~45s | ~49 | 68 | 66% |
| `qwen3.8:27b` | think | ~54s | ~17 | **90** | 90% |
| `gemma4:31b` | think | ~79s | ~17 | **93** | 93% |
| `qwen3.6:35b` | think | ~89s | ~37 | **90** | 82% |
| `qwen3.5:27b-nvfp4` | think | ~103s | ~20 | 89 | 86% |
| `qwen3.5:27b` | think | ~109s | ~14 | **91** | 84% |

The speed/quality leaders at the interactive tier:
- **`qwen3-coder:30b`** — score 94 in ~4s. The fastest model in the benchmark is also the highest-scoring local one. There is no tradeoff to make here; it is the default local recommendation.
- **`qwen3.6:35b-a3b-coding-nvfp4` nothink** — score 85 in ~5s. The best score under 10s after qwen3-coder.
- **`qwen3.6:35b` nothink** — score 89 in ~11s. Second-highest nothink score, at a latency that is still comfortably interactive.
- **`qwen3.5:35b-a3b-coding-nvfp4` think** — score 88 in ~13s. The only thinking-mode configuration with interactive latency; every other think config is ~43s or slower.
- **`qwen3.8:27b` nothink** — score 91 in ~23s. The highest-scoring dense nothink option, but it costs ~5× the latency of qwen3-coder for a 3-point *lower* score. Prefer it over qwen3-coder only when qwen3-coder's specific 007 weakness matters and think mode isn't an option.

The `qwen3.5:4b-nvfp4` models are notably fast (~85–86 tok/s) due to MLX memory bandwidth — faster token generation than anything else in this table, but the quality floor at 4B limits their usefulness.

---

## Summary and Recommendations

### Top 3 — Non-thinking (interactive use)

1. **`qwen3-coder:30b`** — score 94, pass 88%, ~4s/task, ~18–20GB RAM, 100% run consistency. The fastest and highest-scoring local model, full stop. Passes every test in every language except 007. The default choice unless you specifically need reliable multi-step reasoning.

2. **`qwen3.8:27b` nothink** — score 91, pass 84%, ~23s/task, ~16–18GB RAM. The highest-scoring dense nothink model. Slower than the MoE options but the most capable single non-thinking pass, and the only configuration where C# is not the weakest language.

3. **`qwen3.6:35b` nothink** — score 89, pass 81%, ~11s/task, ~18–20GB RAM. The middle ground: 4 points below qwen3.8:27b nothink at half the latency.

**Speed pick:** `qwen3.6:35b-a3b-coding-nvfp4` nothink (score 85, ~5s, ~18–20GB) — MoE latency at a score that ties qwen3-coder's predecessors.
**36GB machines:** `gemma4:26b` nothink (score 86, ~9s, ~14–16GB) — the qwen3.6 and MoE options need 18–20GB. `qwen3.5:27b-nvfp4` nothink (score 86, ~21s, ~14–16GB) is the next step up if speed is less critical.

### Top 3 — Thinking (quality-first use)

1. **`gemma4:31b` think** — score 93, pass 93%, ~79s/task, ~18–20GB RAM, 97% think-mode consistency. The reasoning ceiling: 11/12 on 007, the only local model to approach Claude there. Best when quality is the priority and ~79s/turn latency is acceptable.

2. **`qwen3.5:27b` think** — score 91, pass 84%, ~109s/task, ~16–18GB RAM, 100% consistency. Ties for the best local score outside qwen3-coder. The latency confines it to batch/background use, but its perfect run consistency makes it ideal there.

3. **`qwen3.5:35b-a3b-coding-nvfp4` think** — score 88, pass 83%, ~13s/task, ~18–20GB RAM. The only thinking configuration with interactive latency. First choice when you want a reasoning pass without batch-tier wait times.

**007-class tasks:** `qwen3.8:27b` think (score 90, ~54s, 7/12 on 007) is the fastest configuration that reliably clears the Beatles test outright, and faster than both gemma4:31b think and qwen3.5:27b think.

### Quick reference by machine

| Machine | Recommendation | Notes |
|---------|---------------|-------|
| 48GB (M3 Max, default) | `qwen3-coder:30b` | Score 94 at ~4s; top score and top speed |
| 48GB (M3 Max, max nothink quality) | `qwen3.8:27b` nothink | Score 91 at ~23s; best single non-thinking pass |
| 48GB (M3 Max, interactive thinking) | `qwen3.5:35b-a3b-coding-nvfp4` think | Score 88 at ~13s — the only interactive think option |
| 48GB (M3 Max, hardest tasks) | `gemma4:31b` think | Score 93, 11/12 on 007; ~79s/task |
| 36GB (M3 Pro) | `gemma4:26b` nothink | ~14–16GB, score 86 at ~9s; qwen3.5:27b-nvfp4 for a thinking fallback |
| 16GB (M3 Air) | `qwen3.5:4b` think | Score 68 with ~45s latency; usable for Python/TypeScript |
| Any Mac (zero setup) | `apple-foundationmodel` | Python only; everything else fails |

### Cloud vs. local gap

All three Claude models score 100 — a clean sweep at every level of difficulty. The best local model, `qwen3-coder:30b`, now reaches score 94, and `gemma4:31b` think reaches 93. For routine tasks (001–006, 008), the top local models hit 88–99% pass rates and are fully competitive. The remaining gap is concentrated on 007-level reasoning: `gemma4:31b` think (11/12) and `qwen3.8:27b` think (7/12) are the only local configurations that meaningfully close it, and both still trail Claude's 4/4. For Python- or TypeScript-heavy work with well-defined tasks, the best local models are viable daily-driver alternatives. For tasks requiring the combination of ambiguity tolerance, multi-file schema inference, and precise date arithmetic that 007 exercises, Claude remains the reliable choice.

---

## Real-World Usage: Agentic Harness Potential

This benchmark measures first-shot generation — one prompt, one response, no feedback. A real coding assistant harness (Claude Code, Open Code, Aider, etc.) works differently: the model sees file contents, runs code, reads error output, and iterates over multiple turns. That changes the picture meaningfully.

### How a harness changes the failure profile

The dominant failure category for local models in this benchmark is **compile errors**: Go missing imports, C# `.csx` idiom errors, TypeScript undefined references. In a harness, the model gets the compiler output back and can fix it. A model that generates `./solution.go:4: "strconv" imported and not used` would likely resolve it in one additional tool call. Compile errors that look like model failures here are largely mechanical corrections in an agentic context.

**Logic errors are different** — usually. If the model computes the wrong algorithm and has no way to verify correctness, a harness can't help. Tie-breaking in 002 (word frequency) falls into this category: the model either knows to sort alphabetically on ties or it doesn't, and running the code produces no useful signal either way.

**007 is a different case.** The expected output format (`expected_format.json`) is provided directly in the prompt — the model already has the oracle. In a harness, it could run its code, compare the actual JSON output against the expected format, identify the specific discrepancy (wrong age for John Lennon, incorrect name split), and fix it. The 007 failures in this benchmark are mostly "close but wrong on one field" — exactly the kind of near-miss that iterative tool use is designed to close. `qwen3.8:27b` nothink is the clearest example: it reaches 6–8/12 on most 007 checks and fails only on age-at-death — a discrepancy that becomes obvious the moment its output is diffed against the example.

The practical implication: Go's 72% and C#'s 64% pass rates here are likely pessimistic for harness use. Python's 85% is probably closer to a real ceiling for tests without a built-in reference — but even that ceiling rises for tests like 007 where the model can self-verify against the provided example.

### Tool calling support

Not all models support structured tool calling (function calling). Support is required to use a model as a harness backend at all. The Apple FoundationModel has no tool calling interface via apfel. All Qwen3/3.5/3.6/3.8 and Gemma4 models support it natively through Ollama's tool API.

**Consistency** in the table below is the fraction of (test × language) combinations where all 3 independent runs produced the same outcome — all pass or all fail. High consistency means the model's behavior is predictable; low consistency means the same prompt can succeed or fail depending on random variation, which is frustrating in a harness where you're trying to debug whether a failure is a model problem or a code problem.

| Model | Tool Use | Score | Nothink speed | Consistency | Thinking | RAM |
|-------|----------|-------|---------------|-------------|----------|-----|
| `claude-opus-4-6` | ✓ | 100 | — | (1 run) | — | cloud |
| `claude-sonnet-4-6` | ✓ | 100 | — | (1 run) | — | cloud |
| `claude-haiku-4-5` | ✓ | 100 | — | (1 run) | — | cloud |
| `qwen3-coder:30b` | ✓ | 94 | ~4s | **100%** | — | ~18–20GB |
| `gemma4:31b` | ✓ | 91 | ~24s | 84% | ✓ (93, 97% consistent) | ~18–20GB |
| `qwen3.8:27b` | ✓ | 91 | ~23s | 91% | ✓ (90, 88% consistent) | ~16–18GB |
| `qwen3.6:35b` | ✓ | 89 | ~11s | 94% | ✓ (90, 81% consistent)† | ~18–20GB |
| `qwen3.6:35b-a3b-coding-nvfp4` | ✓ | 85 | ~5s | 91% | ✓ (87, 78% consistent)‡ | ~18–20GB |
| `qwen3.5:35b-a3b-coding-nvfp4` | ✓ | 83 | ~5s | 94% | ✓ (88, 91% consistent) | ~18–20GB |
| `qwen3.5:27b` | ✓ | 86 | ~31s | 81% | ✓ (91, 100% consistent) | ~16–18GB |
| `qwen3.5:27b-nvfp4` | ✓ | 86 | ~21s | 84% | ✓ (89, 91% consistent) | ~14–16GB |
| `gemma4:26b` | ✓ | 86 | ~9s | 88% | ✓ (88, 75% consistent) | ~14–16GB |
| `qwen2.5-coder:7b` | ✓ | 59 | ~5s | 81% | — | ~5GB |
| `qwen3.5:4b` | ✓ | 48 | ~8s | 66% | ✓ (68, 59% consistent) | ~3–4GB |
| `qwen3.5:4b-nvfp4` | ✓ | 44 | ~8s | 59% | ✓ (55, 72% consistent) | ~3GB |
| `apple-foundationmodel` | ✗ | 39 | ~7s | 81% | — | ~4GB§ |

> † `qwen3.6:35b` think: 81% consistent but +1 score over nothink — thinking not recommended for harness use.  
> ‡ `qwen3.6:35b-a3b-coding-nvfp4` think: 78% consistent and had 8 timeouts across 96 attempts; not recommended for harness use.  
> § Apple FoundationModel runs on the Neural Engine with near-zero GPU/CPU load. The ~4GB figure is approximate; actual memory impact is significantly lower than equivalent Ollama models.  
> Thinking mode columns show score and consistency for think mode specifically. Nothink speed is the baseline for interactive harness use.

### Considerations for harness use beyond this benchmark

**Instruction following:** A harness issues multi-turn system-level instructions — "edit only this file," "don't add new dependencies," "stop and ask if ambiguous." This benchmark doesn't measure instruction adherence at all. Models that score well here may still struggle with harness-level directive compliance.

**Context window:** Long coding sessions accumulate file contents, tool outputs, and conversation history. Larger context windows directly reduce the frequency of context truncation mid-task. All Qwen3/3.5/3.6/3.8 and Gemma4 models support 32K+ context via Ollama; `qwen3.8:27b` is natively 256K.

**Turn latency:** Harness interactions often chain many short tool calls (read file, run linter, apply edit). At ~4s/turn, `qwen3-coder:30b` is comfortably interactive. At ~23–31s/turn, the dense 27–35B nothink models are usable but sluggish. Think modes at ~50–109s/turn are impractical for interactive sessions but workable for background tasks.

### Top 3 models to evaluate in an agentic harness — non-thinking

1. **`qwen3-coder:30b`** — Score 94, 100% consistent, ~4s/turn. Highest score, fastest, and perfectly consistent: when something fails you can trust it's the task or environment, not model variance. The clear first candidate.

2. **`qwen3.8:27b` nothink** — Score 91, 91% consistent, ~23s/turn. The highest single-pass quality, and its 007 near-misses (wrong field, right structure) are exactly what harness iteration closes. The latency is the cost.

3. **`qwen3.6:35b` nothink** — Score 89, 94% consistent, ~11s/turn. The best balance of score, consistency, and latency for interactive session use.

### Top 3 models to evaluate in an agentic harness — thinking

1. **`qwen3.5:35b-a3b-coding-nvfp4` think** — Score 88, 91% consistent, ~13s/turn. The only thinking-mode model with interactive latency. First choice when you want a reasoning pass without background-tier waits.

2. **`gemma4:31b` think** — Score 93, 97% consistent, ~79s/turn. The highest-scoring local model and the strongest on 007 (11/12). 97% consistency is the best of any local model in think mode. Best for batch/background agentic tasks where quality is the priority.

3. **`qwen3.5:27b` think** — Score 91, 100% consistent, ~109s/turn. Perfectly consistent — every task reliably passes or reliably fails, ideal for debugging harness behavior. The latency is the tradeoff.

All models listed support tool calling and score well enough that harness iteration would plausibly close the remaining gap on compile-error failures. None have been validated in an actual agentic harness — these are candidates for follow-on testing, not confirmed replacements.

### Context size for harness use

Ollama's default `num_ctx` (2048–4096) is far too small for agentic sessions where context accumulates across file reads, tool outputs, and conversation history. Custom model variants with larger context should be created via `ollama create` before harness testing. Recommended targets for a 48GB machine:

| Model | Recommended ctx | Notes |
|-------|----------------|-------|
| `qwen3-coder:30b` | 131072 (128K) | MoE KV cache is efficient; ~4–8GB at 128K |
| `qwen3.5:35b-a3b-coding-nvfp4` | 131072 (128K) | MoE KV cache; similar efficiency to qwen3-coder |
| `qwen3.6:35b-a3b-coding-nvfp4` | 131072 (128K) | MoE KV cache; same as qwen3.5:35b-a3b |
| `qwen3.8:27b` | 131072 (128K) | Dense 27B, natively 256K; 128K comfortable on 48GB |
| `qwen3.6:35b` | 65536 (64K) | Dense 35B; 128K workable on 48GB but monitor memory |
| `gemma4:31b` | 131072 (128K) | Dense model natively trained at 128K |
| `gemma4:26b` | 131072 (128K) | Natively trained at 128K — no quality degradation |
| `qwen3.5:27b-nvfp4` | 65536 (64K) | Dense attention; 128K workable but monitor memory |

Example Modelfile:
```
FROM qwen3-coder:30b
PARAMETER num_ctx 131072
```

`gemma4` models and the MoE models (`qwen3-coder:30b`, `qwen3.5:35b-a3b`, `qwen3.6:35b-a3b`) are the most confident 128K recommendations: gemma4 because both variants were trained at that length; MoE models because their sparse activation keeps per-token KV cache proportional to active parameters (~3B), not total parameters. `qwen3.8:27b` is natively trained at 256K and handles 128K comfortably. The `qwen3.6:35b` dense model can run at 128K on 48GB but will leave less headroom — 64K is safer. Extending qwen3.5 dense models beyond their 32K training window uses RoPE scaling and may degrade quality on very long contexts.
