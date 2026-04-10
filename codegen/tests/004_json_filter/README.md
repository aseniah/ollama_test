# 004 — JSON Filter

Filter and sort a JSON array based on multiple field conditions.

## Task

The model reads `input/data.json` — an array of objects with fields `name`, `age`, `active`, and `score` — and outputs a JSON array containing only records where `active` is `true` and `age >= 30`, sorted by `name` ascending.

This test runs in all 4 languages (Python, TypeScript, Go, C#) using the same input file.

## Input

`input/data.json` is copied into the working directory at runtime. The input contains a mix of active/inactive and young/old records to test filtering logic.

## Scoring

| Check | Description |
|---|---|
| `valid_json` | stdout parses as a JSON array |
| `correct_count` | exactly 4 records in output |
| `correct_order` | names are `["Alice", "David", "Eva", "Grace"]` in that order |
| `bob_excluded` | Bob (inactive) is not in the result |
| `carol_excluded` | Carol (age < 30) is not in the result |

All checks must pass and exit code must be 0.

## What this tests

Multi-condition filtering and sort correctness. Bob and Carol are the two common wrong-answer traps — one tests the `active` filter, the other tests the `age` filter.
