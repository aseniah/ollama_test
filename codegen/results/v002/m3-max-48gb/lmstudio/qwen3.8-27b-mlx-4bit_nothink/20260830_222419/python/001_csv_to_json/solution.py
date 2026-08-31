import json
import csv

with open('input/data.csv', 'r') as f:
    reader = csv.DictReader(f)
    rows = []
    for row in reader:
        obj = {
            "Name": row['Name'],
            "Age": int(row['Age']),
            "Email": row['Email'],
            "Score": float(row['Score'])
        }
        rows.append(obj)

print(json.dumps(rows))