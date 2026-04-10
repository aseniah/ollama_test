# Findings Report Instructions

This file guides generation of versioned findings reports (e.g. `FINDINGS_V002.md`).
See `FINDINGS_V001.md` for a complete example of the expected output.
See `CLAUDE.md` for jq query patterns to extract data from the results JSONL files.

---

## Before You Start

1. **Identify the target version.** Results live in `results/v{NNN}/`. List the subdirectories
   to see which models were run: `ls results/v002/` (or whichever version). Note which models
   have results and how many runs each has.

2. **Name the output file** `FINDINGS_V{NNN}.md` to match the results version.

3. **Survey what's present.** Not every version will have the same models, thinking variants,
   quantization variants, or Apple model runs. Adapt the report structure to what was actually
   run — omit sections that don't apply (e.g. no Apple run → no Apple section; no thinking
   variants → no Thinking vs. Non-Thinking section).

4. **Pull all numbers from the data.** Don't estimate or carry over numbers from a prior
   findings doc. Use jq to compute scores and pass rates fresh from the current version's JSONL
   files. CLAUDE.md has the query patterns.

---

## Metrics

Use two metrics throughout. **Score is primary; pass rate is secondary.**

- **Score** — per-test normalized partial credit. For each test, compute the fraction of
  correctness assertions that passed (e.g. 3/4 checks = 75%), then average equally across all
  test/language combinations. Treats every test as equally important and rewards near-misses
  over complete failures.
  Report as a whole number (e.g. `score 79`).

- **Pass rate** — binary reliability. Fraction of tests where the solution fully passed all
  checks end-to-end. Report as a percentage with fraction (e.g. `79% (76/96)`).

When the two diverge, explain what it means: a high score / lower pass rate means near-misses;
a low score / low pass rate means mostly broken output.

---

## Report Structure

### Introduction
Brief description of the benchmark: what it tests, how it works, machine used, scope (N models,
N tests × 4 languages, N runs each — note if any models ran fewer times).

### TL;DR
5–8 bullets covering the headline results. Lead each bullet with score as primary metric.
Include: best local model, surprise performers, thinking vs non-thinking summary (if applicable),
quantization tradeoff (if applicable), Apple model verdict (if run).

### Overall Rankings
Table with all model/mode combinations ranked by score. Columns: Rank, Model, Mode, Score,
Pass Rate, Avg Time/task. List thinking and non-thinking variants separately.

Include a paragraph explaining the two metrics and what divergence between them means. Note
whether Claude models were run with extended thinking or not, and any timing caveats (e.g.
Claude API times include round-trip latency, not directly comparable to on-device Ollama times).

### Model Architectures
Note which models are MoE (Mixture of Experts) vs dense, and which are code-specialized vs
general-purpose. Explain what MoE means for inference cost. This context helps explain why some
models punch above their apparent parameter count.

### Thinking vs. Non-Thinking *(if thinking variants were run)*
Table: Model | Nothink score | Think score | Score delta | Time cost (with multiplier).
Discuss where thinking helps most (smaller models tend to gain more) and any outliers where
thinking hurts.

### Quantization *(if quantized variants were run)*
Table: Model | Standard score | Quantized score | Score delta | Speed delta.
Name the quantization format and explain what it means for the hardware. Discuss the tradeoff
at different model sizes — degradation is typically worse at smaller scales.

### Language Breakdown
Table: Language | Pass Rate | Notes. Then a per-model language breakdown table.
Explain the structural reasons languages underperform — Go's strict compiler surfacing missing
imports, C#'s `.csx` scripting environment requiring specific idioms (`Args[0]`, no NuGet,
built-in CSV/JSON patterns). Call out which models are most language-balanced.

### Test Difficulty
For each test, include:
- Header: `### NNN — Test name — score X / pass Y%` (averaged across all models and runs)
- One-line description of what the test asks for
- Explanation of **what specifically caused failures** — not just the score, but why models
  stumbled. What checks failed? Was it a compile issue, a format issue, a logic edge case?
  Was it language-specific? Did the score/pass gap reveal something about partial credit?
- Language pass rate table: Python | TypeScript | Go | C#

Go deeper than just reporting numbers. Examples of what to explain:
- Which edge case or format requirement trips models
- Whether failures are complete misses or near-misses (score vs pass rate divergence)
- Language-specific failure patterns (e.g. C# failing on JSON deserialization, Go on missing imports)
- Any verifier behavior worth noting (vacuous checks, binary-only tests, etc.)

### Per-Model Profiles
One section per model. Header: `### model-name — score X / pass Y%`
Cover: failure mode breakdown (compile/runtime/logic), language strengths and weaknesses,
consistency across runs (how many test/language combos showed mixed results), speed and memory
requirements, and practical fit (what kind of machine/use case is this model suited for).

### Apple FoundationModel *(if run)*
Separate section. Cover:
- Score and pass rate overall
- Language breakdown table (score + pass rate)
- Test breakdown table (score + pass rate), sorted by score
- Resource profile: this model runs with near-zero GPU and CPU load, fans completely silent —
  unlike Ollama models which drive fans to maximum. Discuss implications for battery life,
  thermals, and device coverage (e.g. 8GB MacBook Air where Ollama isn't practical).
- Verdict: where it's viable, where it isn't

### Speed vs. Quality
Table of all timed models: Model | Mode | Avg time/task | tok/s | Score | Pass Rate.
Highlight the best score-per-second options. Note any timing caveats.

### Summary and Recommendations
- Top 3 local models with one-line rationale each
- Quick-reference table by machine type (e.g. 48GB, 36GB, 16GB)
- Cloud vs. local gap: quantify by score, identify where local models are viable (well-defined
  tasks, Python/TypeScript) vs where they fall short (multi-step reasoning, compiled languages,
  hard edge cases)

---

## Tone and Style

- Technical findings document, not a marketing summary. Be specific.
- Tables for anything with more than 2–3 comparable data points.
- When explaining test failures, name the specific mechanism — wrong sequence start, header text
  breaking exact match, wrong arg idiom, vacuous check passing on empty output, etc.
- Score numbers without a `%` sign; pass rates always with `%` and fraction in parens.
- Avg time values with `~` prefix (e.g. `~9.1s`).
