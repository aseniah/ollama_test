import csv
import json
from datetime import date

def calculate_age(birthday_str, as_of_date):
    """Calculate age as of the given date."""
    parts = birthday_str.split('/')
    month = int(parts[0])
    day = int(parts[1])
    year = int(parts[2])
    birth_date = date(year, month, day)
    
    age = as_of_date.year - birth_date.year
    # If birthday hasn't occurred yet this year, subtract 1
    if (as_of_date.month, as_of_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def main():
    as_of = date(2025, 7, 1)
    result = []
    
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        for row in reader:
            # Parse name
            name_parts = row['Name'].split(' ')
            first_name = name_parts[0]
            last_name = name_parts[-1]
            
            # Calculate age
            age = calculate_age(row['Birthday'], as_of)
            
            # Format birthday to YYYY-MM-DD
            bday_parts = row['Birthday'].split('/')
            birthday_formatted = f"{bday_parts[2]}-{bday_parts[0]:>02}-{bday_parts[1]:>02}"
            
            # Build relatives list
            relatives = []
            
            if row['Father'] != 'null':
                parts = row['Father'].split(' ')
                relatives.append({
                    "FirstName": parts[0],
                    "LastName": parts[-1],
                    "Relationship": "Father"
                })
            
            if row['Mother'] != 'null':
                parts = row['Mother'].split(' ')
                relatives.append({
                    "FirstName": parts[0],
                    "LastName": parts[-1],
                    "Relationship": "Mother"
                })
            
            if row['Brother'] != 'null':
                parts = row['Brother'].split(' ')
                relatives.append({
                    "FirstName": parts[0],
                    "LastName": parts[-1],
                    "Relationship": "Brother"
                })
            
            if row['Sister'] != 'null':
                parts = row['Sister'].split(' ')
                relatives.append({
                    "FirstName": parts[0],
                    "LastName": parts[-1],
                    "Relationship": "Sister"
                })
            
            entry = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday_formatted,
                "Age": age,
                "Relatives": relatives
            }
            result.append(entry)
    
    print(json.dumps(result, indent=2))

if __name__ == '__main__':
    main()