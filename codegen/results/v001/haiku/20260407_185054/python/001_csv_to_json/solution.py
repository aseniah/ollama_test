import csv
import json
import sys

data = []
with open('input/data.csv', 'r') as f:
    reader = csv.DictReader(f)
    for row in reader:
        data.append({
            'Name': row['Name'],
            'Age': int(row['Age']),
            'Email': row['Email'],
            'Score': float(row['Score'])
        })

json.dump(data, sys.stdout)
