import csv
import json
from datetime import date

def calc_age(birth, reference):
    age = reference.year - birth.year
    if (reference.month, reference.day) < (birth.month, birth.day):
        age -= 1
    return age

def split_name(full_name):
    parts = full_name.strip().split()
    return parts[0], parts[-1]

def parse_date(s):
    s = s.strip()
    if s.lower() == 'null' or s == '':
        return None
    parts = s.split('/')
    return date(int(parts[2]), int(parts[0]), int(parts[1]))

reference_date = date(2025, 7, 1)
results = []

with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    for row in reader:
        first_name, last_name = split_name(row['Name'])
        birthday = parse_date(row['Birthday'])

        death_date = parse_date(row['Died'])
        if death_date:
            age = calc_age(birthday, death_date)
        else:
            age = calc_age(birthday, reference_date)

        relatives = []
        for col, rel in [('Father', 'Father'), ('Mother', 'Mother'), ('Brother', 'Brother'), ('Sister', 'Sister')]:
            val = row[col].strip()
            if val.lower() != 'null' and val != '':
                r_first, r_last = split_name(val)
                relatives.append({
                    "FirstName": r_first,
                    "LastName": r_last,
                    "Relationship": rel
                })

        results.append({
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": birthday.strftime('%Y-%m-%d'),
            "Age": age,
            "Relatives": relatives
        })

print(json.dumps(results, indent=2))