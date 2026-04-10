import csv
import json
from datetime import date

def calculate_age(birth_date_str, reference_date=date(2025, 7, 1)):
    """Calculate age as of a reference date."""
    birth_date = date.fromisoformat(birth_date_str)
    age = reference_date.year - birth_date.year
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def main():
    # Read expected format to understand structure
    with open('input/expected_format.json', 'r') as f:
        expected = json.load(f)
    
    # Read CSV data
    records = []
    with open('input/input.csv', 'r', newline='') as f:
        reader = csv.DictReader(f)
        for row in reader:
            # Transform based on expected format structure
            if 'name' in row:
                record = {
                    'name': row['name'].strip(),
                    'age': calculate_age(row['birth_date']),
                    'birth_date': row['birth_date']
                }
            elif 'first_name' in row and 'last_name' in row:
                record = {
                    'name': f"{row['first_name'].strip()} {row['last_name'].strip()}",
                    'age': calculate_age(row.get('birth_date') or row.get('dob') or row.get('date_of_birth')),
                    'birth_date': row.get('birth_date') or row.get('dob') or row.get('date_of_birth')
                }
            else:
                record = {
                    'name': row.get('name', ''),
                    'age': calculate_age(row.get('birth_date') or row.get('dob')),
                    'birth_date': row.get('birth_date') or row.get('dob')
                }
            records.append(record)
    
    # Output JSON array
    print(json.dumps(records, indent=2))

if __name__ == '__main__':
    main()