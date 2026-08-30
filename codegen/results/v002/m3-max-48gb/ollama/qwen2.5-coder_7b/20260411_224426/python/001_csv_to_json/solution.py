import csv
import json

# Read the CSV file
with open('input/data.csv', newline='') as csvfile:
    reader = csv.DictReader(csvfile)
    
    # Convert each row to a dictionary with the specified fields
    data = []
    for row in reader:
        data.append({
            "Name": row["Name"],
            "Age": int(row["Age"]),
            "Email": row["Email"],
            "Score": float(row["Score"])
        })

# Output the JSON array
print(json.dumps(data, indent=2))