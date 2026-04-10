import csv
import json

with open('input/data.csv', 'r') as f:
    reader = csv.DictReader(f)
    rows = []
    for row in reader:
        rows.append({
            "Name": row["name"].strip(),
            "Age": int(row["age"]),
            "Email": row["email"].strip(),
            "Score": float(row["score"])
        })

print(json.dumps(rows))