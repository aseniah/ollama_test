# 006 — Bug Fix

Identify and fix a bug in a provided program, then produce the correct output.

## Task

The model is given a buggy FizzBuzz implementation and must fix it, then output the correct FizzBuzz sequence for `n=15`. Output must be only the sequence — no explanation.

The bug is an off-by-one in the loop range. The source code includes an inline comment hinting at the fix, so the model doesn't need to reason about FizzBuzz from scratch — it just needs to follow the hint and produce working code in its target language.

This test runs in all 4 languages (Python, TypeScript, Go, C#). The model writes the solution in its target language.

## Input

`input/buggy.py` contains the buggy Python implementation. Its content is injected directly into the prompt via the `{source_code}` placeholder — the model sees it as context, not as a file to read. The running solution has no access to the input directory.

## Scoring

| Check | Description |
|---|---|
| `correct_sequence` | output exactly matches the 15-line FizzBuzz sequence |
| `first_line_is_1` | first output line is `1` (catches off-by-one) |
| `last_line_is_fizzbuzz` | 15th line is `FizzBuzz` |

All checks must pass and exit code must be 0.

## What this tests

Code repair and output discipline. The fix is hinted in the source; the challenge is translating the corrected logic into the target language and suppressing any explanatory text.
