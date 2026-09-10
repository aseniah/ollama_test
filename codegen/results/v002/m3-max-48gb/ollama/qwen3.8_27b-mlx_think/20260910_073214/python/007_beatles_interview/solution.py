import json
import csv
from datetime import date

# Read the CSV file
with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    rows = list(reader)

def parse_date(date_str):
    """Parse M/D/YYYY or YYYY-MM-DD date string to a date object."""
    parts = date_str.split('/')
    return date(int(parts[2]), int(parts[0]), int(parts[1]))

def calc_age(birth_str, death_str, reference_date):
    """Calculate age. If died, use death date; otherwise use reference date."""
    birth = parse_date(birth_str)
    if death_str and death_str.lower() != 'null':
        ref = parse_date(death_str)
    else:
        ref = reference_date
    
    age = ref.year - birth.year
    if (ref.month, ref.day) < (birth.month, birth.day):
        age -= 1
    return age

reference_date = date(2025, 7, 1)

result = []

for row in rows:
    # Parse name
    name_parts = row['Name'].strip().split()
    first_name = name_parts[0]
    last_name = name_parts[-1]

    # Format birthday as YYYY-MM-DD
    bdate = parse_date(row['Birthday'])
    birthday_str = bdate.strftime('%Y-%m-%d')

    # Calculate age
    age = calc_age(row['Birthday'], row['Died'], reference_date)

    # Build relatives list
    relatives = []
    rel_fields = [
        ('Father', row['Father']),
        ('Mother', row['Mother']),
        ('Brother', row['Brother']),
        ('Sister', row['Sister'])
    ]

    for relationship, rel_name in rel_fields:
        if rel_name and rel_name.lower() != 'null':
            parts = rel_name.strip().split()
            relatives.append({
                "FirstName": parts[0],
                "LastName": parts[-1],
                "Relationship": relationship
            })

    entry = {
        "FirstName": first_name,
        "LastName": last_name,
        "Birthday": birthday_str,
        "Age": age,
        "Relatives": relatives
    }
    result.append(entry)

print(json.dumps(result, indent=3))