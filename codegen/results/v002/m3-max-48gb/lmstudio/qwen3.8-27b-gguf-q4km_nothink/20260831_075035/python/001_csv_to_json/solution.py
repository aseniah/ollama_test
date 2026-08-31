import csv
import json

with open('input/data.csv', 'r', newline='') as f:
    reader = csv.DictReader(f)
    result = []
    for row in reader:
        result.append({
            'Name': row['Name'],
            'Age': int(row['Age']),
            'Email': row['Email'],
            'Score': float(row['Score'])
        })

print(json.dumps(result))