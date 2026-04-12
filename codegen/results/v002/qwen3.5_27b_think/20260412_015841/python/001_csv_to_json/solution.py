import csv
import json

# Read the CSV file
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

# Output as JSON
print(json.dumps(data))