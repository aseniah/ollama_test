import csv
import json
from datetime import date

def parse_date(s):
    """Parse a date string like '10/9/1940' into a date object."""
    parts = s.strip().split('/')
    return date(int(parts[2]), int(parts[0]), int(parts[1]))

def format_date(d):
    """Format a date object as YYYY-MM-DD."""
    return d.strftime('%Y-%m-%d')

def calc_age(birthday, reference_date):
    """Calculate age at reference_date given a birthday."""
    age = reference_date.year - birthday.year
    if (reference_date.month, reference_date.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def parse_name(full_name):
    """Split 'First [Middle] Last' into (FirstName, LastName)."""
    parts = full_name.strip().split()
    first_name = parts[0]
    last_name = parts[-1]
    return first_name, last_name

def parse_relative(name, relationship):
    """Parse a relative's name and relationship into a dict, or None."""
    if name is None or name.strip().lower() == 'null' or name.strip() == '':
        return None
    parts = name.strip().split()
    first_name = parts[0]
    last_name = parts[-1]
    return {
        "FirstName": first_name,
        "LastName": last_name,
        "Relationship": relationship
    }

reference_date = date(2025, 7, 1)

results = []

with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    for row in reader:
        first_name, last_name = parse_name(row['Name'])
        birthday = parse_date(row['Birthday'])

        # Determine reference date for age: death date if exists, else July 1, 2025
        died_str = row['Died'].strip()
        if died_str.lower() == 'null' or died_str == '':
            ref = reference_date
        else:
            ref = parse_date(died_str)

        age = calc_age(birthday, ref)

        # Build relatives list
        relatives = []
        rel_fields = [
            ('Father', 'Father'),
            ('Mother', 'Mother'),
            ('Brother', 'Brother'),
            ('Sister', 'Sister'),
        ]
        for col, rel_name in rel_fields:
            r = parse_relative(row[col], rel_name)
            if r is not None:
                relatives.append(r)

        results.append({
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": format_date(birthday),
            "Age": age,
            "Relatives": relatives
        })

print(json.dumps(results, indent=2))