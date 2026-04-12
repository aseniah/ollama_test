import csv
import json
from datetime import datetime

# Read input CSV file
with open('input/input.csv', mode='r') as csv_file:
    csv_reader = csv.DictReader(csv_file)
    
    # Dictionary to hold data
    data = []
    
    for row in csv_reader:
        # Parse the date from the birthday column
        try:
            birth_date = datetime.strptime(row['Birthday'], '%m/%d/%Y')
        except ValueError:
            raise ValueError(f"Invalid date format in row: {row['Birthday']}")
        
        # Calculate age as of July 1, 2025
        age = (datetime(2025, 7, 1) - birth_date).days // 365
        
        # Create a new entry with calculated age
        new_entry = {
            "FirstName": row['Name'],
            "LastName": row['LastName'],
            "Birthday": row['Birthday'],
            "Age": age,
            "Relatives": []
        }
        
        # Populate relatives from additional columns
        try:
            father = row['Father']
            mother = row['Mother']
            brother = row['Brother']
            sister = row['Sister']
        except KeyError:
            raise KeyError(f"Missing relative data in row: {row}")
        
        if father:
            new_entry['Relatives'].append({
                "FirstName": father,
                "LastName": row['LastName'],
                "Relationship": "Father"
            })
        
        if mother:
            new_entry['Relatives'].append({
                "FirstName": mother,
                "LastName": row['LastName'],
                "Relationship": "Mother"
            })
        
        if brother:
            new_entry['Relatives'].append({
                "FirstName": brother,
                "LastName": row['LastName'],
                "Relationship": "Brother"
            })
        
        if sister:
            new_entry['Relatives'].append({
                "FirstName": sister,
                "LastName": row['LastName'],
                "Relationship": "Sister"
            })
        
        data.append(new_entry)

# Write the JSON array to stdout
json.dump(data, stdout, indent=4)