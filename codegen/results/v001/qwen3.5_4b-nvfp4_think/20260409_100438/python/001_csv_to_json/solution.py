import json
import csv

with open('input/data.csv', 'r') as f:
    reader = csv.DictReader(f)
    
    if not reader.fieldnames:
        print(json.dumps({}))
        return
    
    rows = []
    for row in reader:
        row_data = {
            "Name": row["Name"],
            "Age": int(row["Age"]),
            "Email": row["Email"],
            "Score": float(row["Score"])
        }
        rows.append(row_data)
    
    print(json.dumps(rows))