import csv
import json
from datetime import date

def calculate_age(birthday, reference_date):
    """Calculate age as of the reference date."""
    year, month, day = birthday
    if month > reference_date.month or (month == reference_date.month and day > reference_date.day):
        return reference_date.year - year - 1
    return reference_date.year - year

def parse_date(date_str):
    """Parse a date string in M/D/YYYY format to (year, month, day) tuple."""
    parts = date_str.split('/')
    return (int(parts[2]), int(parts[0]), int(parts[1]))

def main():
    reference_date = date(2025, 7, 1)
    
    output = []
    
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        for row in reader:
            name = row['Name']
            parts = name.split()
            first_name = parts[0]
            last_name = parts[-1]
            
            birthday_str = row['Birthday']
            birthday_parsed = parse_date(birthday_str)
            birthday_iso = f"{birthday_parsed[0]:04d}-{birthday_parsed[1]:02d}-{birthday_parsed[2]:02d}"
            
            age = calculate_age(birthday_parsed, reference_date)
            
            relatives = []
            
            # Father
            father = row['Father']
            if father and father.lower() != 'null':
                father_parts = father.split()
                relatives.append({
                    "FirstName": father_parts[0],
                    "LastName": father_parts[-1],
                    "Relationship": "Father"
                })
            
            # Mother
            mother = row['Mother']
            if mother and mother.lower() != 'null':
                mother_parts = mother.split()
                relatives.append({
                    "FirstName": mother_parts[0],
                    "LastName": mother_parts[-1],
                    "Relationship": "Mother"
                })
            
            # Brother
            brother = row['Brother']
            if brother and brother.lower() != 'null':
                brother_parts = brother.split()
                relatives.append({
                    "FirstName": brother_parts[0],
                    "LastName": brother_parts[-1],
                    "Relationship": "Brother"
                })
            
            # Sister
            sister = row['Sister']
            if sister and sister.lower() != 'null':
                sister_parts = sister.split()
                relatives.append({
                    "FirstName": sister_parts[0],
                    "LastName": sister_parts[-1],
                    "Relationship": "Sister"
                })
            
            entry = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday_iso,
                "Age": age,
                "Relatives": relatives
            }
            output.append(entry)
    
    print(json.dumps(output, indent=2))

if __name__ == '__main__':
    main()