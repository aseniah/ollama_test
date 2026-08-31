import csv
import json
from datetime import datetime

# Calculate age as of July 1, 2025
reference_date = datetime(2025, 7, 1)

# Read the CSV file
with open('input/input.csv', 'r') as file:
    reader = csv.DictReader(file)
    
    output = []
    
    for row in reader:
        # Parse birthday
        birthday = datetime.strptime(row['Birthday'], '%m/%d/%Y')
        
        # Calculate age
        age = reference_date.year - birthday.year
        if reference_date.month < birthday.month or (reference_date.month == birthday.month and reference_date.day < birthday.day):
            age -= 1
            
        # Prepare relatives list
        relatives = []
        
        # Add father
        if row['Father'] != 'null':
            father_name = row['Father'].split()
            relatives.append({
                "FirstName": father_name[0],
                "LastName": father_name[1] if len(father_name) > 1 else "",
                "Relationship": "Father"
            })
            
        # Add mother
        if row['Mother'] != 'null':
            mother_name = row['Mother'].split()
            relatives.append({
                "FirstName": mother_name[0],
                "LastName": mother_name[1] if len(mother_name) > 1 else "",
                "Relationship": "Mother"
            })
            
        # Add brother
        if row['Brother'] != 'null':
            brother_name = row['Brother'].split()
            relatives.append({
                "FirstName": brother_name[0],
                "LastName": brother_name[1] if len(brother_name) > 1 else "",
                "Relationship": "Brother"
            })
            
        # Add sister
        if row['Sister'] != 'null':
            sister_name = row['Sister'].split()
            relatives.append({
                "FirstName": sister_name[0],
                "LastName": sister_name[1] if len(sister_name) > 1 else "",
                "Relationship": "Sister"
            })
            
        # Create output object
        person = {
            "FirstName": row['Name'].split()[0],
            "LastName": row['Name'].split()[1] if len(row['Name'].split()) > 1 else "",
            "Birthday": birthday.strftime('%Y-%m-%d'),
            "Age": age,
            "Relatives": relatives
        }
        
        output.append(person)

# Write to stdout as JSON array
print(json.dumps(output, indent=2))