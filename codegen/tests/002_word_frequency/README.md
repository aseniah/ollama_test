# 002 — Word Frequency

Count word frequencies in a text file, with normalization and deterministic sort order.

## Task

The model reads `input/text.txt`, counts how often each word appears (lowercased, punctuation stripped), and outputs one `word: count` line per unique word. Output is sorted by count descending; ties broken alphabetically ascending.

This test runs in all 4 languages (Python, TypeScript, Go, C#) using the same input file.

## Input

`input/text.txt` is copied into the working directory at runtime.

## Scoring

| Check | Description |
|---|---|
| `valid_format` | every line matches `word: count` |
| `top_word` | first line is `the: 11` |
| `it_count` | `it: 10` appears in output |
| `of_count` | `of: 10` appears in output |
| `ties_sorted` | among count-10 words, order is `it` → `of` → `was` (alphabetical) |

All checks must pass and exit code must be 0.

## What this tests

Text normalization and sort stability — the model must handle lowercasing, punctuation stripping, and a two-key sort (count desc, word asc) correctly.
