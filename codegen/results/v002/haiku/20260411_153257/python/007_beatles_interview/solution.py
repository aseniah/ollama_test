import csv
import json
from datetime import date

def parse_date(s):
    m, d, y = s.strip().split('/')
    return date(int(y), int(m), int(d))

def calc_age(birthday, died, ref):
    end = died if died else ref
    age = end.year - birthday.year
    if (end.month, end.day) < (birthday.month, birthday.day):
        age -= 1
    return age

ref = date(2025, 7, 1)

rows = []
with open('input/input.csv', newline='') as f:
    reader = csv.DictReader(f)
    for row in reader:
        name_parts = row['Name'].split()
        first = name_parts[0]
        last = name_parts[-1]
        birthday = parse_date(row['Birthday'])
        died_str = row['Died'].strip()
        died = parse_date(died_str) if died_str and died_str.lower() != 'null' else None
        age = calc_age(birthday, died, ref)
        
        relatives = []
        for rel in ['Father', 'Mother', 'Brother', 'Sister']:
            val = row[rel].strip()
            if val and val.lower() != 'null':
                parts = val.split()
                rel_first = parts[0]
                rel_last = parts[-1]
                relatives.append({
                    "FirstName": rel_first,
                    "LastName": rel_last,
                    "Relationship": rel
                })
        
        rows.append({
            "FirstName": first,
            "LastName": last,
            "Birthday": birthday.strftime('%Y-%m-%d'),
            "Age": age,
            "Relatives": relatives
        })

print(json.dumps(rows, indent=2))
