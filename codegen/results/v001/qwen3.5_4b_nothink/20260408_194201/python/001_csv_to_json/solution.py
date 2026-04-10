import csv
import json

try:
    with open('input/data.csv', 'r', encoding='utf-8') as f:
        reader = csv.DictReader(f)
        data = []
        for row in reader:
            entry = {
                "Name": str(row.get("Name", "")),
                "Age": int(row.get("Age", 0)),
                "Email": str(row.get("Email", "")),
                "Score": float(row.get("Score", 0.0))
            }
            data.append(entry)
        print(json.dumps(data))
except FileNotFoundError:
    pass