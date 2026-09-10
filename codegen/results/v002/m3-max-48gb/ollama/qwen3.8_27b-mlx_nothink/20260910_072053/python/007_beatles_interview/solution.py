import csv
import json
from datetime import date

def calculate_age(birthday_str, reference_date):
    """Calculate age as of reference_date."""
    # Parse birthday from MM/DD/YYYY format
    parts = birthday_str.strip().split('/')
    month = int(parts[0])
    day = int(parts[1])
    year = int(parts[2])
    birth_date = date(year, month, day)
    
    # Calculate age
    age = reference_date.year - birth_date.year
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def process_name(full_name):
    """Split full name into first and last name."""
    parts = full_name.strip().split()
    if len(parts) == 1:
        return parts[0], ""
    else:
        first_name = parts[0]
        last_name = " ".join(parts[1:])
        return first_name, last_name

def parse_relative(relative_str):
    """Parse a relative string like 'Alfred Lennon' into a dict with FirstName, LastName, and null check."""
    if not relative_str or relative_str.strip().lower() == 'null' or relative_str.strip() == '':
        return None
    first_name, last_name = process_name(relative_str)
    return {
        "FirstName": first_name,
        "LastName": last_name
    }

# Reference date: July 1, 2025
reference_date = date(2025, 7, 1)

result = []

with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    for row in reader:
        first_name, last_name = process_name(row['Name'])
        birthday_str = row['Birthday'].strip()
        
        # Format birthday as YYYY-MM-DD
        parts = birthday_str.split('/')
        birthday_formatted = f"{parts[2]}-{parts[0].zfill(2)}-{parts[1].zfill(2)}"
        
        # Calculate age
        age = calculate_age(birthday_str, reference_date)
        
        # Build relatives list
        relatives = []
        
        father_rel = parse_relative(row['Father'])
        if father_rel:
            father_rel['Relationship'] = 'Father'
            relatives.append(father_rel)
        
        mother_rel = parse_relative(row['Mother'])
        if mother_rel:
            mother_rel['Relationship'] = 'Mother'
            relatives.append(mother_rel)
        
        brother_rel = parse_relative(row['Brother'])
        if brother_rel:
            brother_rel['Relationship'] = 'Brother'
            relatives.append(brother_rel)
        
        sister_rel = parse_relative(row['Sister'])
        if sister_rel:
            sister_rel['Relationship'] = 'Sister'
            relatives.append(sister_rel)
        
        person = {
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": birthday_formatted,
            "Age": age,
            "Relatives": relatives
        }
        result.append(person)

print(json.dumps(result, indent=3))