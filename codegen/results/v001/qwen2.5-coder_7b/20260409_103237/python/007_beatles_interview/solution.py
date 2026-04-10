import csv
from datetime import datetime

# Define the input and expected files
input_file = 'input/input.csv'
expected_format_file = 'input/expected_format.json'

# Read the expected format to understand the structure of the output
with open(expected_format_file, 'r') as f:
    expected_output = f.read()

# Parse the expected output to get the keys we need in our JSON array
import json
expected_keys = list(json.loads(expected_output)[0].keys())

# Initialize an empty list to store the results
results = []

# Read the input CSV file
with open(input_file, 'r') as csvfile:
    csvreader = csv.DictReader(csvfile)
    
    for row in csvreader:
        # Calculate age as of July 1, 2025
        dob = datetime.strptime(row['dob'], '%Y-%m-%d')
        target_date = datetime(2025, 7, 1)
        age = target_date.year - dob.year - ((target_date.month, target_date.day) < (dob.month, dob.day))
        
        # Create a dictionary with the required keys and values
        result_row = {
            'name': row['name'],
            'gender': row['gender'],
            'age': age,
            'email': row['email']
        }
        
        # Append the result to the list
        results.append(result_row)

# Convert the list of dictionaries to a JSON array and print it
import json
print(json.dumps(results, indent=2))