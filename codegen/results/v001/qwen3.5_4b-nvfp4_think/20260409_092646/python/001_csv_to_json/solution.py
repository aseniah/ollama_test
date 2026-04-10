import csv
import json

with open('input/data.csv', 'r') as f:
    reader = csv.DictReader(f)
    rows = [row for row in reader]
    output = {
        "Name": row["Name"],
        "Age": int(row["Age"]),
        "Email": row["Email"],
        "Score": float(row["Score"])
    }

print(json.dumps([output]))