import csv
import json
from datetime import date

def parse_date(s):
    """Parse M/D/YYYY to (year, month, day)"""
    parts = s.strip().split('/')
    return int(parts[0]), int(parts[1]), int(parts[2])

def format_iso(year, month, day):
    return f"{year:04d}-{month:02d}-{day:02d}"

def calc_age(birth_year, birth_month, birth_day, ref_date):
    """Calculate age as of ref_date (a date object)."""
    if (ref_date.month, ref_date.day) >= (birth_month, birth_day):
        return ref_date.year - birth_year
    else:
        return ref_date.year - birth_year - 1

ref_date = date(2025, 7, 1)

with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    records = []
    for row in reader:
        # Parse name
        name_parts = row['Name'].strip().split()
        first_name = name_parts[0]
        last_name = name_parts[-1]

        # Parse birthday
        bm, bd, by = parse_date(row['Birthday'])
        birthday_iso = format_iso(by, bm, bd)

        # Calculate age
        if row['Died'].strip().lower() != 'null' and row['Died'].strip():
            dm, dd, dy = parse_date(row['Died'])
            death_date = date(dy, dm, dd)
            age = calc_age(by, bm, bd, death_date)
        else:
            age = calc_age(by, bm, bd, ref_date)

        # Parse relatives
        relatives = []
        for rel_col, rel_name in [('Father', 'Father'), ('Mother', 'Mother'), ('Brother', 'Brother'), ('Sister', 'Sister')]:
            val = row[rel_col].strip()
            if val.lower() != 'null' and val:
                rel_parts = val.split(' ', 1)
                relatives.append({
                    "FirstName": rel_parts[0],
                    "LastName": rel_parts[1] if len(rel_parts) > 1 else "",
                    "Relationship": rel_name
                })

        records.append({
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": birthday_iso,
            "Age": age,
            "Relatives": relatives
        })

print(json.dumps(records, indent=3))