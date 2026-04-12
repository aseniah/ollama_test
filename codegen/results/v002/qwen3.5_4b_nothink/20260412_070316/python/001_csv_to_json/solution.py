import csv
import json

# Read the CSV file
with open('input/data.csv', 'r') as csvfile:
    reader = csv.DictReader(csvfile)
    data = []
    
    # Process each row
    for row in reader:
        obj = {
            "Name": str(row["Name"]).strip(),
            "Age": int(row["Age"]),
            "Email": str(row["Email"]).strip(),
            "Score": float(row["Score"])
        }
        data.append(obj)

# Output JSON to stdout
print(json.dumps(data))