import csv
import json
import sys

def process_csv_to_json():
    input_file = 'input/data.csv'
    data_list = []

    try:
        with open(input_file, mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # Construct the object with specific type casting
                entry = {
                    "Name": str(row['Name']),
                    "Age": int(row['Age']),
                    "Email": str(row['Email']),
                    "Score": float(row['Score'])
                }
                data_list.append(entry)
        
        # Output valid JSON array to stdout
        sys.stdout.write(json.dumps(data_list, indent=2))
    except FileNotFoundError:
        # Based on the prompt, we assume the file exists, 
        # but we handle it gracefully just in case.
        pass

if __name__ == "__main__":
    process_csv_to_json()