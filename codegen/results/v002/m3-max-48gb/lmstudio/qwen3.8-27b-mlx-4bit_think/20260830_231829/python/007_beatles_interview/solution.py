import csv
import json
from datetime import date

def parse_date(s):
    parts = s.strip().split('/')
    return date(int(parts[2]), int(parts[0]), int(parts[1]))

def calc_age(birth, reference):
    age = reference.year - birth.year
    if (reference.month, reference.day) < (birth.month, birth.day):
        age -= 1
    return age

reference_date = date(2025, 7, 1)

results = []

with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    for row in reader:
        # Parse name
        name_parts = row['Name'].strip().split()
        first_name = name_parts[0]
        last_name = name_parts[-1]

        # Parse birthday
        birthday = parse_date(row['Birthday'])
        birthday_iso = birthday.isoformat()

        # Calculate age
        if row['Died'].strip() != 'null':
            died = parse_date(row['Died'])
            age = calc_age(birthday, died)
        else:
            age = calc_age(birthday, reference_date)

        # Build relatives
        relatives = []
        for col, rel in [('Father', 'Father'), ('Mother', 'Mother'), ('Brother', 'Brother'), ('Sister', 'Sister')]:
            val = row[col].strip()
            if val != 'null':
                rel_parts = val.split()
                relatives.append({
                    "FirstName": rel_parts[0],
                    "LastName": rel_parts[-1],
                    "Relationship": rel
                })

        results.append({
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": birthday_iso,
            "Age": age,
            "Relatives": relatives
        })

print(json.dumps(results, indent=2))