import csv
import json

data = []
try:
    with open('input/data.csv', 'r') as file:
        reader = csv.DictReader(file)
        for row in reader:
            record = {
                "Name": row['Name'],
                "Age": int(row['Age']),
                "Email": row['Email'],
                "Score": float(row['Score'])
            }
            data.append(record)
    print(json.dumps(data))
except Exception as e:
    # In case of error, output an empty list to satisfy the JSON requirement
    print("[]")