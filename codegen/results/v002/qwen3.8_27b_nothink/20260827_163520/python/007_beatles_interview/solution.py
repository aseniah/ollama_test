import csv
import json
from datetime import date

def calculate_age(birthday_str, as_of=date(2025, 7, 1)):
    # Parse birthday (MM/DD/YYYY)
    parts = birthday_str.split('/')
    month = int(parts[0])
    day = int(parts[1])
    year = int(parts[2])
    birthday = date(year, month, day)
    
    age = as_of.year - birthday.year
    # Check if birthday hasn't occurred yet this year
    if (as_of.month, as_of.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def main():
    results = []
    
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        for row in reader:
            # Extract first and last name
            full_name = row['Name'].strip()
            name_parts = full_name.split()
            first_name = name_parts[0]
            last_name = name_parts[-1]
            
            # Parse birthday
            birthday_str = row['Birthday'].strip()
            
            # Calculate age
            age = calculate_age(birthday_str)
            
            # Build relatives list
            relatives = []
            
            # Father
            father = row.get('Father', 'null').strip()
            if father and father.lower() != 'null':
                father_parts = father.split()
                relatives.append({
                    "FirstName": father_parts[0],
                    "LastName": father_parts[-1],
                    "Relationship": "Father"
                })
            
            # Mother
            mother = row.get('Mother', 'null').strip()
            if mother and mother.lower() != 'null':
                mother_parts = mother.split()
                relatives.append({
                    "FirstName": mother_parts[0],
                    "LastName": mother_parts[-1],
                    "Relationship": "Mother"
                })
            
            # Brother
            brother = row.get('Brother', 'null').strip()
            if brother and brother.lower() != 'null':
                brother_parts = brother.split()
                relatives.append({
                    "FirstName": brother_parts[0],
                    "LastName": brother_parts[-1],
                    "Relationship": "Brother"
                })
            
            # Sister
            sister = row.get('Sister', 'null').strip()
            if sister and sister.lower() != 'null':
                sister_parts = sister.split()
                relatives.append({
                    "FirstName": sister_parts[0],
                    "LastName": sister_parts[-1],
                    "Relationship": "Sister"
                })
            
            # Format birthday to YYYY-MM-DD
            b_parts = birthday_str.split('/')
            formatted_birthday = f"{int(b_parts[2])}-{int(b_parts[0]):02d}-{int(b_parts[1]):02d}"
            
            entry = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": formatted_birthday,
                "Age": age,
                "Relatives": relatives
            }
            results.append(entry)
    
    print(json.dumps(results, indent=2))

if __name__ == '__main__':
    main()