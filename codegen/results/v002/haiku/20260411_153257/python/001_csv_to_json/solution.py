import csv
import json
import sys

rows = []
with open('input/data.csv', newline='') as f:
    reader = csv.DictReader(f)
    for row in reader:
        rows.append({
            'Name': row['Name'],
            'Age': int(row['Age']),
            'Email': row['Email'],
            'Score': float(row['Score'])
        })

print(json.dumps(rows))
