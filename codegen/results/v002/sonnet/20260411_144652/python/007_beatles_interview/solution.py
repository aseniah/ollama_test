import csv
import json
import sys
from datetime import date, datetime

CALC_DATE = date(2025, 7, 1)

def parse_name(full_name):
    parts = full_name.strip().split()
    if len(parts) >= 2:
        return parts[0], parts[-1]
    return full_name, ""

def calc_age(birthday, died):
    end = died if died else CALC_DATE
    age = end.year - birthday.year
    if (end.month, end.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def parse_date(s):
    return datetime.strptime(s.strip(), "%m/%d/%Y").date()

rows = []
with open("input/input.csv", newline="") as f:
    reader = csv.DictReader(f)
    for row in reader:
        birthday = parse_date(row["Birthday"])
        died = parse_date(row["Died"]) if row["Died"].strip().lower() != "null" else None
        first, last = parse_name(row["Name"])
        age = calc_age(birthday, died)
        relatives = []
        for rel in ["Father", "Mother", "Brother", "Sister"]:
            val = row[rel].strip()
            if val.lower() != "null":
                rfirst, rlast = parse_name(val)
                relatives.append({"FirstName": rfirst, "LastName": rlast, "Relationship": rel})
        rows.append({
            "FirstName": first,
            "LastName": last,
            "Birthday": birthday.strftime("%Y-%m-%d"),
            "Age": age,
            "Relatives": relatives,
        })

print(json.dumps(rows, indent=2))
