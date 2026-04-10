import csv
import json

try:
    with open('input/data.csv', mode='r', encoding='utf-8') as file:
        reader = csv.DictReader(file)
        result = []
        for row in reader:
            obj = {
                "Name": row.get("Name", ""),
                "Age": int(row.get("Age", 0)),
                "Email": row.get("Email", ""),
                "Score": float(row.get("Score", 0.0))
            }
            result.append(obj)
        
        print(json.dumps(result, indent=None, separators=(',', ':')))
except FileNotFoundError:
    print("[]")
except Exception as e:
    print("[]")