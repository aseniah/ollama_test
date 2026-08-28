import csv
import json
from datetime import date

def parse_date(date_str):
    parts = date_str.split('/')
    return date(int(parts[2]), int(parts[0]), int(parts[1]))

def calculate_age(birthday, reference_date):
    age = reference_date.year - birthday.year
    if (reference_date.month, reference_date.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def main():
    results = []
    
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        for row in reader:
            # Parse name
            name_parts = row['Name'].split()
            first_name = name_parts[0]
            last_name = name_parts[-1]
            
            # Parse birthday
            birthday = parse_date(row['Birthday'])
            birthday_str = f"{birthday.year:04d}-{birthday.month:02d}-{birthday.day:02d}"
            
            # Calculate age
            if row['Died'] and row['Died'].strip() != 'null':
                died = parse_date(row['Died'].strip())
                age = calculate_age(birthday, died)
            else:
                reference = date(2025, 7, 1)
                age = calculate_age(birthday, reference)
            
            # Build relatives
            relatives = []
            for rel_type in ['Father', 'Mother', 'Brother', 'Sister']:
                rel_val = row[rel_type].strip() if row[rel_type] else ''
                if rel_val and rel_val != 'null':
                    rel_parts = rel_val.split()
                    relatives.append({
                        'FirstName': rel_parts[0],
                        'LastName': rel_parts[-1],
                        'Relationship': rel_type
                    })
            
            results.append({
                'FirstName': first_name,
                'LastName': last_name,
                'Birthday': birthday_str,
                'Age': age,
                'Relatives': relatives
            })
    
    print(json.dumps(results, indent=2))

main()