import json
import csv

with open('input/data.csv', 'r', encoding='utf-8') as f:
    reader = csv.DictReader(f)
    data = list(reader)

output = []
for row in data:
    output.append({
        "Name": row["Name"],
        "Age": int(row["Age"]),
        "Email": row["Email"],
        "Score": float(row["Score"])
    })

print(json.dumps(output))