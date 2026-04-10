import csv
import json
import sys

# Read the CSV file and convert to JSON
with open('input/data.csv', 'r') as csvfile:
    reader = csv.DictReader(csvfile)
    data = []
    
    for row in reader:
        data.append({
            "Name": row.get('Name', ''),
            "Age": int(row.get('Age', 0)),
            "Email": row.get('Email', ''),
            "Score": float(row.get('Score', 0.0))
        })

# Output as JSON to stdout
print(json.dumps(data))