# Code Generation Benchmark — Findings

**Date:** 2026-04-09  
**Machine:** MacBook Pro M3 Max (48GB unified memory)  
**Scope:** 11 model configurations, 8 tests × 4 languages, 3 runs each (Claude models: 1 run each)

---

## Introduction

This benchmark evaluates how well various AI models generate correct, executable code across four languages (Python, TypeScript, Go, C#) and eight tasks ranging from simple algorithms to complex CSV transformation. The goal is to understand which local Ollama models are viable coding assistant alternatives to cloud-hosted Claude models, and whether tradeoffs like quantization or extended thinking modes are worth it in practice.

Each model generates a solution in response to a task prompt. The solution is executed, and a per-test verifier checks correctness. Results capture both a binary pass/fail and a partial-credit check score.

---

## TL;DR

**Claude models set the ceiling:** Opus at 97%, Sonnet at 94%, Haiku at 84%.

**Best local model: `qwen3-coder:30b`** at 79% — the only Ollama model that got close to Haiku, and it did so with the fastest generation time of any local model tested (~3.6s avg). It's a Mixture of Experts (MoE) model, giving it 30B-class knowledge at dense-3B inference cost. Requires ~18–20GB RAM.

**`gemma4:26b` is the surprise:** 78% overall, with nothink mode alone hitting 79% — essentially tied with qwen3-coder. It's the most language-balanced local model, requires roughly 14–16GB, and is meaningfully faster than the qwen3.5 27B family. A strong second choice.

**Thinking mode helps consistently**, but the time cost is large (2–6× slower). For smaller models (4B), thinking jumps the score from 39 to 62 (+23 points) and is worth the wait. For larger models (27B), thinking adds modest score gains (+7–10) for a 5× time penalty.

**nvfp4 quantization is an MLX-based format** optimized for Apple Silicon's unified memory architecture. It enables lower memory footprint and faster token generation — but at a meaningful accuracy cost, especially at 4B scale where it drops scores from 39 to 29 (nothink) and pass rates from 34% to 18%. Worth considering at 27B if you're memory-constrained; avoid at 4B.

**Apple's FoundationModel underdelivers at score 39 / 28% pass rate**, and runs in a completely different resource profile: inference barely registered on either the GPU or CPU during testing, with no fan activity — unlike every Ollama model, which drove fans to maximum. This efficiency is real but doesn't compensate for the accuracy gap outside of Python.

---

## Overall Rankings

Each thinking/non-thinking variant is listed separately since they represent meaningfully different configurations. Claude models don't have a thinking field in results (they use standard API mode).

Two metrics are reported for each model:

- **Score** — a nuanced grade of overall capability. Each test is scored as the fraction of its correctness assertions that passed (e.g. a solution that passes 3 of 4 checks contributes 75% for that test), then averaged equally across all tests. This treats every test as equally important and rewards near-misses rather than treating them as equivalent to complete failures.
- **Pass rate** — binary reliability. What fraction of tests did the solution fully pass end-to-end? This answers the practical question: how often will it actually work?

The two numbers tell different stories. A model with a high score but lower pass rate (Sonnet: 99 / 94%) is a capable model whose failures are near-misses — almost right, wrong on one detail. A model with a low score and low pass rate (qwen3.5:4b-nvfp4 nothink: 29 / 18%) is producing mostly broken output.

This is also why Sonnet ranks above Opus by score despite a lower pass rate. Opus passes one more test outright (31 vs 30), but its single failure misses several checks — a harder stumble. Sonnet's two failures are both near-misses that clear almost all their checks, pulling its average score above Opus's. By pass rate, Opus wins; by capability, they're effectively tied with Sonnet slightly ahead on the near-misses metric.

| Rank | Model | Mode | Score | Pass Rate | Avg Time/task |
|------|-------|------|-------|-----------|--------------|
| 1 | `claude-sonnet-4-6` | — | **99** | 94% (30/32) | — |
| 2 | `claude-opus-4-6` | — | **97** | 97% (31/32) | — |
| 3 | `claude-haiku-4-5` | — | **86** | 84% (27/32) | — |
| 4 | `gemma4:26b` | nothink | **81** | 79% (76/96) | ~9.1s |
| 5 | `qwen3-coder:30b` | — | **80** | 79% (76/96) | ~3.6s |
| 6 | `qwen3.5:27b` | think | **79** | 77% (74/96) | ~140s |
| 7 | `gemma4:26b` | think | **78** | 76% (73/96) | ~41.9s |
| 8 | `qwen3.5:27b-nvfp4` | think | **72** | 71% (68/96) | ~124s |
| 9 | `qwen3.5:27b` | nothink | **69** | 68% (65/96) | ~27.2s |
| 10 | `qwen3.5:27b-nvfp4` | nothink | **65** | 65% (62/96) | ~20.8s |
| 11 | `qwen3.5:4b` | think | **62** | 59% (57/96) | ~40.1s |
| 12 | `qwen2.5-coder:7b` | — | **58** | 51% (49/96) | ~4.4s |
| 13 | `qwen3.5:4b-nvfp4` | think | **47** | 41% (39/96) | ~30.6s |
| 14 | `apple-foundationmodel` | — | **39** | 28% (27/96) | ~6.9s |
| 14 | `qwen3.5:4b` | nothink | **39** | 34% (33/96) | ~8.6s |
| 16 | `qwen3.5:4b-nvfp4` | nothink | **29** | 18% (17/96) | ~5.2s |

> **Avg time** is average wall clock time per test/language combination from generation start to result. Codegen requires 100–500+ output tokens, making these times much longer than what you might expect from a simple single-turn prompt. Claude timing was not instrumented in this run (all zeros were recorded); subsequent runs will capture it. Claude times will include API round-trip latency and are not directly comparable to Ollama on-device measurements.
>
> **Thinking column:** Claude models run standard inference — no extended thinking enabled. This is equivalent to "nothink" mode on the Ollama models.

---

## Model Architectures

Not all models are architecturally equivalent. A few distinctions worth knowing:

**Mixture of Experts (MoE):** Rather than activating all parameters for every token, MoE models route each token through a small subset of "expert" sub-networks. This means a large total parameter count with far fewer active parameters per inference pass — delivering large-model knowledge at small-model inference cost.

- **`qwen3-coder:30b`** — MoE: 30B total parameters, ~3.3B active per token. This explains its combination of high pass rate and fast generation speed.

**Code-specialized models:** Some models in this benchmark are general-purpose; others were specifically fine-tuned on code.

- **`qwen2.5-coder:7b`** — Part of the Qwen 2.5 Coder series, trained specifically on code corpora. Despite this specialization, its score 58 / 51% pass rate is the weakest among larger-context Ollama models, showing that code-specific fine-tuning at 7B doesn't overcome the parameter count ceiling for multi-step codegen tasks.
- **`qwen3-coder:30b`** — Also code-specialized, but at 30B scale (with MoE efficiency). The combination works; it's the top local performer.

**General-purpose models evaluated for coding:**

- **`gemma4:26b`** — Google's Gemma 4 family. Tested here as a general-purpose model for coding suitability; performs surprisingly well.
- **`qwen3.5` family** — General-purpose models from the Qwen 3.5 generation with optional thinking mode.

---

## Thinking vs. Non-Thinking

Models that support a thinking/reasoning mode were tested in both modes. In every case, thinking improved capability — sometimes dramatically.

| Model | Nothink | Think | Score delta | Time cost |
|-------|---------|-------|-------------|-----------|
| `qwen3.5:4b` | 39 | 62 | **+23** | ~8.6s → ~40.1s (4.6×) |
| `qwen3.5:4b-nvfp4` | 29 | 47 | **+18** | ~5.2s → ~30.6s (5.9×) |
| `qwen3.5:27b` | 69 | 79 | **+10** | ~27.2s → ~140s (5.2×) |
| `qwen3.5:27b-nvfp4` | 65 | 72 | **+7** | ~20.8s → ~124s (6.0×) |
| `gemma4:26b` | 81 | 78 | **-3** | ~9.1s → ~41.9s (4.6×) |

**Key observations:**

- For `qwen3.5:4b`, thinking mode jumps the score by 23 points. At only 4B parameters, the model benefits enormously from the additional reasoning pass — it effectively becomes a different tier of model.
- The gains shrink as the base model gets larger. `qwen3.5:27b` gains 10 points for a 5× time cost.
- `gemma4:26b` is the outlier: thinking mode slightly *hurts* performance (-3) while quadrupling generation time. Gemma4 appears to reason best in a single pass. Non-thinking is the right mode for this model.

**Recommendation:** Use thinking mode for `qwen3.5:4b` — the quality gain is large enough to justify the wait. Skip thinking for `gemma4:26b`. For `qwen3.5:27b`, thinking mode is worth it only in batch or non-interactive contexts where you can afford ~140s per task.

---

## Quantization: nvfp4 vs Standard

The qwen3.5 models were tested in both their standard GGUF form and as `nvfp4` variants. The nvfp4 format is an **MLX-native quantization** developed for Apple Silicon's unified memory architecture. Unlike GGUF (which was designed for CPU/GPU split inference), MLX models operate natively in Apple's unified memory space, enabling more efficient memory use and faster token generation on M-series hardware.

The tradeoff: faster and lighter, but with quality loss.

| Model | Standard | nvfp4 | Score delta | Speed delta |
|-------|----------|-------|-------------|-------------|
| `qwen3.5:27b` nothink | 69 | 65 | -4 | ~27.2s → ~20.8s (24% faster) |
| `qwen3.5:27b` think | 79 | 72 | -7 | ~140s → ~124s (11% faster) |
| `qwen3.5:4b` nothink | 39 | 29 | **-10** | ~8.6s → ~5.2s (40% faster) |
| `qwen3.5:4b` think | 62 | 47 | **-15** | ~40.1s → ~30.6s (24% faster) |

**At 27B**, nvfp4 is a defensible tradeoff: a score drop of 4–7 points for notably faster generation and lower memory footprint. If you're on a machine where the full 27B model is tight on memory, the nvfp4 variant may be the practical choice.

**At 4B**, nvfp4 causes severe degradation — a score drop of 10–15 points depending on mode. The model is small enough that 4-bit quantization removes too much representational capacity. Not recommended.

---

## Language Breakdown

Pass rates aggregated across all models and runs.

| Language | Pass Rate | Notes |
|----------|-----------|-------|
| Python | **68%** | Highest overall; most models are well-calibrated here |
| TypeScript | **64%** | Strong but more variance across models |
| Go | **56%** | Compiler strictness reliably surfaces missing imports |
| C# | **52%** | `.csx` scripting argument access trips most models |

Go and C# consistently underperform Python. Go failures cluster on compile errors from missing or incorrect imports — Go's strict compiler surfaces issues that Python or TypeScript would let slide at runtime. C# issues are largely the `.csx` script runtime argument handling: models default to `args[0]` or `Environment.GetCommandLineArgs()[1]`, both wrong; the correct form for dotnet-script is `Args[0]`. C# also struggles with JSON deserialization idioms in script mode (004 JSON filter) and CSV/file handling patterns that differ from full .NET project conventions.

**Per model, language strengths vary:**

| Model | Python | TypeScript | Go | C# |
|-------|--------|-----------|----|----|
| `claude-opus-4-6` | 100% | 100% | 100% | 88% |
| `claude-sonnet-4-6` | 100% | 100% | 88% | 88% |
| `claude-haiku-4-5` | 88% | 88% | 88% | 75% |
| `qwen3-coder:30b` | 79% | 83% | 71% | 83% |
| `gemma4:26b` | 75% | 81% | 73% | **81%** |
| `qwen3.5:27b` | 75% | 77% | **79%** | 58% |
| `qwen3.5:27b-nvfp4` | 67% | 75% | 75% | 54% |
| `qwen2.5-coder:7b` | 67% | 54% | 33% | 50% |
| `qwen3.5:4b` | 52% | 56% | 42% | 38% |
| `apple-foundationmodel` | **75%** | 25% | 8% | 4% |

Notable patterns:
- `gemma4:26b` is the most **language-balanced** local model — no language falls below 73%. The most consistent breadth of any Ollama model.
- `qwen3.5:27b` has a relative strength in **Go** (79%), unusual for a general model.
- `qwen3-coder:30b` is especially strong at **C# and TypeScript** (both 83%) relative to its overall rate, suggesting the coding-specific training helps with typed/compiled languages.
- `apple-foundationmodel` shows a striking cliff: 75% in Python but 25% TypeScript, 8% Go, 4% C#. It appears heavily biased toward Python and struggles badly with compiled or typed languages.

---

## Test Difficulty

Score and pass rate across all models for each test. Score reflects partial credit (checks passed / total checks per test, averaged); pass rate is binary end-to-end success. Where the two diverge, score reveals how close models are getting even when they don't fully pass.

### 008 — Prime numbers — score **94** / pass 82%

> Given integer N as a CLI argument, print all prime numbers up to N, one per line.

The largest score/pass gap in the benchmark (+12 points), and the source of this gap is specific: the verifier checks three things — `correct_sequence`, `one_excluded` (1 is not prime), and `no_composite`. The latter two both evaluate to vacuously true on empty output, so every compile failure (Go/C#, missing imports) earns 2/3 checks despite producing nothing. Among logic failures, models almost universally know the prime algorithm — the stumble is formatting: several produce a header line like `Prime numbers up to and including 50:` that breaks the exact `correct_sequence` match. The algorithm itself is rarely wrong.

Python is nearly universal at 98%; Go the weakest at 67%, almost entirely compile errors on missing import declarations.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 98% | 83% | 67% | 81% |

---

### 005 — Unit test writer — score **84** / pass 76%

> Given a function implementation (in the target language), write unit tests for it in plain output format: each test prints `PASS: <description>` or `FAIL: <description>` to stdout.

The function under test is a discount calculator with tiered pricing logic. Models must write tests that run the function and print `PASS: <description>` for each passing assertion — with zero `FAIL:` lines (the function is correct; all tests should pass). The verifier checks three things: ran clean, has 4+ `PASS:` lines, and has no `FAIL:` lines.

Two failure modes dominate. First: models reach for testing frameworks (pytest, jest, `testing.T`, NUnit) instead of the custom format — the tests run cleanly but produce no `PASS:` output matching the expected pattern. Second: models compute wrong expected values for the discount tiers, causing their tests to output `FAIL:` lines because the function (correctly) disagrees with their expected results. The score/pass gap (+8) comes from `ran_clean` being a free check point in nearly all failures — almost every model at least produces code that executes.

Language scores are unusually even, confirming this is a format/logic comprehension problem independent of language skill.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 76% | 81% | 71% | 74% |

---

### 002 — Word frequency — score **75** / pass 74%

> Read a text file, count word frequency (case-insensitive), and print the top 10 words sorted by frequency descending, then alphabetically for ties.

The word counting itself is usually not the problem — models handle case folding and frequency correctly. The trip point is tie-breaking: the test data has three words tied at 10 occurrences (`it`, `of`, `was`) and they must be sorted alphabetically when frequency is equal. Models that sort only by frequency (or sort all words alphabetically regardless of frequency) fail the `ties_sorted` check. When they fail, they tend to fail all five checks together — it's not a partial miss, it's a complete format or logic mismatch. Python leads at 86%; C# and Go trail primarily due to compile errors and output format differences.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 86% | 74% | 69% | 67% |

---

### 003 — Fibonacci — score **71** / pass 71%

> Accept integer N as a CLI argument. Print all Fibonacci numbers up to N, one per line. Sequence starts: 1, 1, 2, 3, 5...

Score equals pass rate — both checks (`correct_sequence` and `no_extra_output`) fail together when anything is wrong, so there's no partial credit. Compile errors dominate: 23 of 48 total failures are Go or C# compile errors (Go missing imports, C# `.csx` arg handling), accounting for most of the shortfall.

Among logic failures, the sequence itself is the stumble. The expected output starts `1, 1, 2, 3, 5...` with two leading 1s. Common wrong outputs observed: starting at 0 (mathematical convention vs. the prompt's explicit specification), outputting only one leading 1 and producing `1, 2, 3, 5...`, or repeating extra 1s at the start (`1, 1, 1, 2...`). C# failures also include truncation due to wrong arg parsing — models use the script filename instead of the integer argument. Python leads at 86%; Go's 62% is nearly all compile errors.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 86% | 74% | 62% | 64% |

---

### 004 — JSON filter — score **70** / pass 70%

> Read a JSON array of objects with fields (name, age, active, score). Output only records where `active == true` and `score >= threshold`, sorted by score descending.

Score equals pass rate — failures are complete misses with no partial credit. Python is near-perfect at 95%: list comprehensions and `json` module make this near-trivial. Go (69%) and TypeScript (83%) failures are mostly compile or runtime issues. The biggest drag is C# at 31%: reading a JSON array from a file, deserializing it to typed objects, filtering with LINQ, and re-serializing in `.csx` script mode is an uncommon enough pattern that most models default to approaches that work in full .NET projects but fail in the script environment. Models that pass do so by using `dynamic` or `JsonDocument` — models that fail typically try to define classes inline or use library features unavailable in script mode.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 95% | 83% | 69% | 31% |

---

### 001 — CSV to JSON — score **67** / pass 65%

> Read a CSV file and output it as a JSON array, with correct type coercion: age as integer, score as float.

Score is slightly above pass rate — some failures pass structural checks (valid JSON, correct row count) but fail on type coercion. The type coercion requirement is the specific trip point: models often parse the CSV correctly and produce valid JSON but output all values as strings, failing `age_is_int` and `score_is_float`. This pattern is most visible in Go and TypeScript logic failures, where the structure is right but types are wrong. Python at 79% handles type inference more naturally; C# at 48% struggles with both type coercion and CSV handling patterns in the `.csx` scripting environment.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 79% | 74% | 60% | 48% |

---

### 006 — Bug fix — score **38** / pass 38%

> A FizzBuzz implementation with a range bug (starts at 0 instead of 1, uses `< n` instead of `<= n`). Find and fix the bug; output only the corrected FizzBuzz sequence for n=15.

Score equals pass rate exactly — the only test where this is true. Bug fix is perfectly binary: models either produce the correct sequence or they don't; there are no partial checks to earn. You get it right or you get nothing.

**Python is an outlier at 17% pass**, while TypeScript (40%), Go (45%), and C# (50%) score higher. Nearly every Ollama model gets 0% on the Python version specifically, including models that pass it in all other languages. Claude models and the Apple model both pass Python bug fix 100%. The leading hypothesis is that the Python buggy code includes an inline comment (`# Bug: should be range(1, n + 1)`) that confuses Ollama models — either causing them to produce extra commentary, make unrelated changes, or echo the hint without genuinely fixing the logic. Compiled languages force a simpler code-only response pattern.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 17% | 40% | 45% | 50% |

---

### 007 — Beatles interview — score **9** / pass 3%

> Read a CSV of Beatles members (name, born, died). Using an example JSON file to infer the transformation rules, produce a JSON array with each member's name and age (at death for deceased, current age for living).

The benchmark's hardest test by design. It requires: parsing a CSV with mixed present/absent death dates, inferring transformation logic from an example JSON, computing ages correctly with date arithmetic (handling living vs. deceased members differently), and producing output that matches the expected JSON structure exactly. Only 5 records fully passed across 168 total attempts — all from Claude models. No Ollama model passed a single run.

Score is 3× higher than pass rate (9 vs 3%), revealing that many models produce output with some structural validity — valid JSON, correct member names — while failing on the harder checks like age calculation and date formatting. The per-model score breakdown exposes a clear capability hierarchy even within "everyone failed":

| Model | Mode | Score | Passes |
|-------|------|-------|--------|
| `claude-sonnet-4-6` | — | **89** | 2/4 |
| `claude-opus-4-6` | — | **75** | 3/4 |
| `claude-haiku-4-5` | — | 14 | 0/4 |
| `gemma4:26b` | nothink | 14 | 0/12 |
| `gemma4:26b` | think | 11 | 0/12 |
| `qwen3.5:27b` | nothink | 9 | 0/12 |
| `qwen3.5:27b` | think | 9 | 0/12 |
| `qwen3-coder:30b` | — | 7 | 0/12 |
| `qwen3.5:27b-nvfp4` | nothink/think | 5 | 0/12 |
| `qwen3.5:4b` | nothink | 4 | 0/12 |
| `qwen2.5-coder:7b` | — | 3 | 0/12 |
| `qwen3.5:4b` | think | 2 | 0/12 |
| `qwen3.5:4b-nvfp4` | think | 2 | 0/12 |
| `apple-foundationmodel` | — | 1 | 0/12 |
| `qwen3.5:4b-nvfp4` | nothink | **0** | 0/12 |

Sonnet scores 89 despite only 2/4 passes — its two failures are near-misses clearing most checks. Haiku and gemma4:26b nothink both score 14 with zero passes, meaning they're producing valid JSON with some correct fields but failing on the computed values. At the bottom, `qwen3.5:4b-nvfp4` nothink scores 0 — not a single check passed across all 12 runs.

| Python | TypeScript | Go | C# |
|--------|-----------|----|----|
| 5% | 5% | 2% | 0% |

---

## Per-Model Profiles

### claude-opus-4-6 — score 97 / pass 97%
Near-perfect on both metrics. One failure (C#, likely the `.csx` args pattern); logic and compile errors are essentially absent. Only model to score 100% in Go. The reference ceiling for this benchmark.

### claude-sonnet-4-6 — score 99 / pass 94%
Ranks first by score despite two more failures than Opus — those failures are near-misses that clear almost all their checks, pulling the score above Opus's. No compile errors across any run, the only model to achieve this. The score gap between Sonnet (99) and Opus (97) reflects capability; the pass rate gap (94% vs 97%) reflects that Sonnet technically stumbles on two more tests end-to-end.

### claude-haiku-4-5 — score 86 / pass 84%
Score and pass rate are close, meaning Haiku's failures tend to be more complete misses than near-misses. Five failures overall: mostly runtime and C#-related. The weakest Claude on C# (75%) but still ahead of all local models. A viable API option when cost matters.

### qwen3-coder:30b — score 80 / pass 79%
**Best local model.** MoE architecture (30B total, ~3.3B active per token) delivers 30B-class knowledge at dense-3B inference speed: ~3.6s average per codegen task, ~85 tok/s. Failure modes are well-distributed (7 compile, 7 runtime, 5 logic). Strong in C# and TypeScript (both 83%). Inconsistency across runs is low (4/32 combos showed mixed results). Requires ~18–20GB RAM — comfortable on 48GB, tight on 36GB, not feasible on 16GB.

### gemma4:26b — score 81 nothink / 78 think — pass 79% / 76%
**Second best local model and the most balanced.** Edges qwen3-coder by one score point in nothink mode (81 vs 80). Thinking mode marginally hurts both score and pass rate — use nothink. Most language-balanced local model tested, with no language falling below 73%. Generation time in nothink mode (~9s) is far more practical than the qwen3.5:27b family. Memory footprint approximately 14–16GB — lighter than qwen3-coder:30b and viable on 36GB machines. Inconsistency is low (3/32 combos with mixed results across runs).

### qwen3.5:27b — score 79 think / 69 nothink — pass 77% / 68%
Thinking mode brings a meaningful score jump (+10) but at ~140s per task — impractical for interactive use. Nothink at score 69 is weaker than both qwen3-coder:30b and gemma4:26b nothink. Score and pass rate track closely in both modes, meaning failures tend to be complete misses. More inconsistency than gemma4: 6/32 combos have mixed results across runs in thinking mode. Best suited to batch or async workflows.

### qwen3.5:27b-nvfp4 — score 72 think / 65 nothink — pass 71% / 65%
MLX quantization gives a speed advantage (~20.8s vs ~27.2s nothink) at a score cost of 7 points in thinking mode, 4 in nothink. At 27B, the gap is tolerable. Worth considering if memory is the constraint; otherwise the standard weights are better.

### qwen2.5-coder:7b — score 58 / pass 51%
Score noticeably above pass rate (58 vs 51%), suggesting it's getting partway through tests it doesn't fully pass — earning structural or partial checks even on failures. Codegen demands longer reasoning chains where 7B parameters hit a ceiling. Go is a notable weak spot (33%). Compile errors are the dominant failure mode. The primary advantage is size and accessibility: at ~4.4s and 4.7GB, it runs on any machine.

### qwen3.5:4b — score 62 think / 39 nothink — pass 59% / 34%
A tale of two modes. Nothink score of 39 is poor; thinking jumps it to 62 (+23 points). The gap between score and pass rate is larger in nothink mode (39 vs 34%) than think mode (62 vs 59%), suggesting nothink failures earn more partial credit while thinking mode failures are more often complete misses when they do fail. Highly inconsistent across runs (14/32 combos show mixed results in thinking mode). Best treated as a fallback for memory-constrained machines.

### qwen3.5:4b-nvfp4 — score 47 think / 29 nothink — pass 41% / 18%
Not recommended. The score/pass gap in nothink mode (29 vs 18%) is the widest of any model — it's earning partial credit on roughly a third of its failures but almost never crossing the finish line. Even thinking mode at score 47 underperforms standard qwen3.5:4b nothink (score 39) on a per-point basis relative to time cost. The most inconsistent model tested (18/32 combos with mixed results in thinking mode). The memory savings don't justify the accuracy loss at this size.

---

## Apple FoundationModel

**Score: 39 / Pass rate: 28% (27/96)**

The Apple on-device model ran via [apfel](https://github.com/Arthur-Ficial/apfel), an OpenAI-compatible wrapper for Apple's FoundationModels framework on macOS 26 Tahoe. It was tested on the same 8 tests × 4 languages, 3 runs each.

### Accuracy breakdown

| Language | Score | Pass Rate |
|----------|-------|-----------|
| Python | **79** | 75% |
| TypeScript | **37** | 25% |
| Go | **21** | 8% |
| C# | **18** | 4% |

| Test | Score | Pass Rate |
|------|-------|-----------|
| 006 — Bug fix | **75** | 75% |
| 008 — Prime numbers | **78** | 33% |
| 005 — Unit test writer | **42** | 8% |
| 003 — Fibonacci | **33** | 33% |
| 002 — Word frequency | **30** | 25% |
| 001 — CSV to JSON | **27** | 25% |
| 004 — JSON filter | **27** | 25% |
| 007 — Beatles interview | **1** | 0% |

The model performs adequately in Python (score 79 / 75% pass) — on par with qwen2.5-coder:7b — but collapses in TypeScript (score 37 / 25%), Go (score 21 / 8%), and especially C# (score 18 / 4%). This is consistent across all three runs. The bug fix result (score 75 / 75% pass) is a notable bright spot, outperforming most Ollama models on that specific test — it handles the Python case cleanly where most Ollama models fail.

The prime numbers score (78) is conspicuously higher than its pass rate (33%): this mirrors the vacuous-check phenomenon noted in the test breakdown — empty output on compile failures earns 2/3 prime-number checks. The unit test writer score (42) is far above its pass rate (8%) for the same reason: every run that executes cleanly earns 1/3 checks regardless of output quality.

### Resource profile

The most distinctive characteristic of the Apple model is how little it asks of the hardware. During all Ollama model runs, GPU utilization was high and fans ran at maximum. During Apple model runs, neither the GPU nor CPU showed meaningful load — Activity Monitor showed essentially idle behavior on both, and the fans stayed off entirely throughout.

Whether this reflects Neural Engine inference, a particularly efficient on-device runtime, or simply output generation fast enough that GPU utilization never has time to register is unclear. What's observable is that the model runs without thermal impact: no fan noise, no heat output, no visible load spikes.

This has practical implications:

- **Battery life** — No GPU saturation means dramatically lower power draw during use.
- **Thermals** — Critical for fanless machines (MacBook Air). Sustained Ollama inference will thermal-throttle an Air; the Apple model won't.
- **Broad device coverage** — Works on any Apple Silicon Mac with Apple Intelligence enabled, including 8GB MacBook Airs where Ollama models aren't practical.

### Verdict

Not competitive with mid-tier Ollama models for general-purpose code generation. Score 39 overall, with near-zero scores in Go (21) and C# (18), disqualifies it as a primary coding backend. It could serve as a zero-configuration fallback for Python-focused users on constrained hardware, particularly on 8GB MacBook Airs where thermal and memory constraints rule out Ollama entirely.

---

## Speed vs. Quality

For models where timing data is available:

| Model | Mode | Avg time/task | tok/s | Score | Pass Rate |
|-------|------|--------------|-------|-------|-----------|
| `qwen3-coder:30b` | — | ~3.6s | ~85 | **80** | 79% |
| `qwen2.5-coder:7b` | — | ~4.4s | ~73 | 58 | 51% |
| `qwen3.5:4b-nvfp4` | nothink | ~5.2s | ~85 | 29 | 18% |
| `apple-foundationmodel` | — | ~6.9s | ~65 | 39 | 28% |
| `qwen3.5:4b` | nothink | ~8.6s | ~46 | 39 | 34% |
| `gemma4:26b` | nothink | ~9.1s | ~64 | **81** | 79% |
| `qwen3.5:27b-nvfp4` | nothink | ~20.8s | ~20 | 65 | 65% |
| `qwen3.5:27b` | nothink | ~27.2s | ~13 | 69 | 68% |
| `qwen3.5:4b-nvfp4` | think | ~30.6s | — | 47 | 41% |
| `qwen3.5:4b` | think | ~40.1s | — | 62 | 59% |
| `gemma4:26b` | think | ~41.9s | — | 78 | 76% |
| `qwen3.5:27b-nvfp4` | think | ~124s | — | 72 | 71% |
| `qwen3.5:27b` | think | ~140s | — | 79 | 77% |

`qwen3-coder:30b` and `gemma4:26b` (nothink) are the clear leaders on score per second: both score 80/81 while the next best is 79 at 140 seconds per task. qwen3-coder:30b is faster; gemma4:26b is lighter on memory.

---

## Summary and Recommendations

### Top 3 local Ollama models for code generation

1. **`qwen3-coder:30b`** — score 80, pass rate 79%, ~3.6s/task, ~18–20GB RAM. Fastest local model, strong across all languages. First choice for 48GB machines.

2. **`gemma4:26b` (nothink)** — score 81, pass rate 79%, ~9.1s/task, ~14–16GB RAM. Marginally higher score than qwen3-coder, most language-balanced, lighter on memory. Better fit for 36GB machines where qwen3-coder is tight.

3. **`qwen3.5:27b` (think)** — score 79, pass rate 77%, ~140s/task. High capability ceiling but only practical for batch or background use. Skip if you need interactive response times.

### Quick reference by machine

| Machine | Recommendation | Notes |
|---------|---------------|-------|
| 48GB (M3 Max) | `qwen3-coder:30b` | Best quality + fastest local speed |
| 36GB (M3 Pro) | `gemma4:26b` nothink | qwen3-coder pushes memory; gemma4 lighter + tied quality |
| 16GB (M3 Air) | `qwen3.5:4b` think | 59% with patience; nothing great at this tier |
| Any Mac (zero setup) | `apple-foundationmodel` | Python only; everything else fails |

### Cloud vs. local gap

By score, the best local models (80–81) sit roughly 5 points below Haiku (86) and 18 below Sonnet (99). By pass rate, the gap is similar: best local at 79% vs Haiku at 84%, Sonnet at 94%. For tasks requiring high reliability across all languages, cloud Claude is significantly ahead. For Python-heavy or TypeScript-heavy work with acceptable failure rates, local models are viable — particularly on well-defined tasks like prime numbers and unit test writing where local models approach 80–90% pass rates.

The gap is most pronounced on multi-step reasoning under ambiguity (007 Beatles interview), where local models score 0% pass rate and Claude models score reliably. Bug fix in Python (17% pass rate for local models vs 100% for Claude) is another sharp dividing line — one where the score metric understates the gap, since bug fix is perfectly binary with no partial credit.
