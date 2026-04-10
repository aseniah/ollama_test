import csv
import json

with open('input/data.csv', 'r', newline='', encoding='utf-8') as f:
    reader = csv.DictReader(f)
    rows_data = []
    for row in reader:
        data = {
            "Name": row["name"],
            "Age": int(row["age"]),
            "Email": row["email"],
            "Score": float(row["score"])
        }
        rows_data.append(data)
    
    print(json.dumps(rows_data))