# 007 — Beatles Interview

The hardest test. Transform a CSV using rules inferred from an example output file, including age calculation and nested object construction.

## Task

The model reads `input/input.csv` (Beatles biographical data) and `input/expected_format.json` (a single example record showing the desired output shape), infers the transformation rules, and outputs a JSON array for all four Beatles. Ages are calculated as of July 1, 2025.

The model must figure out from the example: field renaming, date formatting, age calculation, relative filtering (exclude null relationships), and nested structure.

This test runs in all 4 languages (Python, TypeScript, Go, C#) using the same input files.

## Input

Both files are copied into the working directory at runtime.

- `input/input.csv` — raw data: names, birthdays, death dates, relatives
- `input/expected_format.json` — one example record showing the target JSON shape

## Scoring

| Check | Description |
|---|---|
| `valid_json` | stdout parses as a JSON array |
| `four_records` | exactly 4 records |
| `correct_first_names` | all of John, James (Paul), Ringo, George present |
| `correct_last_names` | all of Lennon, McCartney, Starr, Harrison present |
| `birthday_format` | all `Birthday` fields match `YYYY-MM-DD` |
| `age_john` | John Lennon's age is 40 (died 1980, before cutoff) |
| `age_paul` | Paul McCartney's age is 83 |
| `age_ringo` | Ringo Starr's age is 84 |
| `age_george` | George Harrison's age is 58 (died 2001, before cutoff) |
| `null_relatives_excluded` | John's null-relationship relatives are not included |
| `relative_count_george` | George has exactly 4 relatives |

All checks must pass and exit code must be 0.

## What this tests

Rule inference from examples, date arithmetic, conditional data inclusion, and nested JSON construction. Most models fail entirely — wrong CSV column assumptions, incorrect age calculations, and missing output are the most common failure modes.

## Maintenance note

When a living Beatle dies, three files must be updated: `input/input.csv` (fill in `Died`), `input/expected_format.json` (update `Age`), and `verify.py` (update the hardcoded age check).
