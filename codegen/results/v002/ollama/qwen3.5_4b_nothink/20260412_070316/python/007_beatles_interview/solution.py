import csv
import json
from datetime import date

def parse_date(date_str):
    parts = date_str.split('/')
    day, month, year = int(parts[0]), int(parts[1]), int(parts[2])
    return date(year, month, day)

def calculate_age(birthday, reference_date):
    return (reference_date - birthday).days / 365.25
    # Rounding to nearest integer
    return int((reference_date - birthday).days / 365.25 + 0.5)

def main():
    reference_date = date(2025, 7, 1)
    
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        results = []
        
        for row in reader:
            # Parse birthday
            birthday = parse_date(row['Birthday'])
            age = calculate_age(birthday, reference_date)
            
            # Parse father
            father = None
            if row['Father'] != 'null':
                father = {
                    "FirstName": row['Father'].split()[0],
                    "LastName": row['Father'].split()[1],
                    "Relationship": "Father"
                }
            
            # Parse mother
            mother = None
            if row['Mother'] != 'null':
                mother = {
                    "FirstName": row['Mother'].split()[0],
                    "LastName": row['Mother'].split()[1],
                    "Relationship": "Mother"
                }
            
            relatives = []
            
            if father:
                relatives.append(father)
            if mother:
                relatives.append(mother)
            
            if row['Brother'] != 'null':
                relatives.append({
                    "FirstName": row['Brother'].split()[0],
                    "LastName": row['Brother'].split()[1],
                    "Relationship": "Brother"
                })
            
            if row['Sister'] != 'null':
                relatives.append({
                    "FirstName": row['Sister'].split()[0],
                    "LastName": row['Sister'].split()[1],
                    "Relationship": "Sister"
                })
            
            results.append({
                "FirstName": row['Name'].split()[0],
                "LastName": row['Name'].split()[1],
                "Birthday": f"{birthday.day:02d}-{birthday.month:02d}-{birthday.year:04d}",
                "Age": age,
                "Relatives": relatives
            })
    
    print(json.dumps(results, indent=2))

if __name__ == '__main__':
    main()