import csv
import json
from datetime import date

def main():
    # Read the CSV file
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        rows = list(reader)
    
    # Reference date: July 1, 2025
    reference_date = date(2025, 7, 1)
    
    output = []
    
    for row in rows:
        name_parts = row['Name'].split()
        first_name = name_parts[0]
        last_name = name_parts[1]
        
        # Convert birthday from MM/DD/YYYY to YYYY-MM-DD
        bday_parts = row['Birthday'].split('/')
        birthday = f"{bday_parts[2]}-{bday_parts[0]}-{bday_parts[1]}"
        birth_date = date.fromisoformat(birthday)
        
        # Calculate age as of July 1, 2025
        # Age is number of full years before or since birthday, as of July 1
        age = 2025 - 1940  # Initial calculation based on birth year
        # Check if birthday in 2025 has already passed (Jan-Jun vs Jul)
        bday_month = int(bday_parts[0])
        if bday_month < 7:
            age = 2025 - birth_date.year - 1
        else:
            age = 2025 - birth_date.year
        
        # Handle relatives (skip if null)
        relatives = []
        father_name = row['Father']
        if father_name != 'null':
            father_parts = father_name.split()
            relatives.append({
                "FirstName": father_parts[0],
                "LastName": father_parts[1],
                "Relationship": "Father"
            })
        
        mother_name = row['Mother']
        if mother_name != 'null':
            mother_parts = mother_name.split()
            relatives.append({
                "FirstName": mother_parts[0],
                "LastName": mother_parts[1],
                "Relationship": "Mother"
            })
        
        brother_name = row['Brother']
        if brother_name != 'null':
            brother_parts = brother_name.split()
            relatives.append({
                "FirstName": brother_parts[0],
                "LastName": brother_parts[1],
                "Relationship": "Brother"
            })
        
        sister_name = row['Sister']
        if sister_name != 'null':
            sister_parts = sister_name.split()
            relatives.append({
                "FirstName": sister_parts[0],
                "LastName": sister_parts[1],
                "Relationship": "Sister"
            })
        
        person = {
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": birthday,
            "Age": age,
            "Relatives": relatives
        }
        output.append(person)
    
    # Output JSON to stdout
    print(json.dumps(output, indent=2))

if __name__ == '__main__':
    main()