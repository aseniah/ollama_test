#!/usr/bin/env python3
import csv
import json
from datetime import date
from dateutil.relativedelta import relativedelta

# Read expected output format to understand structure
with open('input/expected_format.json', 'r') as f:
    expected = json.load(f)

# Read CSV input
with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    rows = list(reader)

# Calculate ages as of July 1, 2025
reference_date = date(2025, 7, 1)
ages = []

for row in rows:
    # Find birth date field (common formats: birthdate, dob, date_of_birth)
    birth_date = None
    for key in ['birth_date', 'date_of_birth', 'dob', 'birthdate']:
        if key in row:
            birth_date = row[key].strip()
            break
    
    if birth_date:
        try:
            birth_date = date.strptime(birth_date, '%Y-%m-%d')
        except ValueError:
            try:
                birth_date = date.strptime(birth_date, '%m/%d/%Y')
            except ValueError:
                birth_date = date.strptime(birth_date, '%d/%m/%Y')
        
        # Calculate age as of July 1, 2025
        age = relativedelta(reference_date, birth_date).years
        if age < 0:
            age = 0
        
        # Use expected format structure
        entry = {}
        for key in expected:
            if key in row:
                entry[key] = row[key]
        entry['age'] = age
        ages.append(entry)

# Output JSON array to stdout
print(json.dumps(ages, indent=2))