import json
import csv

try:
    with open('input/data.csv', 'r', encoding='utf-8') as f:
        reader = csv.DictReader(f)
        data = [row for row in reader if all(key in row for key in ['name', 'age', 'email', 'score'])]
        # Clean and format fields
        formatted_data = [{
            "Name": row['name'].strip(),
            "Age": int(row['age']),
            "Email": row['email'].strip(),
            "Score": float(row['score'])
        } for row in data if row]  # Skip empty rows
except FileNotFoundError:
    print(json.dumps([]))  # Empty list on file not found