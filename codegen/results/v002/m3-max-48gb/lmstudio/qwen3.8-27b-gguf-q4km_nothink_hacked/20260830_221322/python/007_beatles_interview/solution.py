import csv
import json
from datetime import date

REFERENCE_DATE = date(2025, 7, 1)

def calculate_age(birthday_str):
    """Calculate age as of July 1, 2025."""
    month, day, year = birthday_str.split('/')
    birthday = date(int(year), int(month), int(day))
    age = REFERENCE_DATE.year - birthday.year
    if (REFERENCE_DATE.month, REFERENCE_DATE.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def parse_name(full_name):
    """Split full name into first and last name (first word is first name, rest is last name)."""
    parts = full_name.strip().split()
    first_name = parts[0]
    last_name = ' '.join(parts[1:]) if len(parts) > 1 else ''
    return first_name, last_name

def parse_relative(name_str):
    """Parse a relative's name string into first and last name."""
    if name_str is None or name_str.strip() == 'null':
        return None
    parts = name_str.strip().split()
    if len(parts) == 0:
        return None
    first_name = parts[0]
    last_name = ' '.join(parts[1:]) if len(parts) > 1 else ''
    return {
        "FirstName": first_name,
        "LastName": last_name
    }

results = []

with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    for row in reader:
        first_name, last_name = parse_name(row['Name'])
        age = calculate_age(row['Birthday'])
        
        relatives = []
        
        father = row.get('Father', '').strip()
        if father and father != 'null':
            rel = parse_relative(father)
            if rel:
                rel["Relationship"] = "Father"
                relatives.append(rel)
        
        mother = row.get('Mother', '').strip()
        if mother and mother != 'null':
            rel = parse_relative(mother)
            if rel:
                rel["Relationship"] = "Mother"
                relatives.append(rel)
        
        brother = row.get('Brother', '').strip()
        if brother and brother != 'null':
            rel = parse_relative(brother)
            if rel:
                rel["Relationship"] = "Brother"
                relatives.append(rel)
        
        sister = row.get('Sister', '').strip()
        if sister and sister != 'null':
            rel = parse_relative(sister)
            if rel:
                rel["Relationship"] = "Sister"
                relatives.append(rel)
        
        # Convert birthday format from MM/DD/YYYY to YYYY-MM-DD
        m, d, y = row['Birthday'].split('/')
        birthday_iso = f"{int(y)}-{int(m):02d}-{int(d):02d}"
        
        person = {
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": birthday_iso,
            "Age": age,
            "Relatives": relatives
        }
        results.append(person)

print(json.dumps(results, indent=2))