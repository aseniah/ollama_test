import csv
import json
from datetime import date

# Reference date for age calculation
ref_date = date(2025, 7, 1)

def parse_date(s):
    """Parse MM/DD/YYYY to a date object."""
    parts = s.strip().split('/')
    return date(int(parts[2]), int(parts[0]), int(parts[1]))

def format_date(d):
    """Format a date object as YYYY-MM-DD."""
    return d.strftime('%Y-%m-%d')

def calc_age(birth, death):
    """Calculate age: at death if dead, otherwise as of ref_date."""
    if death:
        target = death
    else:
        target = ref_date
    age = target.year - birth.year
    if (target.month, target.day) < (birth.month, birth.day):
        age -= 1
    return age

def parse_name(full_name):
    """Split 'First Middle... Last' into first and last name."""
    parts = full_name.strip().split()
    first = parts[0]
    last = parts[-1]
    return first, last

def parse_relatives(row):
    """Build relatives list from CSV row, skipping null entries."""
    relatives = []
    for rel_type in ['Father', 'Mother', 'Brother', 'Sister']:
        val = row[rel_type].strip()
        if val.lower() != 'null':
            first, last = parse_name(val)
            relatives.append({
                "FirstName": first,
                "LastName": last,
                "Relationship": rel_type
            })
    return relatives

# Read CSV
with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    records = []
    for row in reader:
        birth = parse_date(row['Birthday'])
        death_str = row['Died'].strip()
        death = parse_date(death_str) if death_str.lower() != 'null' else None

        first, last = parse_name(row['Name'])

        entry = {
            "FirstName": first,
            "LastName": last,
            "Birthday": format_date(birth),
            "Age": calc_age(birth, death),
            "Relatives": parse_relatives(row)
        }
        records.append(entry)

# Output JSON
print(json.dumps(records, indent=2))