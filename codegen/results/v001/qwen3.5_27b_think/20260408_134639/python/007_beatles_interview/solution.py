import csv
import json
from datetime import date

def calculate_age(birth_date_str, reference_date=date(2025, 7, 1)):
    """Calculate age as of July 1, 2025"""
    birth_year, birth_month, birth_day = map(int, birth_date_str.split('-'))
    age = reference_date.year - birth_year
    if (reference_date.month, reference_date.day) < (birth_month, birth_day):
        age -= 1
    return age

def main():
    reference_date = date(2025, 7, 1)
    output = []
    
    with open('input/input.csv', 'r', encoding='utf-8') as f:
        reader = csv.DictReader(f)
        for row in reader:
            person_data = {
                'name': row.get('name', ''),
                'age': calculate_age(row.get('date_of_birth', '1990-01-01'), reference_date),
                'email': row.get('email', ''),
                'department': row.get('department', '')
            }
            output.append(person_data)
    
    print(json.dumps(output))

if __name__ == '__main__':
    main()