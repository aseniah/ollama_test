# 001 — CSV to JSON

Read a CSV file and convert it to a JSON array, preserving correct types.

## Task

The model is given `input/data.csv` (5 rows, 4 columns: Name, Age, Email, Score) and must output a JSON array to stdout. Each row becomes an object. No other output is allowed.

This test runs in all 4 languages (Python, TypeScript, Go, C#) using the same input file.

## Input

`input/data.csv` is copied into the working directory at runtime. The model reads it relative to the script.

## Scoring

| Check | Description |
|---|---|
| `valid_json` | stdout parses as a JSON array |
| `row_count` | array has exactly 5 elements |
| `age_is_int` | every `Age` field is an integer (not a string) |
| `score_is_float` | every `Score` field is a number |
| `correct_first_name` | first record's `Name` is `"Alice Johnson"` |

All checks must pass and exit code must be 0.

## What this tests

Type awareness — the model must produce correctly typed values (`Age` as int, `Score` as float) rather than treating all fields as strings.
