import csv
import json

# Read the CSV file
with open('input/data.csv', 'r') as f:
    reader = csv.DictReader(f)
    
    # Parse each row into a JSON object
    data = []
    for row in reader:
        record = {
            'Name': row['Name'],
            'Age': int(row['Age']),
            'Email': row['Email'],
            'Score': float(row['Score'])
        }
        data.append(record)
    
    # Output as JSON array to stdout
    print(json.dumps(data))