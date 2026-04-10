# 005 — Unit Test Writer

Write tests for a provided function using a prescribed output format (no test framework).

## Task

The model is given a `calculate_discount(price, quantity)` function and must write a program that tests it. Each test case must print either:

```
PASS: <description>
FAIL: <description> (expected: X, got: Y)
```

The prompt requires at least 4 test cases covering boundary conditions at quantities 9, 10, 49, and 50. The function must be copied inline — not imported from a file.

This test runs in all 4 languages (Python, TypeScript, Go, C#). The model writes the solution in its target language.

## Input

`input/source.py` contains the reference implementation in Python. Its content is injected directly into the prompt via the `{source_code}` placeholder — the model sees it as context, not as a file to read. The running solution has no access to the input directory.

## Scoring

| Check | Description |
|---|---|
| `ran_clean` | exit code is 0 |
| `has_pass_lines` | at least 4 lines starting with `PASS:` |
| `no_fail_lines` | no lines starting with `FAIL:` |

Additionally, `spontaneous_tests` is recorded (but does not affect pass/fail) — it flags whether the model reached for a test framework (pytest, jest, etc.) despite being asked for plain output.

## What this tests

Format compliance is the primary challenge. Models frequently use `PASSED:`, print summaries instead of per-test lines, or use a test framework. The exact `PASS:` / `FAIL:` prefix is required.

## Notes

This test has a high failure rate across all models. The format constraint is intentional — format-following under instruction is part of what's being measured.
