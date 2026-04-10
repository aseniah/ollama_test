import csv
import json

# Read the CSV file
with open('input/data.csv', 'r') as csvfile:
    reader = csv.DictReader(csvfile)
    
    # Prepare to collect all rows
    rows = []
    
    # Iterate over each row
    for row in reader:
        # Create a dictionary from the row
        row_dict = {
            'Name': row['Name'],
            'Age': int(row['Age']),
            'Email': row['Email'],
            'Score': float(row['Score'])
        }
        # Append the dictionary to the rows list
        rows.append(row_dict)

# Convert the list of dictionaries to a JSON array
json_array = json.dumps(rows, indent=4)

# Output the JSON array
print(json_array)