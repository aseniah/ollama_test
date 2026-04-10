import csv
import json
import sys
from datetime import date

AGE_REFERENCE = date(2025, 7, 1)


def calc_age(birthday: date, died: date | None) -> int:
    end = died if died is not None else AGE_REFERENCE
    age = end.year - birthday.year
    if (end.month, end.day) < (birthday.month, birthday.day):
        age -= 1
    return age


def parse_date(s: str) -> date | None:
    if not s or s.strip().lower() == "null":
        return None
    return date.fromisoformat(s.strip())


def parse_csv_date(s: str) -> date | None:
    if not s or s.strip().lower() == "null":
        return None
    # Format: M/D/YYYY
    parts = s.strip().split("/")
    return date(int(parts[2]), int(parts[0]), int(parts[1]))


def split_name(full: str) -> tuple[str, str]:
    parts = full.strip().split()
    return parts[0], parts[-1]


def main():
    input_path = "input/input.csv"
    results = []

    with open(input_path, newline="", encoding="utf-8") as f:
        reader = csv.DictReader(f)
        for row in reader:
            birthday = parse_csv_date(row["Birthday"])
            died = parse_csv_date(row["Died"])
            age = calc_age(birthday, died)

            first, last = split_name(row["Name"])

            relatives = []
            for rel in ("Father", "Mother", "Brother", "Sister"):
                val = row.get(rel, "").strip()
                if val and val.lower() != "null":
                    rf, rl = split_name(val)
                    relatives.append({
                        "FirstName": rf,
                        "LastName": rl,
                        "Relationship": rel,
                    })

            results.append({
                "FirstName": first,
                "LastName": last,
                "Birthday": birthday.strftime("%Y-%m-%d"),
                "Age": age,
                "Relatives": relatives,
            })

    print(json.dumps(results, indent=2))


if __name__ == "__main__":
    main()
