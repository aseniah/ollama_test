import csv
import json
import sys
from datetime import date

REFERENCE_DATE = date(2025, 7, 1)

def parse_birthday(s):
    m, d, y = s.strip().split('/')
    return date(int(y), int(m), int(d))

def calc_age(birthday, died_str):
    if died_str and died_str.strip().lower() != 'null':
        m, d, y = died_str.strip().split('/')
        end = date(int(y), int(m), int(d))
    else:
        end = REFERENCE_DATE
    age = end.year - birthday.year
    if (end.month, end.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def split_name(full):
    parts = full.strip().split()
    return parts[0], parts[-1]

def main():
    rows = []
    with open('input/input.csv', newline='') as f:
        reader = csv.DictReader(f)
        for row in reader:
            name_parts = row['Name'].strip().split()
            first = name_parts[0]
            last = name_parts[-1]
            birthday = parse_birthday(row['Birthday'])
            age = calc_age(birthday, row.get('Died', ''))
            relatives = []
            for rel in ('Father', 'Mother', 'Brother', 'Sister'):
                val = row.get(rel, '').strip()
                if val and val.lower() != 'null':
                    rf, rl = split_name(val)
                    relatives.append({
                        'FirstName': rf,
                        'LastName': rl,
                        'Relationship': rel
                    })
            rows.append({
                'FirstName': first,
                'LastName': last,
                'Birthday': birthday.strftime('%Y-%m-%d'),
                'Age': age,
                'Relatives': relatives
            })
    print(json.dumps(rows, indent=2))

main()
