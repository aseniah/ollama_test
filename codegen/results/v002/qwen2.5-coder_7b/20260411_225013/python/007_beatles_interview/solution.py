import csv
from datetime import datetime

# Define the target date for age calculation
target_date = datetime(2025, 7, 1)

# Read the input CSV file and process its contents
with open('input/input.csv', mode='r') as csvfile:
    reader = csv.DictReader(csvfile)
    output_data = []
    
    # Process each row in the CSV file
    for row in reader:
        first_name, last_name = row['Name'].split()
        birthday = datetime.strptime(row['Birthday'], '%m/%d/%Y')
        
        # Calculate age as of July 1, 2025
        age = target_date.year - birthday.year - ((target_date.month, target_date.day) < (birthday.month, birthday.day))
        
        # Collect relative information
        relatives = []
        
        if row['Father']:
            relatives.append({
                "FirstName": row['Father'].split()[0],
                "LastName": row['Father'].split()[1] if len(row['Father'].split()) > 1 else "",
                "Relationship": "Father"
            })
        
        if row['Mother']:
            relatives.append({
                "FirstName": row['Mother'].split()[0],
                "LastName": row['Mother'].split()[1] if len(row['Mother'].split()) > 1 else "",
                "Relationship": "Mother"
            })
        
        if row['Brother']:
            relatives.append({
                "FirstName": row['Brother'].split()[0],
                "LastName": row['Brother'].split()[1] if len(row['Brother'].split()) > 1 else "",
                "Relationship": "Brother"
            })
        
        if row['Sister']:
            relatives.append({
                "FirstName": row['Sister'].split()[0],
                "LastName": row['Sister'].split()[1] if len(row['Sister'].split()) > 1 else "",
                "Relationship": "Sister"
            })
        
        # Append the processed data to the output list
        output_data.append({
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": birthday.strftime('%Y-%m-%d'),
            "Age": age,
            "Relatives": relatives
        })

# Output the JSON array
import json
print(json.dumps(output_data, indent=2))