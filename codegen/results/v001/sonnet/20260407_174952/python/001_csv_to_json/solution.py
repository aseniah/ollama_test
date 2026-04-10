import csv
import json

with open("input/data.csv", newline="") as f:
    reader = csv.DictReader(f)
    rows = [
        {
            "Name": row["Name"],
            "Age": int(row["Age"]),
            "Email": row["Email"],
            "Score": float(row["Score"]),
        }
        for row in reader
    ]

print(json.dumps(rows))
