import csv
import json

# Define the input file path
input_file = 'input/data.csv'

# List to store the parsed data
data_rows = []

# Read the CSV file
with open(input_file, mode='r', encoding='utf-8') as file:
    reader = csv.DictReader(file)
    
    for row in reader:
        # Parse and convert data types
        person = {
            "Name": row["Name"],
            "Age": int(row["Age"]),
            "Email": row["Email"],
            "Score": float(row["Score"])
        }
        data_rows.append(person)

# Output the JSON array to stdout
print(json.dumps(data_rows))