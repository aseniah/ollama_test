import csv
import json

result = []

with open('input/data.csv', 'r', newline='') as f:
    reader = csv.DictReader(f)
    for row in reader:
        obj = {
            "Name": row["Name"],
            "Age": int(row["Age"]),
            "Email": row["Email"],
            "Score": float(row["Score"])
        }
        result.append(obj)

print(json.dumps(result))