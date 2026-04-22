# Code Generation Benchmark — Findings

**Date:** 2026-04-12 to 2026-04-19  
**Machine:** MacBook Pro M3 Max (48GB unified memory)  
**Scope:** 19 model configurations, 8 tests × 4 languages, 3 runs each (Claude models: 1 run each)

---

## Introduction

This benchmark evaluates how well various AI models generate correct, executable code across four languages (Python, TypeScript, Go, C#) and eight tasks ranging from simple algorithms to complex CSV transformation. The goal is to understand which local Ollama models are viable coding assistant alternatives to cloud-hosted Claude models, and whether tradeoffs like quantization or extended thinking modes are worth it in practice.

Each model generates a solution in response to a task prompt. The solution is executed, and a per-test verifier checks correctness. Results capture both a binary pass/fail and a partial-credit check score.

---

## TL;DR

**Claude models hit the ceiling:** Opus, Sonnet, and Haiku all achieved 100% score and pass rate — perfect first-shot performance across all 32 test/language cells. All three variants are in a class of their own.

**`gemma4:31b` with thinking is the local high-water mark:** score 93 / 93% pass. It passes the hardest test (007 Beatles interview) **11 out of 12 times** in think mode — no other local model comes close.

**`qwen3.6:35b` nothink scores 89 in ~11s** — 2 points below gemma4:31b nothink (91) but at less than half the time, and 1 point above `qwen3.5:35b-a3b-coding-nvfp4` think despite running without extended reasoning. It's the highest-scoring local nothink model under ~20s, and a strong daily-driver option on 48GB machines.

**Thinking barely helps qwen3.6 models**: `qwen3.6:35b` think gains only +1 point over nothink at 8.2× the time cost. `qwen3.6:35b-a3b-coding-nvfp4` think gains +2 at 14.2× the time. Both are the worst thinking ROI in the benchmark — nothink is unambiguously the right mode for these models.

**`qwen3.6:35b-a3b-coding-nvfp4` nothink ties `qwen3-coder:30b`** at score 85, at the same ~5s speed, with a +2 improvement over its qwen3.5 predecessor. The MoE speed advantage continues: both models deliver ~5s interactive responses at score 85.

**Thinking mode helps across the board for qwen3.5 models**, but the value scales with model size. The 4b models gain +16 points for a 5.9× time cost; the 27b gains +9 points for a 3.5× cost. For `gemma4:26b`, thinking adds only +1 point at 4.7× the time — nothink is clearly the right mode for that model. `gemma4:31b` follows the same pattern (+2 at 3.3×), but the think mode's 007 performance (11/12) justifies it for complex tasks.

**nvfp4 quantization is nearly lossless at 27B**: -1 point nothink, -4 think, with a 31% and 5% speed gain respectively. At 4B, the tradeoff deteriorates sharply in think mode (-8 points for 57% faster), though remains tolerable in nothink (-2 points, similar speed).

**Apple's FoundationModel scores 30 / 34% pass rate**, which is low but not the whole story. It runs with essentially no thermal signature — no fan noise, no visible GPU load — unlike every Ollama model which saturates the GPU and fans immediately.

---

## Overall Rankings

Each thinking/non-thinking variant is listed separately since they represent meaningfully different configurations. Claude models don't have a thinking field in results (they use standard API mode).

Two metrics are reported for each model:

- **Score** — a nuanced grade of overall capability. Each test is scored as the fraction of its correctness assertions that passed (e.g. a solution that passes 3 of 4 checks contributes 75% for that test), then averaged equally across all checks. This treats tests with more assertions as higher-weight contributors and rewards near-misses rather than treating them as equivalent to complete failures.
- **Pass rate** — binary reliability. What fraction of tests did the solution fully pass end-to-end? This answers the practical question: how often will it actually work?

When score exceeds pass rate, the model is earning partial credit on failures — near-misses that clear some checks. When pass rate exceeds score, failures are clean passes but the model has some weak spots pulling down the average. `qwen3-coder:30b` (score 85 / pass 88%) exhibits the latter: it passes almost everything cleanly but when it fails (exclusively on 007), it earns 50% partial credit from that test's 11 checks while the test's high check count pulls down the overall average disproportionately.

| Rank | Model | Mode | Score | Pass Rate | Avg Time/task |
|------|-------|------|-------|-----------|--------------|
| 1 | `claude-haiku-4-5` | — | **100** | 100% (32/32) | — |
| 1 | `claude-opus-4-6` | — | **100** | 100% (32/32) | — |
| 1 | `claude-sonnet-4-6` | — | **100** | 100% (32/32) | — |
| 4 | `gemma4:31b` | think | **93** | 93% (89/96) | ~79s |
| 5 | `gemma4:31b` | nothink | **91** | 85% (82/96) | ~24s |
| 6 | `qwen3.6:35b` | think | **90** | 82% (79/96) | ~90s |
| 7 | `qwen3.6:35b` | nothink | **89** | 81% (78/96) | ~11s |
| 8 | `qwen3.5:35b-a3b-coding-nvfp4` | think | **88** | 83% (80/96) | ~12s |
| 9 | `qwen3.6:35b-a3b-coding-nvfp4` | think | **87** | 78% (75/96) | ~71s |
| 10 | `qwen3.5:27b` | think | **86** | 84% (81/96) | ~106s |
| 11 | `qwen3.6:35b-a3b-coding-nvfp4` | nothink | **85** | 80% (77/96) | ~5s |
| 11 | `qwen3-coder:30b` | — | **85** | 88% (84/96) | ~4s |
| 13 | `qwen3.5:35b-a3b-coding-nvfp4` | nothink | **83** | 79% (76/96) | ~5s |
| 14 | `qwen3.5:27b-nvfp4` | think | **82** | 86% (83/96) | ~101s |
| 15 | `gemma4:26b` | think | **80** | 88% (84/96) | ~43s |
| 16 | `gemma4:26b` | nothink | **79** | 79% (76/96) | ~9s |
| 17 | `qwen3.5:27b` | nothink | **77** | 80% (77/96) | ~31s |
| 18 | `qwen3.5:27b-nvfp4` | nothink | **76** | 80% (77/96) | ~21s |
| 19 | `qwen3.5:4b` | think | **54** | 66% (63/96) | ~45s |
| 20 | `qwen2.5-coder:7b` | — | **49** | 53% (51/96) | ~5s |
| 21 | `qwen3.5:4b-nvfp4` | think | **46** | 49% (47/96) | ~19s |
| 22 | `qwen3.5:4b` | nothink | **38** | 45% (43/96) | ~8s |
| 23 | `qwen3.5:4b-nvfp4` | nothink | **36** | 41% (39/96) | ~8s |
| 24 | `apple-foundationmodel` | — | **30** | 34% (33/96) | ~7s |

> **Avg time** is average wall clock time per test/language combination from generation start to result. Codegen tasks require 100–500+ output tokens; thinking-mode tasks often exceed 1000 tokens. Claude timing was not instrumented; times will include API round-trip latency and are not comparable to Ollama on-device measurements.

---

## Model Architectures

**Mixture of Experts (MoE):** Rather than activating all parameters for every token, MoE models route each token through a small subset of expert sub-networks. This delivers large-model knowledge at small-model inference cost.

- **`qwen3-coder:30b`** — MoE: 30B total parameters, ~3.3B active per token. This directly explains its combination of top-tier local score and ~4s average inference time.
- **`qwen3.5:35b-a3b-coding-nvfp4`** — MoE: 35B total parameters, ~3B active per token ("a3b" = active 3B). Despite having more total parameters than qwen3-coder, it runs at the same ~5s nothink speed for the same reason — only ~3B parameters fire per token. Weight loading still requires memory proportional to total parameters.
- **`qwen3.6:35b-a3b-coding-nvfp4`** — MoE: same architecture as qwen3.5:35b-a3b (35B total, ~3B active, nvfp4 quantized, code-specialized). The 3.6 generation brings modest score improvement (+2 nothink) at the same ~5s inference speed.

**Code-specialized models:**

- **`qwen2.5-coder:7b`** — Fine-tuned specifically on code corpora. Despite specialization, score 49 / 53% pass rate is the weakest of the larger Ollama models, showing that code-specific training at 7B doesn't compensate for the parameter ceiling on multi-step codegen tasks.
- **`qwen3-coder:30b`** — Also code-specialized, at 30B MoE scale. The combination works: it's the fastest competitive local model.
- **`qwen3.5:35b-a3b-coding-nvfp4`** / **`qwen3.6:35b-a3b-coding-nvfp4`** — Code-specialized MoE variants. The nvfp4 quantization compresses the 35B weights for efficient inference on Apple Silicon.

**General-purpose models evaluated for coding:**

- **`gemma4:26b`** / **`gemma4:31b`** — Google's Gemma 4 family, not code-specialized, yet strong and exceptionally language-balanced. The 31B variant is a dense model (not MoE) that runs slower than the 26B but at significantly higher quality — particularly on 007-level reasoning tasks.
- **`qwen3.5` family** — General-purpose Qwen 3.5 generation with optional thinking mode.
- **`qwen3.6:35b`** — Dense 35B model (not MoE) from the Qwen 3.6 generation. General-purpose with optional thinking. Despite being larger than qwen3.5:27b, it runs faster (~11s vs ~31s nothink), suggesting significant architectural or inference efficiency improvements in the 3.6 generation. Score 89 nothink is the highest of any non-gemma4:31b local model.

---

## Thinking vs. Non-Thinking

| Model | Nothink | Think | Score delta | Time cost |
|-------|---------|-------|-------------|-----------|
| `qwen3.5:4b` | 38 | 54 | **+16** | ~7.7s → ~45s (5.9×) |
| `qwen3.5:4b-nvfp4` | 36 | 46 | **+10** | ~8s → ~19s (2.4×) |
| `qwen3.5:27b` | 77 | 86 | **+9** | ~31s → ~106s (3.5×) |
| `qwen3.5:35b-a3b-coding-nvfp4` | 83 | 88 | **+5** | ~5s → ~12s (2.4×) |
| `qwen3.5:27b-nvfp4` | 76 | 82 | **+6** | ~21s → ~101s (4.8×) |
| `gemma4:31b` | 91 | 93 | **+2** | ~24s → ~79s (3.3×) |
| `gemma4:26b` | 79 | 80 | **+1** | ~9s → ~43s (4.7×) |
| `qwen3.6:35b` | 89 | 90 | **+1** | ~11s → ~90s (8.2×) |
| `qwen3.6:35b-a3b-coding-nvfp4` | 85 | 87 | **+2** | ~5s → ~71s (14.2×) |

**Key observations:**

- For `qwen3.5:4b`, thinking mode gains 16 points — a tier shift. At 4B parameters, the extra reasoning pass overcomes a significant portion of the base model's limitations.
- The gain compresses as model size increases. At 27B, thinking adds 9 points for a 3.5× time cost — still meaningful but not transformative.
- `gemma4:26b` is the exception in the qwen3.5 generation: thinking adds 1 point at 4.7× the time. Nothink is unambiguously the right mode for that model. `gemma4:31b` shows the same pattern (+2 points at 3.3× cost) — but unlike the 26B, the think mode's 007 performance (11/12 passes vs 5/12 nothink) makes think mode compelling for complex tasks even though the score delta is small.
- `qwen3.5:35b-a3b-coding-nvfp4` is the most time-efficient thinking upgrade in this set: +5 points for 2.4× cost starting from an already fast ~5s baseline. The resulting ~12s think latency is fully interactive.
- **`qwen3.6` models show the worst thinking ROI in the benchmark.** `qwen3.6:35b` gains only +1 score point at 8.2× the time cost. `qwen3.6:35b-a3b-coding-nvfp4` gains +2 at 14.2× — the most expensive thinking upgrade relative to benefit of any model tested. Both models appear to be largely "thinking-saturated" in nothink mode; extended reasoning adds little that the base inference doesn't already produce. Thinking is not recommended for either qwen3.6 model.
- For `qwen3.5:4b-nvfp4` think specifically, the speed gain from quantization compresses the time multiplier to 2.4× — the most time-efficient path to a thinking-mode result for a 4B model.

**Recommendation:** Use thinking for `qwen3.5:4b` — the quality gain justifies the wait. Skip thinking for `gemma4:26b`, `qwen3.6:35b`, and `qwen3.6:35b-a3b-coding-nvfp4`. For `gemma4:31b`, prefer think mode if 007-level tasks appear in your workload; otherwise nothink is fine. For `qwen3.5:35b-a3b`, think mode is cheap enough (~12s) to use routinely. For `qwen3.5:27b`, thinking is practical only for batch or background work.

---

## Quantization: nvfp4 vs Standard

The nvfp4 format is an **MLX-native quantization** for Apple Silicon's unified memory architecture. Unlike GGUF (designed for CPU/GPU split inference), MLX models operate natively in Apple's unified memory space.

| Model | Standard | nvfp4 | Score delta | Speed delta |
|-------|----------|-------|-------------|-------------|
| `qwen3.5:27b` nothink | 77 | 76 | **-1** | ~31s → ~21s (32% faster) |
| `qwen3.5:27b` think | 86 | 82 | **-4** | ~106s → ~101s (5% faster) |
| `qwen3.5:4b` nothink | 38 | 36 | **-2** | ~7.7s → ~8s (~same) |
| `qwen3.5:4b` think | 54 | 46 | **-8** | ~45s → ~19s (57% faster) |

**At 27B nothink**, nvfp4 is an excellent tradeoff: 1 point lost for 32% faster generation and lower memory. A clear win when memory is tight.

**At 27B think**, the speed advantage nearly vanishes (5% faster) while costing 4 score points. The thinking-mode token volume is high enough that the quantization speed benefit gets absorbed. Standard weights are preferable when using thinking mode.

**At 4B nothink**, nvfp4 is barely different: 2 points lower at essentially the same speed. Neither compelling nor harmful.

**At 4B think**, nvfp4 is 57% faster with an 8-point score drop. Whether that's acceptable depends on the use case — at a base score of 46, the degradation is more noticeable proportionally.

---

## Language Breakdown

Pass rates aggregated across all models and runs.

| Language | Pass Rate | Notes |
|----------|-----------|-------|
| Python | **84%** | Highest overall; most models are well-calibrated here |
| TypeScript | **75%** | Strong but noticeably more variance across models |
| Go | **61%** | Compiler strictness reliably surfaces missing imports |
| C# | **58%** | `.csx` scripting environment requires specific idioms |

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
| `qwen3.6:35b` | think | 88% | 88% | 79% | 75% |
| `qwen3.6:35b` | nothink | 79% | 88% | 88% | 71% |
| `qwen3.5:35b-a3b-coding-nvfp4` | think | 88% | 88% | 83% | 75% |
| `qwen3-coder:30b` | — | 88% | 88% | 88% | 88% |
| `qwen3.6:35b-a3b-coding-nvfp4` | nothink | 88% | 88% | 83% | 63% |
| `qwen3.5:35b-a3b-coding-nvfp4` | nothink | 88% | 83% | 83% | 63% |
| `qwen3.6:35b-a3b-coding-nvfp4` | think | 63%† | 88% | 83% | 79% |
| `gemma4:26b` | nothink | 83% | 83% | 67% | 83% |
| `gemma4:26b` | think | 88% | **100%** | 71% | 92% |
| `qwen3.5:27b` | nothink | 88% | 79% | 83% | 71% |
| `qwen3.5:27b` | think | 88% | 88% | 88% | 75% |
| `qwen3.5:27b-nvfp4` | nothink | 88% | 92% | 83% | 58% |
| `qwen3.5:27b-nvfp4` | think | 92% | 96% | 88% | 71% |
| `qwen2.5-coder:7b` | — | 83% | 67% | 29% | 33% |
| `qwen3.5:4b` | nothink | 71% | 50% | 33% | 25% |
| `qwen3.5:4b` | think | 88% | 67% | 50% | **58%** |
| `qwen3.5:4b-nvfp4` | nothink | 71% | 42% | 29% | 21% |
| `qwen3.5:4b-nvfp4` | think | 75% | 58% | 33% | 29% |
| `apple-foundationmodel` | — | **75%** | 38% | 13% | 13% |

> † Python rate for `qwen3.6:35b-a3b-coding-nvfp4` think is suppressed by 3 timeout failures in Python cells across 3 runs. The underlying Python capability is likely comparable to the nothink variant.

Notable patterns:
- `gemma4:31b` think achieves **100% in both Python and TypeScript** — matching Claude's perfect mark in those languages. Go (96%) is near-perfect; C# (75%) is the only meaningful weakness.
- `qwen3-coder:30b` achieves perfectly uniform 88% across all four languages — the most balanced non-Claude model by language. `qwen3.6:35b-a3b-coding-nvfp4` nothink matches this (88% in Python and TypeScript, 83% in Go, 63% in C#).
- `qwen3.6:35b` nothink shows an unusual pattern: 88% in TypeScript and Go, but only 79% in Python. This is surprising and may reflect a shift in training focus or test-case interaction in the 3.6 generation. Think mode recovers Python to 88% but Go drops to 79%.
- Thinking mode is especially impactful on **C#** for smaller models: `qwen3.5:4b` jumps from 25% to 58% (+33pp). C# requires very specific `.csx` scripting idioms that smaller models recall more reliably with extended reasoning.
- `gemma4:26b` think reaches **100% TypeScript** — matched by `gemma4:31b` think, making the Gemma 4 family distinctively strong on TypeScript in think mode.
- `qwen3.5:27b-nvfp4` nothink has an unusual TypeScript peak (92%) that exceeds its Python rate (88%), and think mode pushes it further to 96%. This appears consistent and may reflect training data distribution in the quantized weights.
- `apple-foundationmodel` shows an extreme cliff: 75% Python, then 38% TypeScript, 13% Go, 13% C#.

---

## Test Difficulty

Score and pass rate across all models for each test. Score reflects partial credit (checks passed / total checks, check-weighted average); pass rate is binary end-to-end success.

### 006 — Bug fix — score **97** / pass 97%

> A FizzBuzz implementation with a range bug. Fix the bug; return only the corrected source code.

The highest-scoring test and the most binary: score equals pass rate almost exactly, meaning failures are near-total misses with minimal partial credit. The fix requires identifying that the loop starts at 0 instead of 1 and terminates at `< n` instead of `<= n`. Nearly every model gets this right across all languages. Only 5 failures across 168 attempts, all Go compile errors from models that produced syntactically invalid code.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 100% | 98% | 93% | 98% |

---

### 008 — Prime numbers — score **85** / pass 85%

> Given integer N as a CLI argument, print all prime numbers up to N, one per line.

Score and pass rate are nearly equal here too — this test has little partial credit opportunity; failures tend to be complete. Failures split roughly evenly between Go/C# compile errors and logic issues. Among logic failures, models occasionally produce a header line or include 1 in the output. 1 is not a prime number by definition (primes must be greater than 1; this has been the standard mathematical definition since the late 19th century) — the prompt doesn't need to say so explicitly, and models that include it are making a mathematical error. Python is nearly universal at 98%; Go and C# trail due to compile failures.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 98% | 88% | 76% | 76% |

---

### 005 — Unit test writer — score **83** / pass 73%

> Given a function implementation, write unit tests that print `PASS: <description>` or `FAIL: <description>` to stdout for each assertion.

The score/pass gap (+10) is the largest in the benchmark after 007. Two failure modes dominate. First: models reach for testing frameworks (pytest, jest, `testing.T`, NUnit) instead of the plain-output format — the tests run cleanly but produce no `PASS:` output, earning only the `ran_clean` check. Second: models compute wrong expected values for the discount tiers, causing `FAIL:` lines when the correct function disagrees with their expectations. Since `ran_clean` is nearly universal (the function under test is valid code that compiles), nearly every failure earns 1/3 checks regardless of output quality — inflating the score above the pass rate.

Language scores are unusually even: this is a format comprehension problem, not a language-specific compilation challenge.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 90% | 74% | 64% | 62% |

---

### 004 — JSON filter — score **75** / pass 75%

> Read a JSON array of objects; output only records where `active == true` and `score >= threshold`, sorted by score descending.

Score and pass rate are equal — failures earn no partial credit. Python (95%) is near-perfect: list comprehensions and `json` module make this near-trivial. Go (79%) and TypeScript (90%) perform well. C# (36%) is the outlier: reading a JSON array, deserializing to typed objects, filtering with LINQ, and re-serializing in `.csx` script mode requires `JsonDocument` or `dynamic` patterns; most models default to full-project approaches that fail to compile. This is the single largest language gap in the benchmark.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 95% | 90% | 79% | 36% |

---

### 002 — Word frequency — score **73** / pass 71%

> Read a text file, count word frequency (case-insensitive), print top 10 sorted by frequency descending, alphabetically for ties.

The word counting and case folding are rarely wrong. The trip point is tie-breaking: the test data has three words tied at 10 occurrences (`it`, `of`, `was`) that must appear alphabetically when frequency is equal. Models that sort only by frequency fail the `ties_sorted` check and typically fail all five checks together — it's an all-or-nothing miss when it fails. Python (95%) handles this naturally; Go (60%) and C# (62%) trail on compile failures and string-handling idiom mismatches.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 95% | 69% | 60% | 62% |

---

### 003 — Fibonacci — score **72** / pass 72%

> Accept integer N as CLI argument. Print all Fibonacci numbers up to N, one per line. Sequence starts: 1, 1, 2, 3, 5...

Score equals pass rate — both checks fail together when anything is wrong, so no partial credit exists. Compile errors dominate the failure pool: Go and C# together account for roughly half of all failures. Among logic failures, the common errors are starting at 0 instead of 1 (mathematical convention vs. the prompt's explicit spec), outputting only one leading 1 (`1, 2, 3, 5...`), or truncating due to C#'s wrong arg parsing pattern.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 86% | 74% | 60% | 69% |

---

### 001 — CSV to JSON — score **72** / pass 70%

> Read a CSV file and output it as a JSON array with type coercion: age as integer, score as float.

Score slightly exceeds pass rate: some failures produce structurally valid JSON with correct row count but fail type coercion, earning 2/5 checks. The type requirement is the specific trip point — models parse the CSV correctly and produce valid JSON but output all values as strings, failing `age_is_int` and `score_is_float`. Python (93%) handles type inference naturally. C# (55%) and Go (50%) trail on compile errors and stricter type coercion requirements.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 93% | 83% | 50% | 55% |

---

### 007 — Beatles interview — score **34** / pass 13%

> Read a CSV of Beatles members. Using a provided example JSON to infer the transformation, produce a JSON array with each member's first name, last name, birthday, age (at death for deceased, current age for living), and non-null relatives.

The benchmark's hardest test. It requires: parsing a CSV with mixed present/absent death dates, inferring the transformation from an example JSON, computing ages correctly via date arithmetic (handling living vs. deceased differently), and matching the exact JSON structure. 38 records fully passed across 216 attempts — all from Claude models and gemma4 configurations.

Score (34) is more than 2.5× pass rate (13%), revealing that many models produce structurally valid output while failing the computed fields. The per-model breakdown shows a clear capability hierarchy:

| Model | Mode | Score | Passes |
|-------|------|-------|--------|
| `claude-haiku-4-5` | — | **100** | 4/4 |
| `claude-opus-4-6` | — | **100** | 4/4 |
| `claude-sonnet-4-6` | — | **100** | 4/4 |
| `gemma4:31b` | think | **92** | 11/12 |
| `gemma4:31b` | nothink | **83** | 5/12 |
| `qwen3.6:35b-a3b-coding-nvfp4` | think | **73** | 1/12 |
| `qwen3.5:27b` | think | **62** | 0/12 |
| `qwen3.6:35b` | think | **61** | 1/12 |
| `gemma4:26b` | nothink | **54** | 0/12 |
| `qwen3.6:35b` | nothink | **52** | 0/12 |
| `qwen3-coder:30b` | — | **50** | 0/12 |
| `gemma4:26b` | think | **50** | 6/12 |
| `qwen3.5:27b-nvfp4` | think | **46** | 3/12 |
| `qwen3.5:27b` | nothink | **43** | 0/12 |
| `qwen3.5:27b-nvfp4` | nothink | **42** | 1/12 |
| `qwen3.6:35b-a3b-coding-nvfp4` | nothink | **42** | 0/12 |
| `qwen3.5:35b-a3b-coding-nvfp4` | think | **40** | 0/12 |
| `qwen3.5:35b-a3b-coding-nvfp4` | nothink | **30** | 0/12 |
| `qwen2.5-coder:7b` | — | **16** | 0/12 |
| `qwen3.5:4b-nvfp4` | think | **8** | 0/12 |
| `qwen3.5:4b` | think | **5** | 0/12 |
| `qwen3.5:4b` | nothink | **4** | 0/12 |
| `apple-foundationmodel` | — | **0** | 0/12 |
| `qwen3.5:4b-nvfp4` | nothink | **0** | 0/12 |

`gemma4:31b` think (11/12) is the standout result — nearly matching Claude's 4/4 mark and far ahead of every other local model. The one miss is a single language cell where the output was structurally correct but failed a computed field check. `gemma4:31b` nothink also passes outright 5/12 times (score 83), making it the second-strongest local model on this test. The 31B scale appears to cross a reasoning threshold that the 26B model could not reliably reach.

The MoE models (`qwen3-coder:30b`, `qwen3.5:35b-a3b`) both score 30–50 but achieve zero full passes — structurally correct output with wrong computed fields. This pattern is consistent: MoE architectures produce near-correct 007 output but not reliably enough to clear all 11 checks simultaneously.

Thinking mode is what unlocks 007 for gemma4 — the 26B nothink scores 54 but earns zero passes; 26B think reaches 50 (lower score from a different failure distribution) but achieves 6/12 full passes. `qwen3.6:35b-a3b-coding-nvfp4` think (score 73, 1/12 pass) is the first MoE model to achieve a 007 full pass — and its score (73) is the highest of any model that earns 1 or fewer passes, exceeding `qwen3.5:27b` think (62, 0 passes). `qwen3.6:35b` think also achieves 1/12 with score 61. The `qwen3.5:27b` think score (62) remains the highest of any 0-pass model in the qwen3.5 generation, meaning it produces output with significant field-level correctness without clearing all checks simultaneously. The most common failure for mid-tier models: wrong age calculation for John Lennon and George Harrison specifically, often because models compute current age from birthdate without checking the death date field.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 14% | 21% | 7% | 10% |

---

## Per-Model Profiles

### claude-opus-4-6 — score 100 / pass 100%

Perfect on both metrics across all 32 test/language combinations. No compile errors, no logic failures, no runtime errors. The reference ceiling.

### claude-sonnet-4-6 — score 100 / pass 100%

Matches Opus on every metric. No failures in any test or language. Consistent with Opus as the cloud ceiling — the two models are interchangeable on this benchmark.

### claude-haiku-4-5 — score 100 / pass 100%

Clean sweep — all 32 cells passing. A notable result: Haiku is a significantly cheaper and faster model than Sonnet or Opus, yet matches them perfectly on this benchmark. Well-defined codegen tasks with clear correctness criteria appear to fall well within Haiku's capability range.

### qwen3-coder:30b — score 85 / pass 88%

**Fastest competitive local model.** MoE architecture (30B total, ~3.3B active) delivers ~73 tok/s and ~4s average per task. All 12 failures are on 007 (Beatles interview) — it passes every other test in every language without exception. On 007, it produces valid JSON with correct structure but fails age calculations (especially John Lennon and George Harrison). Uniquely balanced across languages: 88% in all four. Failure modes are almost entirely logic-based (9 logic, 3 compile), meaning compile errors are rare. Requires ~18–20GB RAM; comfortable on 48GB, tight on 36GB.

### gemma4:31b — score 93 think / 91 nothink — pass 93% / 85%

**Local ceiling.** Think mode (score 93 / 93% pass) is the highest any local model achieved on this benchmark — 7 points above `qwen3.5:27b` think and 8 points above qwen3-coder. The defining result is **007**: 11 out of 12 full passes in think mode (score 92 for that test), compared to gemma4:26b think's 6/12 and every other local model at 3 or fewer. The 31B scale crosses a reasoning threshold the 26B model approached but could not reliably sustain.

Nothink mode (score 91 / 85% pass) is also strong, with 5/12 007 passes — more than gemma4:26b think achieved. Language balance in think mode: Python 100%, TypeScript 100%, Go 96%, C# 75%. Go and TypeScript are near-perfect; C# is the remaining gap. Nothink: Python 92%, TypeScript 96%, Go 88%, C# 67%. Failures are concentrated on 003 fibonacci and 008 prime numbers (compile errors in Go/C#), with only one 007 miss in think mode.

The downside is speed: ~18 tok/s dense inference at 31B parameters. Nothink averages ~24s/task; think ~79s. Compare to gemma4:26b nothink at ~9s — the 31B is slower for roughly the same nothink score but wins decisively on 007 and overall pass rate. Requires ~18–20GB RAM. Think mode is the recommended configuration if 007-level multi-step reasoning tasks appear in your workload; the latency (~79s) places it in the batch/background tier for interactive use.

### qwen3.6:35b — score 90 think / 89 nothink — pass 82% / 81%

Score 89 nothink at ~11s trails only `gemma4:31b` nothink (91) among nothink local models — and at less than half the time (11s vs ~24s). Notably, it exceeds `qwen3.5:35b-a3b-coding-nvfp4` think (88) without needing extended reasoning. Despite being a dense 35B model, it runs considerably faster than `qwen3.5:27b` nothink (~31s for a smaller 27B model), suggesting meaningful architectural or inference efficiency improvements in the 3.6 generation.

Think mode (+1 point, 8.2× time cost to ~90s) is the worst thinking ROI in the benchmark. The nothink model is already nearly fully capable; extended reasoning delivers no meaningful lift. Use nothink.

Language balance in nothink: TypeScript 88%, Go 88%, C# 71%, Python 79%. The Python rate is unexpectedly low relative to the overall score — nothink failures concentrate on 004 JSON filter (C# scripting idioms) and 002 word frequency (Python ties). Think mode recovers Python to 88% but Go drops to 79%, an unusual tradeoff. 007 Beatles: 0 full passes nothink (score 52), 1 full pass think (score 61). The nothink 007 score (52) is the highest of any Qwen-family nothink model, above `qwen3-coder:30b` (50) and the qwen3.5 variants. The think-mode 007 score (61) narrowly trails `qwen3.5:27b` think (62) but achieves an outright pass (1/12) that qwen3.5:27b never managed (0/12).

Consistency: 94% nothink, 81% think. High nothink consistency makes it reliable for harness use. Requires ~18–20GB RAM (full 35B dense weights).

### gemma4:26b — score 80 think / 79 nothink — pass 88% / 79%

Strong competitor at the 26B weight class, and the fastest dense model at this capability tier (~9s nothink, ~43s think). Think mode pushes it to score 80 / 88% pass and enables 007 passes (6/12) — the most of any local model except gemma4:31b. In nothink mode (score 79 / 79%) it remains highly competitive. Language balance differs by mode: nothink sits at 83% across Python, TypeScript, and C# with Go lagging at 67%; think mode pushes TypeScript to 100% and C# to 92% but Go only reaches 71%. Failure modes are well-distributed (18 compile, 3 runtime, 12 logic), with no single language dominating the failure pool. Requires ~14–16GB RAM.

### qwen3.5:35b-a3b-coding-nvfp4 — score 88 think / 83 nothink — pass 83% / 79%

**Speed leader among non-trivial local models.** MoE architecture (35B total, ~3B active per token) delivers ~82 tok/s — the fastest of any model in the benchmark except the 4B nvfp4 variants. Nothink averages ~5s/task, think ~12s. At those latencies, the score results are exceptional: 83 nothink and 88 think comfortably beat all non-gemma4:31b, non-Claude models.

Think mode is unusually cheap here: +5 score for 2.4× time cost from a fast baseline, producing a ~12s fully-interactive thinking-mode experience. No other model in the benchmark delivers a thinking-mode result this quickly at this quality level.

Language balance in nothink: Python 88%, TypeScript 83%, Go 83%, C# 63%. Think adds ~5pp to C# (75%) and TypeScript (88%) and is largely flat elsewhere. C# is the persistent weak point at both settings. 007 Beatles: 0 passes in both modes (nothink score 30, think score 40). Like qwen3-coder:30b (also MoE), the sparse activation pattern appears to be a ceiling for 007-level multi-step reasoning — it produces partially correct output but can't reliably clear all 11 checks.

Consistency is high: 94% nothink, 91% think — meaning failures are predictable rather than random. Memory footprint: weight loading for a 35B nvfp4 model requires approximately the full quantized weight set (~18–20GB), despite the ~3B active-per-token inference cost.

### qwen3.6:35b-a3b-coding-nvfp4 — score 87 think / 85 nothink — pass 78% / 80%

The MoE successor to `qwen3.5:35b-a3b-coding-nvfp4`. Same architecture: 35B total, ~3B active per token, code-specialized, nvfp4 quantized for Apple Silicon. The inference speed is unchanged at ~5s nothink — MoE speed comes from active parameters, not version.

Score 85 nothink (+2 over qwen3.5:35b-a3b nothink's 83) ties `qwen3-coder:30b` — both at 85, both at ~4–5s. Language balance nothink: Python 88%, TypeScript 88%, Go 83%, C# 63%. C# is the persistent weak point, unchanged from the qwen3.5 generation.

Think mode (+2 over nothink, 14.2× time cost to ~71s) is the worst thinking ROI in the benchmark in absolute time terms — 14× slower for 2 points. Additionally, think mode produced **8 timeouts** across 96 attempts (001 Python, 005 Python/Go, 008 C#, 002 Python, 001 TypeScript), where the extended reasoning chain exceeded the benchmark's timeout threshold. This not only suppresses the measured think score but indicates instability — the model occasionally "runs away" in think mode on certain tasks. The true think score without timeouts may be modestly higher than 87, but the timeout behavior itself is a practical problem. Nothink is the right mode for this model.

007 Beatles think: score 73 (1/12 passes) — the highest 007 score of any MoE or non-gemma4 model, and the first MoE model to achieve a full 007 pass. This is notable; the qwen3.5 generation's a3b scored 40 think with 0 passes. However, the result comes with the timeout caveat above. Nothink 007: score 42, 0 passes — consistent with the qwen3.5 predecessor (score 40).

Consistency: 91% nothink, 78% think. The think consistency decline (from 91% in qwen3.5:35b-a3b think) is driven by timeouts. Requires ~18–20GB RAM.

### qwen3.5:27b — score 86 think / 77 nothink — pass 84% / 80%

Thinking mode is the story here: +9 score, +4 pass rate, at ~3.5× the time (~31s → ~106s). In thinking mode it scores 86 — third among local models, behind gemma4:31b and qwen3.6:35b; in nothink mode it's competitive but unremarkable. Failure breakdown in thinking mode: a handful of C# compile/runtime errors and 004 C# runtime failures (JSON node parent conflict), plus 007. Requires ~16GB RAM.

### qwen3.5:27b-nvfp4 — score 82 think / 76 nothink — pass 86% / 80%

The nvfp4 variant trades 4 score points in think mode for 5% faster generation — a poor exchange at this scale given how small the speed gain is. In nothink mode, -1 point for 32% faster is a strong tradeoff; this is the mode where nvfp4 earns its place. A notable quirk: TypeScript is consistently this model's strongest language — 92% nothink (highest among nothink local models) and 96% think (second only to gemma4 think at 100%). This appears consistent across runs rather than noise.

### qwen2.5-coder:7b — score 49 / pass 53%

The weakest of the explicitly Ollama-targeted models. Score slightly below pass rate (49 vs 53%), which is unusual: it suggests the failures are on tests where the model earns little partial credit. Go (29%) is a notable weak point — nearly all compile errors from missing or undefined imports. C# (33%) is similarly poor. Python (83%) is its strongest area by a significant margin. The main advantage is minimal footprint: ~4.7GB, ~4.9s per task, runs on any machine. Codegen tasks requiring more than one logical step consistently expose the 7B parameter ceiling.

### qwen3.5:4b — score 54 think / 38 nothink — pass 66% / 45%

Thinking mode transforms this model: +16 points, +21 pass rate points. In nothink mode it's a weak performer across the board (Python 71%, TypeScript 50%, Go 33%, C# 25%). Think mode shifts the picture meaningfully: Python reaches 88%, TypeScript 67%, Go 50%, C# 58%. C# in particular — 25% nothink vs 58% think — benefits most from the reasoning pass, likely because the `.csx` idioms require more deliberate recall. Failure mode breakdown shows compile errors dominating (47 total across both modes). High run-to-run variance compared to larger models.

### qwen3.5:4b-nvfp4 — score 46 think / 36 nothink — pass 49% / 41%

The nvfp4 variant is most impactful here in think mode: 57% faster than standard 4b think mode at the cost of 8 score points. The nothink variant is essentially the same speed as standard at 2 points lower — not recommended. Even in think mode, score 46 underperforms standard nothink qwen3.5:4b by 8 points while running faster. The compile error rate is the highest of any model (66 compile failures), making Go and C# near-unusable.

---

## Apple FoundationModel

**Score: 30 / Pass rate: 34% (33/96)**

The Apple on-device model ran via [apfel](https://github.com/Arthur-Ficial/apfel), an OpenAI-compatible wrapper for Apple's FoundationModels framework on macOS 26 Tahoe. It was tested on the same 8 tests × 4 languages, 3 runs each.

### Accuracy breakdown

| Language | Score | Pass Rate |
|----------|-------|-----------|
| Python | **63** | 75% |
| TypeScript | **33** | 38% |
| Go | **12** | 13% |
| C# | **11** | 13% |

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

The model performs adequately in Python (score 63 / 75% pass) — competitive with qwen2.5-coder:7b in that language specifically — and collapses in TypeScript (33 / 38%), Go (12 / 13%), and C# (11 / 13%). Bug fix stands out at score 92 / 92%: the model handles the FizzBuzz correction well across languages, outperforming most Ollama models on that specific test. The unit test writer score (44) is far above its pass rate (8%) — a `ran_clean` partial credit pattern, where code compiles but produces no matching `PASS:` output.

### Resource profile

The most distinctive characteristic of this model is what it doesn't do. During every Ollama model run, GPU utilization was high, fans ran at maximum, and the machine was audibly working. During Apple model runs, neither the GPU nor CPU showed meaningful load — Activity Monitor showed essentially idle behavior on both, and the fans stayed completely off throughout all inference runs.

The observable effect:

- **Battery life** — No GPU saturation means dramatically lower power draw. Sustained Ollama inference on a MacBook Pro drains the battery at GPU-level rates; this model doesn't.
- **Thermals** — Critical for fanless machines. An 8GB MacBook Air will thermal-throttle under sustained Ollama inference; this model has no thermal footprint.
- **Device coverage** — Runs on any Apple Silicon Mac with Apple Intelligence enabled, including 8GB devices where even qwen3.5:4b is impractical.

### Verdict

Not competitive with mid-tier Ollama models for multi-language code generation. Scores below 15 in Go and C# disqualify it as a general coding backend. As a Python-only zero-configuration assistant on constrained hardware — specifically 8GB MacBook Airs where Ollama isn't viable — it has a niche. The thermal profile is real and matters in those contexts.

---

## Speed vs. Quality

Claude generation times were not instrumented; those rows are omitted from this table.

| Model | Mode | Avg time/task | tok/s | Score | Pass Rate |
|-------|------|--------------|-------|-------|-----------|
| `qwen3-coder:30b` | — | ~4s | ~73 | **85** | 88% |
| `qwen3.5:35b-a3b-coding-nvfp4` | nothink | ~5s | ~82 | **83** | 79% |
| `qwen3.6:35b-a3b-coding-nvfp4` | nothink | ~5s | — | **85** | 80% |
| `qwen2.5-coder:7b` | — | ~5s | ~66 | 49 | 53% |
| `apple-foundationmodel` | — | ~7s | ~63 | 30 | 34% |
| `qwen3.5:4b` | nothink | ~8s | ~46 | 38 | 45% |
| `qwen3.5:4b-nvfp4` | nothink | ~8s | ~83 | 36 | 41% |
| `gemma4:26b` | nothink | ~9s | ~68 | **79** | 79% |
| `qwen3.6:35b` | nothink | ~11s | — | **89** | 81% |
| `qwen3.5:35b-a3b-coding-nvfp4` | think | ~12s | ~81 | **88** | 83% |
| `qwen3.5:4b-nvfp4` | think | ~19s | ~84 | 46 | 49% |
| `qwen3.5:27b-nvfp4` | nothink | ~21s | ~20 | 76 | 80% |
| `gemma4:31b` | nothink | ~24s | ~18 | **91** | 85% |
| `qwen3.5:27b` | nothink | ~31s | ~13 | 77 | 80% |
| `gemma4:26b` | think | ~43s | ~69 | 80 | 88% |
| `qwen3.5:4b` | think | ~45s | ~47 | 54 | 66% |
| `qwen3.6:35b-a3b-coding-nvfp4` | think | ~71s | — | **87** | 78% |
| `gemma4:31b` | think | ~79s | ~18 | **93** | 93% |
| `qwen3.6:35b` | think | ~90s | — | **90** | 82% |
| `qwen3.5:27b-nvfp4` | think | ~101s | ~20 | 82 | 86% |
| `qwen3.5:27b` | think | ~106s | ~13 | **86** | 84% |

> tok/s not instrumented for qwen3.6 models in this run.

The speed/quality leaders at the interactive tier:
- **`qwen3-coder:30b`** — score 85 in 4s. The fastest competitive local model.
- **`qwen3.6:35b-a3b-coding-nvfp4` nothink** — score 85 in ~5s. Ties qwen3-coder on score at essentially the same speed; +2 over its qwen3.5 predecessor.
- **`qwen3.5:35b-a3b-coding-nvfp4` nothink** — score 83 in ~5s. The right choice when you want the fast thinking-mode path (~12s, score 88) — qwen3.6:35b-a3b think at ~71s is not an interactive option.
- **`qwen3.6:35b` nothink** — score 89 in ~11s. The highest score of any local nothink model. An 11s latency is easily interactive for most workflows, and the quality gain over qwen3-coder (+4 points) is substantial.
- **`gemma4:26b` nothink** — score 79 in ~9s. Best fit for 36GB machines where the 35B models are tight on memory.
- **`qwen3.5:35b-a3b-coding-nvfp4` think** — score 88 in ~12s. The fastest thinking-mode option; qwen3.6:35b-a3b think is not recommended due to timeout instability and poor ROI.

The qwen3.5:4b-nvfp4 models are notably fast (~83–84 tok/s) due to MLX memory bandwidth — faster token generation than anything else in this table, but the quality floor at 4B limits their usefulness.

---

## Summary and Recommendations

### Top 3 — Non-thinking (interactive use)

These models run without an extended reasoning pass and are suited to interactive sessions where per-turn latency matters.

1. **`qwen3.6:35b` nothink** — score 89, pass 81%, ~11s/task, ~18–20GB RAM. The highest-scoring nothink local model in the benchmark. At ~11s it's slower than the MoE options but well within interactive range, and 4 points above them. Strong for 48GB machines where quality is the priority.

2. **`qwen3.6:35b-a3b-coding-nvfp4` nothink** — score 85, pass 80%, ~5s/task, ~18–20GB RAM. Ties `qwen3-coder:30b` on score at essentially the same speed. Best when you want maximum speed without sacrificing quality. Note: thinking mode for this model is not recommended (poor ROI, timeout instability).

3. **`qwen3-coder:30b`** — score 85, pass 88%, ~4s/task, ~18–20GB RAM. The fastest competitive local model at 100% run consistency. The pass rate edge (88% vs 80%) and perfect consistency make it the safer choice when reliability matters more than score.

**36GB machines:** `gemma4:26b` nothink (score 79, ~9s, ~14–16GB) — the qwen3.6 and MoE models require 18–20GB. `qwen3.5:27b-nvfp4` nothink (score 76, ~21s, ~14–16GB) is the next step up if speed is less critical.

### Top 3 — Thinking (quality-first use)

These models use extended reasoning and produce meaningfully better results at the cost of higher per-turn latency. Best suited for complex tasks, longer sessions, or agentic harness use where latency per turn is less critical.

1. **`gemma4:31b` think** — score 93, pass 93%, ~79s/task, ~18–20GB RAM. The local ceiling. 11/12 on 007 (Beatles interview) — the only local model to approach Claude's perfect 4/4 on that test. Think mode is 97% consistent across runs. The best thinking-mode choice when quality is the priority and latency is acceptable.

2. **`qwen3.5:35b-a3b-coding-nvfp4` think** — score 88, pass 83%, ~12s/task, ~18–20GB RAM. The most time-efficient thinking-mode model in the benchmark. Score 88 at ~12s/task is unmatched — fully interactive for most workflows. 91% consistent. First choice when you want thinking mode without batch latency. (Prefer qwen3.5 over qwen3.6 for the a3b thinking mode — qwen3.6:35b-a3b think shows timeout instability.)

3. **`qwen3.5:27b` think** — score 86, pass 84%, ~106s/task, ~16–18GB RAM. Best option for 36GB machines where the 31B dense model won't fit. 100% run consistency makes it ideal for background/batch agentic use.

### Quick reference by machine

| Machine | Recommendation | Notes |
|---------|---------------|-------|
| 48GB (M3 Max, quality) | `qwen3.6:35b` nothink | Score 89 at ~11s; best nothink quality |
| 48GB (M3 Max, speed) | `qwen3.6:35b-a3b-coding-nvfp4` nothink | Score 85 at ~5s; ties qwen3-coder, +2 over qwen3.5 predecessor |
| 48GB (M3 Max, thinking) | `qwen3.5:35b-a3b-coding-nvfp4` think | Score 88 at ~12s — the only interactive thinking-mode option |
| 48GB (M3 Max, quality-first) | `gemma4:31b` think | Score 93, 11/12 on 007; ~79s/task |
| 36GB (M3 Pro) | `gemma4:26b` nothink | ~14–16GB, score 79 at ~9s; qwen3.5:27b-nvfp4 for thinking-mode fallback |
| 16GB (M3 Air) | `qwen3.5:4b` think | Score 54 with ~45s latency; usable for Python/TypeScript |
| Any Mac (zero setup) | `apple-foundationmodel` | Python only; everything else fails |

### Cloud vs. local gap

All three Claude models score 100 — a clean benchmark sweep at every level of difficulty. The gap to the best local models continues narrowing: `qwen3.6:35b` nothink reaches score 89 — the highest of any local model in nothink mode — and `gemma4:31b` think reaches 93. For routine tasks (001–006, 008), the top local models hit 88–96%+ pass rates and are fully competitive. The remaining gap is concentrated on 007-level reasoning: only gemma4:31b in think mode meaningfully closes it (11/12 passes, pass rate 92% vs Claude's 100%). `qwen3.6:35b` think (1/12 pass) and `qwen3.6:35b-a3b-coding-nvfp4` think (1/12 pass) show the qwen3.6 generation is improving on 007, but not yet at a level comparable to gemma4:31b. For Python- or TypeScript-heavy work with well-defined tasks, the best local models are viable daily-driver alternatives. For tasks requiring the combination of ambiguity tolerance, multi-file schema inference, and precise date arithmetic that 007 exercises, Claude remains the reliable choice.

---

## Real-World Usage: Agentic Harness Potential

This benchmark measures first-shot generation — one prompt, one response, no feedback. A real coding assistant harness (Claude Code, Open Code, Aider, etc.) works differently: the model sees file contents, runs code, reads error output, and iterates over multiple turns. That changes the picture meaningfully.

### How a harness changes the failure profile

The dominant failure category for local models in this benchmark is **compile errors**: Go missing imports, C# `.csx` idiom errors, TypeScript undefined references. In a harness, the model gets the compiler output back and can fix it. A model that generates `./solution.go:4: "strconv" imported and not used` would likely resolve it in one additional tool call. Compile errors that look like model failures here are largely mechanical corrections in an agentic context.

**Logic errors are different** — usually. If the model computes the wrong algorithm and has no way to verify correctness, a harness can't help. Tie-breaking in 002 (word frequency) falls into this category: the model either knows to sort alphabetically on ties or it doesn't, and running the code produces no useful signal either way.

**007 is a different case.** The expected output format (`expected_format.json`) is provided directly in the prompt — the model already has the oracle. In a harness, it could run its code, compare the actual JSON output against the expected format, identify the specific discrepancy (wrong age for John Lennon, incorrect name split), and fix it. The 007 failures in this benchmark are mostly "close but wrong on one field" — exactly the kind of near-miss that iterative tool use is designed to close. 007 may be one of the *best* candidates for harness improvement among the tests here.

The practical implication: Go's 61% and C#'s 58% pass rates here are likely pessimistic for harness use. Python's 84% is probably closer to a real ceiling for tests without a built-in reference — but even that ceiling rises for tests like 007 where the model can self-verify against the provided example.

### Tool calling support

Not all models support structured tool calling (function calling). Support is required to use a model as a harness backend at all. The Apple FoundationModel has no tool calling interface via apfel. All Qwen3/3.5 and Gemma4 models support it natively through Ollama's tool API.

**Consistency** in the table below is the fraction of (test × language) combinations where all 3 independent runs produced the same outcome — all pass or all fail. High consistency means the model's behavior is predictable; low consistency means the same prompt can succeed or fail depending on random variation, which is frustrating in a harness where you're trying to debug whether a failure is a model problem or a code problem.

| Model | Tool Use | Score | Nothink speed | Consistency | Thinking | RAM |
|-------|----------|-------|---------------|-------------|----------|-----|
| `claude-opus-4-6` | ✓ | 100 | — | (1 run) | — | cloud |
| `claude-sonnet-4-6` | ✓ | 100 | — | (1 run) | — | cloud |
| `claude-haiku-4-5` | ✓ | 100 | — | (1 run) | — | cloud |
| `gemma4:31b` | ✓ | 91 | ~24s | 84% | ✓ (93, 97% consistent) | ~18–20GB |
| `qwen3.6:35b` | ✓ | 89 | ~11s | **94%** | ✓ (90, 81% consistent)† | ~18–20GB |
| `qwen3.6:35b-a3b-coding-nvfp4` | ✓ | 85 | ~5s | 91% | ✓ (87, 78% consistent)‡ | ~18–20GB |
| `qwen3.5:35b-a3b-coding-nvfp4` | ✓ | 83 | ~5s | 94% | ✓ (88, 91% consistent) | ~18–20GB |
| `qwen3-coder:30b` | ✓ | 85 | ~4s | **100%** | — | ~18–20GB |
| `qwen3.5:27b` | ✓ | 77 | ~31s | 81% | ✓ (86, 100% consistent) | ~16–18GB |
| `qwen3.5:27b-nvfp4` | ✓ | 76 | ~21s | 84% | ✓ (82, 91% consistent) | ~14–16GB |
| `gemma4:26b` | ✓ | 79 | ~9s | 88% | ✓ (80, 75% consistent) | ~14–16GB |
| `qwen2.5-coder:7b` | ✓ | 49 | ~5s | 81% | — | ~5GB |
| `qwen3.5:4b` | ✓ | 38 | ~8s | 66% | ✓ (54, 59% consistent) | ~3–4GB |
| `qwen3.5:4b-nvfp4` | ✓ | 36 | ~8s | 59% | ✓ (46, 72% consistent) | ~3GB |
| `apple-foundationmodel` | ✗ | 30 | ~7s | 81% | — | ~4GB§ |

> † `qwen3.6:35b` think: 81% consistent but +1 score over nothink — thinking not recommended for harness use.  
> ‡ `qwen3.6:35b-a3b-coding-nvfp4` think: 78% consistent and had 8 timeouts across 96 attempts; not recommended for harness use.  
> § Apple FoundationModel runs on the Neural Engine with near-zero GPU/CPU load. The ~4GB figure is approximate; actual memory impact is significantly lower than equivalent Ollama models.  
> Thinking mode columns show score and consistency for think mode specifically. Nothink speed is the baseline for interactive harness use.

### Considerations for harness use beyond this benchmark

**Instruction following:** A harness issues multi-turn system-level instructions — "edit only this file," "don't add new dependencies," "stop and ask if ambiguous." This benchmark doesn't measure instruction adherence at all. Models that score well here may still struggle with harness-level directive compliance.

**Context window:** Long coding sessions accumulate file contents, tool outputs, and conversation history. Larger context windows directly reduce the frequency of context truncation mid-task. All Qwen3/3.5 and Gemma4 models support 32K+ context via Ollama; qwen3-coder:30b supports 32K.

**Turn latency:** Harness interactions often chain many short tool calls (read file, run linter, apply edit). At 31s/turn, `qwen3.5:27b` nothink is slow but usable. At 106s, think mode is impractical for interactive sessions but workable for background tasks.

### Top 3 models to evaluate in an agentic harness — non-thinking

Best for interactive sessions where per-turn latency matters. The harness's ability to iterate on compile errors reduces the importance of first-shot perfection.

1. **`qwen3.6:35b` nothink** — Score 89, 94% consistent, ~11s/turn. Highest nothink score in the benchmark. The 11s latency is interactive for most session use cases, and the quality advantage (+4 over qwen3-coder) means fewer harness iterations needed to close failures.

2. **`qwen3-coder:30b`** — Score 85, 100% consistent, ~4s/turn. Perfect consistency is especially valuable in a harness: when something fails, you can trust it's the task or environment, not random model variation. Best when speed and predictability matter more than raw score.

3. **`qwen3.6:35b-a3b-coding-nvfp4` nothink** — Score 85, 91% consistent, ~5s/turn. Ties qwen3-coder on score at nearly the same speed; the +2 improvement over `qwen3.5:35b-a3b` is meaningful. Note: do not use think mode in a harness — timeout instability at ~71s/turn with 78% consistency is worse than nothink for iterative sessions.

### Top 3 models to evaluate in an agentic harness — thinking

Best for background or batch sessions, complex multi-file tasks, or cases where quality matters more than latency. Thinking models tend to produce more consistent outputs and handle harder reasoning tasks — both of which compound positively in a multi-turn harness.

1. **`qwen3.5:35b-a3b-coding-nvfp4` think** — Score 88, 91% consistent, ~12s/turn. The only thinking-mode model with interactive latency. Score 88 at ~12s makes it the first thinking-mode option viable for session-by-session use rather than background work only. Top choice when quality matters and you don't want to wait 43–106s per turn.

2. **`gemma4:31b` think** — Score 93, 97% consistent, ~79s/turn. The highest-scoring local model and the only one to pass 007 outright (11/12). 97% consistency is the best of any local model in think mode — nearly every task reliably passes or reliably fails. Best for batch/background agentic tasks where quality is the priority.

3. **`qwen3.5:27b` think** — Score 86, 100% consistent, ~106s/turn. Perfectly consistent — every task either reliably passes or reliably fails, ideal for debugging harness behavior. The 106s/turn is the tradeoff. Best thinking-mode option for 36GB machines where gemma4:31b won't fit.

All models listed support tool calling and score well enough that harness iteration would plausibly close the remaining gap on compile-error failures. None have been validated in an actual agentic harness — these are candidates for follow-on testing, not confirmed replacements.

### Context size for harness use

Ollama's default `num_ctx` (2048–4096) is far too small for agentic sessions where context accumulates across file reads, tool outputs, and conversation history. Custom model variants with larger context should be created via `ollama create` before harness testing. Recommended targets for a 48GB machine:

| Model | Recommended ctx | Notes |
|-------|----------------|-------|
| `qwen3-coder:30b` | 131072 (128K) | MoE KV cache is efficient; ~4–8GB at 128K |
| `qwen3.5:35b-a3b-coding-nvfp4` | 131072 (128K) | MoE KV cache; similar efficiency to qwen3-coder |
| `qwen3.6:35b-a3b-coding-nvfp4` | 131072 (128K) | MoE KV cache; same as qwen3.5:35b-a3b |
| `qwen3.6:35b` | 65536 (64K) | Dense 35B; 128K workable on 48GB but monitor memory |
| `gemma4:31b` | 131072 (128K) | Dense model natively trained at 128K |
| `gemma4:26b` | 131072 (128K) | Natively trained at 128K — no quality degradation |
| `qwen3.5:27b-nvfp4` | 65536 (64K) | Dense attention; 128K workable but monitor memory |

Example Modelfile:
```
FROM qwen3-coder:30b
PARAMETER num_ctx 131072
```

`gemma4` models and the MoE models (`qwen3-coder:30b`, `qwen3.5:35b-a3b`, `qwen3.6:35b-a3b`) are the most confident 128K recommendations: gemma4 because both variants were trained at that length; MoE models because their sparse activation keeps per-token KV cache proportional to active parameters (~3B), not total parameters. The `qwen3.6:35b` dense model can run at 128K on 48GB but will leave less headroom — 64K is safer. Extending qwen3.5 dense models beyond their 32K training window uses RoPE scaling and may degrade quality on very long contexts.
