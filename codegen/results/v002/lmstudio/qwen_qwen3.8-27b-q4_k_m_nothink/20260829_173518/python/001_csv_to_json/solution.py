import csv
import json

output = []

with open('input/data.csv', 'r') as f:
    reader = csv.DictReader(f)
    for row in reader:
        obj = {
            "Name": row['Name'],
            "Age": int(row['Age']),
            "Email": row['Email'],
            "Score": float(row['Score'])
        }
        output.append(obj)

print(json.dumps(output, indent=None))