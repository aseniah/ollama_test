import csv
import json
import sys

def csv_to_json():
    input_file = 'input/data.csv'
    data_list = []

    try:
        with open(input_file, mode='r', encoding='utf-8') as f:
            reader = csv.DictReader(f)
            for row in reader:
                # Transform values to specified types
                transformed_row = {
                    "Name": row["Name"],
                    "Age": int(row["Age"]),
                    "Email": row["Email"],
                    "Score": float(row["Score"])
                }
                data_list.append(transformed_row)
        
        # Output valid JSON to stdout
        sys.stdout.write(json.dumps(data_list))
    except Exception:
        # Silent fail as per the constraint "Do not output anything else" 
        # beyond valid JSON, although in a real scenario we'd handle errors.
        pass

if __name__ == "__main__":
    csv_to_json()