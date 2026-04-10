import csv
import json
from datetime import datetime

def calculate_age(birth_date_str):
    birth_date = datetime.strptime(birth_date_str, '%Y-%m-%d')
    target_date = datetime(2025, 7, 1)
    age = target_date.year - birth_date.year
    if (target_date.month, target_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def main():
    with open('input/input.csv', 'r') as csvfile:
        reader = csv.DictReader(csvfile)
        data = []
        for row in reader:
            transformed_row = {key.lower().replace(' ', '_'): value.strip() for key, value in row.items()}
            
            birth_date_key = None
            for key in transformed_row.keys():
                if 'birth' in key or 'dob' in key:
                    birth_date_key = key
                    break
            
            if birth_date_key and transformed_row.get(birth_date_key):
                age = calculate_age(transformed_row[birth_date_key])
                if birth_date_key != 'birth_date':
                    transformed_row['birth_date'] = transformed_row.pop(birth_date_key)
                transformed_row['age'] = age
            data.append(transformed_row)
        
        print(json.dumps(data))

if __name__ == '__main__':
    main()