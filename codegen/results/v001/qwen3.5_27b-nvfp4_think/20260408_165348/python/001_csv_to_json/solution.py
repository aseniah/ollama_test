import csv
import json

# Read the CSV file
data = []
with open('input/data.csv', 'r') as f:
    reader = csv.DictReader(f)
    for row in reader:
        data.append({
            "Name": row['Name'],
            "Age": int(row['Age']),
            "Email": row['Email'],
            "Score": float(row['Score'])
        })

# Output JSON to stdout
print(json.dumps(data, indent=2))