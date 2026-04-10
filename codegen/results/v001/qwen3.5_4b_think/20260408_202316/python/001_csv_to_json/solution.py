import csv
import json

# Read CSV file and convert to JSON array
with open('input/data.csv', 'r', encoding='utf-8') as f:
    reader = csv.DictReader(f)
    json_array = [
        {
            "Name": row["name"],
            "Age": int(row["age"]),
            "Email": row["email"],
            "Score": float(row["score"])
        }
        for row in reader
    ]

# Output JSON to stdout
print(json.dumps(json_array))