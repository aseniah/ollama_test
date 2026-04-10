# 003 — Fibonacci Sequence

Print all Fibonacci numbers up to a CLI-supplied limit.

## Task

The model accepts a single integer `N` as a command-line argument and prints all Fibonacci numbers up to and including the largest one that does not exceed `N`, one per line. Sequence starts `1, 1, 2, 3, 5, ...`

This test runs in all 4 languages (Python, TypeScript, Go, C#). No input files — just a CLI argument.

## Runtime arguments

The benchmark passes `N=100`, so the expected output is the first 11 Fibonacci numbers (1 through 89).

## Scoring

| Check | Description |
|---|---|
| `correct_sequence` | output lines match `["1","1","2","3","5","8","13","21","34","55","89"]` exactly |
| `no_extra_output` | no additional lines beyond the sequence |

All checks must pass and exit code must be 0.

## What this tests

CLI argument parsing and basic sequence generation. The double `1` at the start is a common stumbling point.
