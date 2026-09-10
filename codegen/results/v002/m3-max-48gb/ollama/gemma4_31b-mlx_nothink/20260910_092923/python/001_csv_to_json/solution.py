import csv
import json
import sys

def convert_csv_to_json(input_file):
    data = []
    try:
        with open(input_file, mode='r', encoding='utf-8') as f:
            reader = csv.DictReader(f)
            for row in reader:
                # Convert types as specified in the requirements
                entry = {
                    "Name": str(row["Name"]),
                    "Age": int(row["Age"]),
                    "Email": str(row["Email"]),
                    "Score": float(row["Score"])
                }
                data.append(entry)
        
        # Output as valid JSON array to stdout
        print(json.dumps(data, indent=2))
        
    except FileNotFoundError:
        # Based on requirements, only valid JSON should be output.
        # In a real scenario, you might handle this differently, 
        # but here we follow the "Do not output anything else" constraint.
        pass

if __name__ == "__main__":
    convert_csv_to_json('input/data.csv')