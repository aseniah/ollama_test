import csv
import json

with open('input/data.csv', 'r') as f:
    reader = csv.DictReader(f)
    data = []
    for row in reader:
        entry = {
            "Name": row['Name'],
            "Age": int(row['Age']),
            "Email": row['Email'],
            "Score": float(row['Score'])
        }
        data.append(entry)

print(json.dumps(data))