import csv
import json

# Read the CSV file
with open('input/data.csv', 'r') as f:
    reader = csv.DictReader(f)
    data = []

    for row in reader:
        # Convert Age to int, Score to float
        data.append({
            'Name': str(row['Name']),
            'Age': int(row['Age']),
            'Email': str(row['Email']),
            'Score': float(row['Score'])
        })

# Output as JSON array to stdout
print(json.dumps(data))