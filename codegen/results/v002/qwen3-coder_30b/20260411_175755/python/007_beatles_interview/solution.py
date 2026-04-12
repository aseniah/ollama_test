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
        # Extract person data
        first_name, last_name = row['Name'].split(' ', 1)
        birthday = row['Birthday']
        
        # Calculate age
        age = calculate_age(birthday, reference_date)
        
        # Build relatives list
        relatives = []
        
        # Add father
        if row['Father'] != 'null':
            father_name = row['Father'].split(' ', 1)
            relatives.append({
                "FirstName": father_name[0],
                "LastName": father_name[1] if len(father_name) > 1 else "",
                "Relationship": "Father"
            })
        
        # Add mother
        if row['Mother'] != 'null':
            mother_name = row['Mother'].split(' ', 1)
            relatives.append({
                "FirstName": mother_name[0],
                "LastName": mother_name[1] if len(mother_name) > 1 else "",
                "Relationship": "Mother"
            })
        
        # Add brother
        if row['Brother'] != 'null':
            brother_name = row['Brother'].split(' ', 1)
            relatives.append({
                "FirstName": brother_name[0],
                "LastName": brother_name[1] if len(brother_name) > 1 else "",
                "Relationship": "Brother"
            })
        
        # Add sister
        if row['Sister'] != 'null':
            sister_name = row['Sister'].split(' ', 1)
            relatives.append({
                "FirstName": sister_name[0],
                "LastName": sister_name[1] if len(sister_name) > 1 else "",
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