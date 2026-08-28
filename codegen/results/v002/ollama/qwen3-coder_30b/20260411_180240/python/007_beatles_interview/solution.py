import csv
from datetime import datetime
import json

def calculate_age(birthday_str, reference_date):
    birthday = datetime.strptime(birthday_str, "%m/%d/%Y")
    age = reference_date.year - birthday.year
    if reference_date.month < birthday.month or (reference_date.month == birthday.month and reference_date.day < birthday.day):
        age -= 1
    return age

# Reference date: July 1, 2025
reference_date = datetime(2025, 7, 1)

# Read CSV file
people = []
with open('input/input.csv', 'r') as file:
    reader = csv.DictReader(file)
    for row in reader:
        # Parse name
        name_parts = row['Name'].split()
        first_name = name_parts[0]
        last_name = name_parts[-1]
        
        # Parse birthday
        birthday = row['Birthday']
        
        # Calculate age
        age = calculate_age(birthday, reference_date)
        
        # Build relatives list
        relatives = []
        
        # Add father
        if row['Father'] != 'null':
            father_parts = row['Father'].split()
            father_first = father_parts[0]
            father_last = father_parts[-1]
            relatives.append({
                "FirstName": father_first,
                "LastName": father_last,
                "Relationship": "Father"
            })
        
        # Add mother
        if row['Mother'] != 'null':
            mother_parts = row['Mother'].split()
            mother_first = mother_parts[0]
            mother_last = mother_parts[-1]
            relatives.append({
                "FirstName": mother_first,
                "LastName": mother_last,
                "Relationship": "Mother"
            })
        
        # Add brother
        if row['Brother'] != 'null':
            brother_parts = row['Brother'].split()
            brother_first = brother_parts[0]
            brother_last = brother_parts[-1]
            relatives.append({
                "FirstName": brother_first,
                "LastName": brother_last,
                "Relationship": "Brother"
            })
        
        # Add sister
        if row['Sister'] != 'null':
            sister_parts = row['Sister'].split()
            sister_first = sister_parts[0]
            sister_last = sister_parts[-1]
            relatives.append({
                "FirstName": sister_first,
                "LastName": sister_last,
                "Relationship": "Sister"
            })
        
        # Create person object
        person = {
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": birthday,
            "Age": age,
            "Relatives": relatives
        }
        
        people.append(person)

# Output JSON array
print(json.dumps(people, indent=2))