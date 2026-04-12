import csv
import json

filename = 'input/data.csv'

output = []

with open(filename, 'r', newline='', encoding='utf-8') as csvfile:
    reader = csv.DictReader(csvfile)
    for row in reader:
        record = {
            "Name": row['Name'],
            "Age": int(row['Age']),
            "Email": row['Email'],
            "Score": float(row['Score'])
        }
        output.append(record)

print(json.dumps(output, ensure_ascii=False))