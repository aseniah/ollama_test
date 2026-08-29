import csv
import json
import sys

with open("input/data.csv", newline="", encoding="utf-8") as f:
    rows = [
        {
            "Name": row["Name"].strip(),
            "Age": int(row["Age"].strip()),
            "Email": row["Email"].strip(),
            "Score": float(row["Score"].strip()),
        }
        for row in csv.DictReader(f)
    ]

sys.stdout.write(json.dumps(rows))