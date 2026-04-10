# 008 — Prime Numbers

Print all prime numbers up to a CLI-supplied limit.

## Task

The model accepts a single integer `N` as a command-line argument and prints all prime numbers up to and including `N`, one per line.

This test runs in all 4 languages (Python, TypeScript, Go, C#). No input files — just a CLI argument.

## Runtime arguments

The benchmark passes `N=50`, so the expected output is the 15 primes from 2 to 47.

## Scoring

| Check | Description |
|---|---|
| `correct_sequence` | output lines exactly match the 15 primes from 2 to 47 |
| `one_excluded` | `1` is not in the output (1 is not prime) |
| `no_composite` | every number in the output is actually prime |

All checks must pass and exit code must be 0.

## What this tests

CLI argument parsing and primality logic. The `one_excluded` check catches a common off-by-one where models start from 1 rather than 2. The `no_composite` check independently validates correctness even if the sequence check fails due to formatting issues.
