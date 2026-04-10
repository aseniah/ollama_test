import csv
import json
from datetime import date

AGE_REFERENCE = date(2025, 7, 1)

def parse_date(s):
    if not s or s.strip().lower() == 'null':
        return None
    parts = s.strip().split('/')
    if len(parts) == 3:
        month, day, year = int(parts[0]), int(parts[1]), int(parts[2])
        return date(year, month, day)
    return None

def calc_age(birthday, as_of):
    age = as_of.year - birthday.year
    if (as_of.month, as_of.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def split_name(full_name):
    parts = full_name.strip().split()
    first = parts[0] if parts else ''
    last = parts[-1] if len(parts) > 1 else ''
    return first, last

def main():
    with open('input/input.csv', newline='') as f:
        reader = csv.DictReader(f)
        rows = list(reader)

    result = []
    rel_columns = ['Father', 'Mother', 'Brother', 'Sister']

    for row in rows:
        full_name = row['Name'].strip()
        first, last = split_name(full_name)

        birthday = parse_date(row['Birthday'])
        died = parse_date(row['Died'])

        age_date = died if died else AGE_REFERENCE
        age = calc_age(birthday, age_date)

        relatives = []
        for rel_type in rel_columns:
            rel_name = row.get(rel_type, '').strip()
            if rel_name and rel_name.lower() != 'null':
                rel_first, rel_last = split_name(rel_name)
                relatives.append({
                    'FirstName': rel_first,
                    'LastName': rel_last,
                    'Relationship': rel_type
                })

        entry = {
            'FirstName': first,
            'LastName': last,
            'Birthday': birthday.strftime('%Y-%m-%d'),
            'Age': age,
            'Relatives': relatives
        }
        result.append(entry)

    print(json.dumps(result, indent=2))

if __name__ == '__main__':
    main()
