import csv
import json
from datetime import date

def parse_date(date_str):
    if date_str == 'null':
        return None
    parts = date_str.split('/')
    return date(int(parts[2]), int(parts[0]), int(parts[1]))

def calculate_age(birthday):
    if birthday is None:
        return None
    reference_date = date(2025, 7, 1)
    age = reference_date.year - birthday.year
    if (reference_date.month, reference_date.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def split_name(full_name):
    parts = full_name.split()
    return parts[0], ' '.join(parts[1:])

def parse_relative(name, relationship):
    if name == 'null' or name is None:
        return None
    first_name, last_name = split_name(name)
    return {
        "FirstName": first_name,
        "LastName": last_name,
        "Relationship": relationship
    }

result = []

with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    for row in reader:
        full_name = row['Name']
        first_name, last_name = split_name(full_name)
        birthday = parse_date(row['Birthday'])
        age = calculate_age(birthday)

        relatives = []
        if row['Father'] != 'null':
            relative = parse_relative(row['Father'], 'Father')
            if relative:
                relatives.append(relative)

        if row['Mother'] != 'null':
            relative = parse_relative(row['Mother'], 'Mother')
            if relative:
                relatives.append(relative)

        if row['Brother'] != 'null':
            relative = parse_relative(row['Brother'], 'Brother')
            if relative:
                relatives.append(relative)

        if row['Sister'] != 'null':
            relative = parse_relative(row['Sister'], 'Sister')
            if relative:
                relatives.append(relative)

        person = {
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": birthday.strftime('%Y-%m-%d') if birthday else None,
            "Age": age,
            "Relatives": relatives
        }
        result.append(person)

print(json.dumps(result, indent=2))
