import csv
import json
from datetime import datetime

# Read input CSV file
input_csv_path = 'input/input.csv'
output_json_path = 'output/output.json'

with open(input_csv_path, mode='r', newline='') as csvfile:
    reader = csv.DictReader(csvfile)
    data = list(reader)

# Calculate ages as of July 1, 2025
current_date = datetime(2025, 7, 1)
output_data = []

for row in data:
    birthdate = datetime.strptime(row['BirthDate'], '%m/%d/%Y')
    age = current_date.year - birthdate.year - ((current_date.month, current_date.day) < (birthdate.month, birthdate.day))
    output_row = {
        "Name": row['Name'],
        "Age": age,
        "Gender": row['Gender'],
        "Address": row['Address']
    }
    output_data.append(output_row)

# Convert to JSON and write to stdout
json_output = json.dumps(output_data, indent=2)
print(json_output)