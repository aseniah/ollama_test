import csv
import json
from datetime import date

def split_name(full_name):
    parts = full_name.strip().split()
    first = parts[0]
    last = parts[-1]
    return first, last

def calculate_age(birthday, died, reference_date):
    birth_year, birth_month, birth_day = map(int, birthday.split('/'))
    birth = date(birth_year, birth_month, birth_day)
    
    if died != 'null':
        d_month, d_day, d_year = map(int, died.split('/'))
        death = date(d_year, d_month, d_day)
    else:
        death = reference_date  # July 1, 2025
    
    age = death.year - birth.year
    if (death.month, death.day) < (birth.month, birth.day):
        age -= 1
    return age

reference = date(2025, 7, 1)
results = []

with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    for row in reader:
        first_name, last_name = split_name(row['Name'])
        
        birthday_parts = row['Birthday'].split('/')
        b_month = int(birthday_parts[0])
        b_day = int(birthday_parts[1])
        b_year = int(birthday_parts[2])
        formatted_birthday = f"{b_year}-{b_month:02d}-{b_day:02d}"
        
        age = calculate_age(row['Birthday'], row['Died'], reference)
        
        relatives = []
        for key in ['Father', 'Mother', 'Brother', 'Sister']:
            value = row[key].strip()
            if value != 'null':
                rel_first, rel_last = split_name(value)
                relatives.append({
                    "FirstName": rel_first,
                    "LastName": rel_last,
                    "Relationship": key
                })
        
        results.append({
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": formatted_birthday,
            "Age": age,
            "Relatives": relatives
        })

print(json.dumps(results, indent=2))