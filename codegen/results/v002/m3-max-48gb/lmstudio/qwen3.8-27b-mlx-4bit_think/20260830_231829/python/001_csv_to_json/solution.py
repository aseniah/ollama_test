import csv
import json

results = []
with open("input/data.csv", "r") as f:
    reader = csv.DictReader(f)
    for row in reader:
        results.append({
            "Name": row["Name"],
            "Age": int(row["Age"]),
            "Email": row["Email"],
            "Score": float(row["Score"])
        })

print(json.dumps(results, indent=2))