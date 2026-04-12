import csv
import json

csv_file_path = 'input/data.csv'
output = []

with open(csv_file_path, mode='r', newline='', encoding='utf-8') as file:
    reader = csv.DictReader(file)
    for row in reader:
        obj = {
            "Name": row['Name'],
            "Age": int(row['Age']),
            "Email": row['Email'],
            "Score": float(row['Score'])
        }
        output.append(obj)

print(json.dumps(output))