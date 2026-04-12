import csv
from datetime import datetime

# Define the transformation rules and output format based on expected_format.json
def transform_data(input_csv):
    output = []
    
    with open(input_csv, newline='') as csvfile:
        reader = csv.DictReader(csvfile)
        
        for row in reader:
            # Calculate age as of July 1, 2025
            birthdate = datetime.strptime(row['Birthday'], '%m/%d/%Y')
            current_date = datetime(2025, 7, 1)
            age = current_date.year - birthdate.year - ((current_date.month, current_date.day) < (birthdate.month, birthdate.day))
            
            # Prepare relatives data
            relatives = []
            if row['Father']:
                relatives.append({'FirstName': row['Father'], 'LastName': '', 'Relationship': 'Father'})
            if row['Mother']:
                relatives.append({'FirstName': row['Mother'], 'LastName': '', 'Relationship': 'Mother'})
            if row['Brother']:
                relatives.append({'FirstName': row['Brother'], 'LastName': '', 'Relationship': 'Brother'})
            if row['Sister']:
                relatives.append({'FirstName': row['Sister'], 'LastName': '', 'Relationship': 'Sister'})
            
            # Add the transformed data to the output list
            output.append({
                "FirstName": row['Name'].split()[0],
                "LastName": row['Name'].split()[1] if len(row['Name'].split()) > 1 else '',
                "Birthday": row['Birthday'],
                "Age": age,
                "Relatives": relatives
            })
    
    return output

# Read the input CSV and transform the data
output_json = transform_data('input/input.csv')

# Output the JSON array to stdout
import json
print(json.dumps(output_json, indent=2))