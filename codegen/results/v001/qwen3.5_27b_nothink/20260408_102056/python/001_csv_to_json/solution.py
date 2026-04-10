import csv
import json

try:
    with open('input/data.csv', 'r', newline='', encoding='utf-8') as csvfile:
        reader = csv.DictReader(csvfile)
        data = []
        for row in reader:
            data.append({
                "Name": row['Name'],
                "Age": int(row['Age']),
                "Email": row['Email'],
                "Score": float(row['Score'])
            })
        print(json.dumps(data))
except Exception as e:
    # In a real scenario, handle the error appropriately, 
    # but the prompt asks for JSON output only. 
    # If an error occurs, we cannot output valid JSON without data, 
    # so we exit silently or with an error message if debugging.
    # Given the strict constraint "Do not output anything else", 
    # we will not print error messages to stdout.
    import sys
    sys.exit(1)