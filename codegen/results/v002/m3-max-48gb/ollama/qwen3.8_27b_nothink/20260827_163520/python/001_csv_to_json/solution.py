import csv
import json

data = []
with open('input/data.csv', 'r', newline='') as f:
    reader = csv.DictReader(f)
    for row in reader:
        data.append({
            'Name': row['Name'].strip(),
            'Age': int(row['Age'].strip()),
            'Email': row['Email'].strip(),
            'Score': float(row['Score'].strip())
        })

print(json.dumps(data, indent=2))