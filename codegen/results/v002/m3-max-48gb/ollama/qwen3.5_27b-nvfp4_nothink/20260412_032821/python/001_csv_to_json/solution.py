import csv
import json

# Path to the input CSV file
input_file = "input/data.csv"

data_list = []

with open(input_file, mode='r', newline='', encoding='utf-8') as file:
    reader = csv.DictReader(file)
    for row in reader:
        record = {
            "Name": row["Name"],
            "Age": int(row["Age"]),
            "Email": row["Email"],
            "Score": float(row["Score"])
        }
        data_list.append(record)

# Output the JSON array to stdout
print(json.dumps(data_list))