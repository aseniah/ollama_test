import csv
import json

with open('input/data.csv', 'r', encoding='utf-8') as file:
    reader = csv.DictReader(file)
    rows = []
    for row in reader:
        rows.append({
            "Name": row["Name"],
            "Age": int(row["Age"]),
            "Email": row["Email"],
            "Score": float(row["Score"])
        })

print(json.dumps(rows, indent=2))