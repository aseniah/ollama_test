import csv
import json
from datetime import date

# Read the CSV file
with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    rows = list(reader)

# Reference date for age calculation
ref_date = date(2025, 7, 1)

result = []

for row in rows:
    # Parse name
    full_name = row['Name'].strip()
    name_parts = full_name.split()
    first_name = name_parts[0]
    last_name = name_parts[-1]
    
    # Parse birthday (format: M/D/YYYY)
    birthday_str = row['Birthday'].strip()
    b_month, b_day, b_year = map(int, birthday_str.split('/'))
    birthday_obj = date(b_year, b_month, b_day)
    birthday_formatted = birthday_obj.isoformat()  # YYYY-MM-DD
    
    # Calculate age as of July 1, 2025
    if (ref_date.month, ref_date.day) >= (birthday_obj.month, birthday_obj.day):
        age = ref_date.year - birthday_obj.year
    else:
        age = ref_date.year - birthday_obj.year - 1
    
    # Build relatives list
    relatives = []
    
    father = row.get('Father', 'null').strip()
    if father and father.lower() != 'null':
        father_parts = father.split()
        relatives.append({
            "FirstName": father_parts[0],
            "LastName": father_parts[-1],
            "Relationship": "Father"
        })
    
    mother = row.get('Mother', 'null').strip()
    if mother and mother.lower() != 'null':
        mother_parts = mother.split()
        relatives.append({
            "FirstName": mother_parts[0],
            "LastName": mother_parts[-1],
            "Relationship": "Mother"
        })
    
    brother = row.get('Brother', 'null').strip()
    if brother and brother.lower() != 'null':
        brother_parts = brother.split()
        relatives.append({
            "FirstName": brother_parts[0],
            "LastName": brother_parts[-1],
            "Relationship": "Brother"
        })
    
    sister = row.get('Sister', 'null').strip()
    if sister and sister.lower() != 'null':
        sister_parts = sister.split()
        relatives.append({
            "FirstName": sister_parts[0],
            "LastName": sister_parts[-1],
            "Relationship": "Sister"
        })
    
    entry = {
        "FirstName": first_name,
        "LastName": last_name,
        "Birthday": birthday_formatted,
        "Age": age,
        "Relatives": relatives
    }
    result.append(entry)

# Output JSON array
print(json.dumps(result, indent=3))