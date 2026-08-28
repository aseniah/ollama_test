import csv
import json
import sys

with open("input/data.csv", newline="", encoding="utf-8-sig") as f:
    reader = csv.DictReader(f)
    rows = []
    for row in reader:
        rows.append({
            "Name": row["Name"].strip(),
            "Age": int(row["Age"].strip()),
            "Email": row["Email"].strip(),
            "Score": float(row["Score"].strip())
        })

    sys.stdout.write(json.dumps(rows, separators=(",", ":")))