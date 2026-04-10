import csv
import json

data = []
with open('input/data.csv', 'r') as file:
    reader = csv.DictReader(file)
    for row in reader:
        data.append({
            'Name': row['Name'],
            'Age': int(row['Age']),
            'Email': row['Email'],
            'Score': float(row['Score'])
        })

print(json.dumps(data))