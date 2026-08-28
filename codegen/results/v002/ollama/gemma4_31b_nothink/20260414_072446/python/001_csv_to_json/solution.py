import csv
import json
import sys

def convert_csv_to_json():
    input_file = 'input/data.csv'
    data = []

    try:
        with open(input_file, mode='r', encoding='utf-8') as f:
            reader = csv.DictReader(f)
            for row in reader:
                # Construct the object with specific type casting
                item = {
                    "Name": str(row["Name"]),
                    "Age": int(row["Age"]),
                    "Email": str(row["Email"]),
                    "Score": float(row["Score"])
                }
                data.append(item)
        
        # Output valid JSON array to stdout
        sys.stdout.write(json.dumps(data, indent=2))
        
    except FileNotFoundError:
        # Depending on requirements, you can handle errors or leave empty
        pass

if __name__ == "__main__":
    convert_csv_to_json()