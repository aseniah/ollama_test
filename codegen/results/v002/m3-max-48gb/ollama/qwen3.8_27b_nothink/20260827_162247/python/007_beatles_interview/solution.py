import csv
import json
from datetime import datetime

# Reference date for age calculation
reference_date = datetime(2025, 7, 1)

def calculate_age(birthday_str):
    """Calculate age as of July 1, 2025."""
    # Birthday is in MM/DD/YYYY format
    birth_date = datetime.strptime(birthday_str, "%m/%d/%Y")
    
    # Calculate age
    age = reference_date.year - birth_date.year
    
    # If the birthday hasn't occurred yet this year, subtract 1
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    
    return age

# Read input CSV
with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    rows = list(reader)

# Build output
output = []
for row in rows:
    # Parse name
    parts = row['Name'].split()
    first_name = parts[0]
    last_name = parts[-1]
    
    # Birthday - convert MM/DD/YYYY to YYYY-MM-DD
    birth_date = datetime.strptime(row['Birthday'], "%m/%d/%Y")
    birthday_formatted = birth_date.strftime("%Y-%m-%d")
    
    # Calculate age
    age = calculate_age(row['Birthday'])
    
    # Build relatives list
    relatives = []
    
    # Father
    if row['Father'] and row['Father'] != 'null':
        father_parts = row['Father'].split()
        relatives.append({
            "FirstName": father_parts[0],
            "LastName": father_parts[-1],
            "Relationship": "Father"
        })
    
    # Mother
    if row['Mother'] and row['Mother'] != 'null':
        mother_parts = row['Mother'].split()
        relatives.append({
            "FirstName": mother_parts[0],
            "LastName": mother_parts[-1],
            "Relationship": "Mother"
        })
    
    # Brother
    if row['Brother'] and row['Brother'] != 'null':
        brother_parts = row['Brother'].split()
        relatives.append({
            "FirstName": brother_parts[0],
            "LastName": brother_parts[-1],
            "Relationship": "Brother"
        })
    
    # Sister
    if row['Sister'] and row['Sister'] != 'null':
        sister_parts = row['Sister'].split()
        relatives.append({
            "FirstName": sister_parts[0],
            "LastName": sister_parts[-1],
            "Relationship": "Sister"
        })
    
    person = {
        "FirstName": first_name,
        "LastName": last_name,
        "Birthday": birthday_formatted,
        "Age": age,
        "Relatives": relatives
    }
    output.append(person)

# Output JSON array
print(json.dumps(output, indent=2))