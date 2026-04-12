import csv
import json

# Read the CSV file and convert to list of dictionaries with appropriate types
with open('input/data.csv', 'r') as f:
    reader = csv.DictReader(f)
    
    data = []
    for row in reader:
        data.append({
            'Name': row['Name'],
            'Age': int(row['Age']),
            'Email': row['Email'],
            'Score': float(row['Score'])
        })

# Output as JSON to stdout
print(json.dumps(data))