import csv
import json
from datetime import date

# Define the target date for age calculation
target_date = date(2025, 7, 1)

# Load the expected output format to understand the structure
with open('input/expected_format.json', 'r') as f:
    expected_format = json.load(f)

# Determine the output format from the expected file
if 'output' in expected_format:
    for field in expected_format['output']:
        if 'age' in field:
            output_format = field['age']
            break
    else:
        output_format = field
else:
    output_format = None

# Infer output structure - assume we need a list with 'name' and 'age' keys
# Adjust if different structure inferred from expected format
output_list = []

# Read input CSV
with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    
    for row in reader:
        # Calculate age as of July 1, 2025
        birth_date = date(int(row['birth_date'][0:4]), int(row['birth_date'][5:7]), int(row['birth_date'][8:]))
        
        # Calculate years difference
        age = target_date.year - birth_date.year
        
        # Adjust for whether birthday has occurred yet this year
        if (target_date.month, target_date.day) < (birth_date.month, birth_date.day):
            age -= 1
        
        # Ensure age doesn't go below zero
        if age < 0:
            age = 0
        
        output_list.append({
            'name': row['name'],
            'age': age
        })

# Output as JSON array
print(json.dumps(output_list))