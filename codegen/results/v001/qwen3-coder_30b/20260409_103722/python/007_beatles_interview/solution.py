import csv
import json
from datetime import date

def calculate_age(birth_date_str):
    birth_date = date.fromisoformat(birth_date_str)
    reference_date = date(2025, 7, 1)
    age = reference_date.year - birth_date.year
    if reference_date.month < birth_date.month or (reference_date.month == birth_date.month and reference_date.day < birth_date.day):
        age -= 1
    return age

output = []
with open('input/input.csv', 'r') as file:
    reader = csv.DictReader(file)
    for row in reader:
        output.append({
            'id': int(row['id']),
            'name': row['name'],
            'age': calculate_age(row['birth_date']),
            'email': row['email']
        })

print(json.dumps(output, indent=None))