import csv
import json
from datetime import date, datetime

def calculate_age(dob_str):
    dob_str = dob_str.strip()
    formats = [
        '%Y-%m-%d',
        '%m/%d/%Y',
        '%d/%m/%Y',
        '%Y/%m/%d'
    ]
    for fmt in formats:
        try:
            dob = datetime.strptime(dob_str, fmt).date()
            break
        except ValueError:
            continue
    else:
        return None
    
    today = date(2025, 7, 1)
    age = today.year - dob.year
    if (today.month, today.day) < (dob.month, dob.day):
        age -= 1
    return age

def main():
    with open('input/input.csv', 'r') as file:
        reader = csv.DictReader(file)
        data = []
        for row in reader:
            entry = {}
            for key, value in row.items():
                lower_key = key.lower().strip()
                if lower_key == 'dateofbirth':
                    age = calculate_age(value)
                    entry['age'] = age
                else:
                    entry[lower_key] = value.strip()
            data.append(entry)
    print(json.dumps(data))

if __name__ == '__main__':
    main()