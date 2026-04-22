import csv
import json
from datetime import date

def calculate_age(birth_date):
    ref_date = date(2025, 7, 1)
    age = ref_date.year - birth_date.year
    if (ref_date.month, ref_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def parse_birthday(bd_str):
    m, d, y = bd_str.strip().split('/')
    return date(int(y), int(m), int(d))

def parse_name(name_str):
    parts = name_str.strip().split()
    return parts[0], parts[-1]

result = []
with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    for row in reader:
        # Clean whitespace from keys and values to prevent parsing issues
        row = {k.strip(): v.strip() for k, v in row.items()}
        
        first_name, last_name = parse_name(row['Name'])
        bd = parse_birthday(row['Birthday'])
        age = calculate_age(bd)
        
        relatives = []
        for rel_type in ['Father', 'Mother', 'Brother', 'Sister']:
            val = row[rel_type]
            if val and val.lower() != 'null':
                r_first, r_last = parse_name(val)
                relatives.append({
                    "FirstName": r_first,
                    "LastName": r_last,
                    "Relationship": rel_type
                })
                
        result.append({
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": bd.strftime('%Y-%m-%d'),
            "Age": age,
            "Relatives": relatives
        })

print(json.dumps(result, indent=2))