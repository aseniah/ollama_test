import csv
import json
from datetime import datetime

def parse_date(date_str):
    """Convert DD/MM/YYYY to YYYY-MM-DD"""
    if date_str is None or date_str.lower() == 'null':
        return None
    parts = date_str.split('/')
    return f"{parts[2]}-{parts[1]}-{parts[0]}"

def calculate_age(birthday):
    """Calculate age as of July 1, 2025"""
    if birthday is None:
        return None
    birth_date = datetime.strptime(birthday, "%Y-%m-%d")
    reference_date = datetime(2025, 7, 1)
    age = reference_date.year - birth_date.year
    # If birthday hasn't occurred yet this year, subtract 1
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def parse_name(name_str):
    """Extract first and last name from full name"""
    parts = name_str.strip().split()
    return parts[0], parts[-1]

def parse_relative(rel_name):
    """Parse relative name and create dict structure"""
    if rel_name is None or rel_name.lower() == 'null':
        return None
    first_name, last_name = parse_name(rel_name)
    return {
        "FirstName": first_name,
        "LastName": last_name
    }

# Read CSV file
with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    output = []
    
    for row in reader:
        first_name, last_name = parse_name(row['Name'])
        birthday = parse_date(row['Birthday'])
        age = calculate_age(birthday)
        
        # Parse relatives
        relatives = []
        
        father_rel = parse_relative(row['Father'])
        if father_rel:
            relatives.append(father_rel)
        
        mother_rel = parse_relative(row['Mother'])
        if mother_rel:
            relatives.append(mother_rel)
        
        brother_rel = parse_relative(row['Brother'])
        if brother_rel:
            relatives.append(brother_rel)
        
        sister_rel = parse_relative(row['Sister'])
        if sister_rel:
            relatives.append(sister_rel)
        
        output.append({
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": birthday,
            "Age": age,
            "Relatives": relatives
        })
    
    # Output JSON array to stdout
    print(json.dumps(output, indent=2))