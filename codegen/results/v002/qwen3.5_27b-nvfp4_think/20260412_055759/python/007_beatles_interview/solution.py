import csv
import json
from datetime import date

# Reference date for age calculation: July 1, 2025
reference_date = date(2025, 7, 1)

with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    result = []
    
    for row in reader:
        # Parse name into first and last
        parts = row['Name'].split()
        first_name = parts[0]
        last_name = parts[-1]
        
        # Parse birthday from M/D/YYYY format
        bday_parts = row['Birthday'].split('/')
        birth_month, birth_day, birth_year = int(bday_parts[0]), int(bday_parts[1]), int(bday_parts[2])
        
        # Calculate age as of July 1, 2025
        age = reference_date.year - birth_year
        if (reference_date.month, reference_date.day) < (birth_month, birth_day):
            age -= 1
        
        # Format birthday in ISO format (YYYY-MM-DD)
        birthday_iso = f"{birth_year}-{birth_month:02d}-{birth_day:02d}"
        
        # Build relatives list from Father, Mother, Brother, Sister columns
        relatives = []
        
        for rel_field, rel_type in [('Father', 'Father'), ('Mother', 'Mother'), 
                                     ('Brother', 'Brother'), ('Sister', 'Sister')]:
            value = row.get(rel_field, '').strip()
            if value and value.lower() != 'null':
                name_parts = value.split()
                relatives.append({
                    'FirstName': name_parts[0],
                    'LastName': name_parts[-1],
                    'Relationship': rel_type
                })
        
        result.append({
            'FirstName': first_name,
            'LastName': last_name,
            'Birthday': birthday_iso,
            'Age': age,
            'Relatives': relatives
        })

# Output only the JSON array to stdout
print(json.dumps(result))