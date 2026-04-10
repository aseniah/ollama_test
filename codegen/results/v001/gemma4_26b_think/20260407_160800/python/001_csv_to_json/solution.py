import csv
import json
import sys

def main():
    file_path = 'input/data.csv'
    json_output = []

    try:
        with open(file_path, mode='r', encoding='utf-8-sig') as csvfile:
            # DictReader uses the first row of the CSV as dictionary keys
            reader = csv.DictReader(csvfile)
            
            for row in reader:
                # Create a new dictionary with the specific required types
                # We map the CSV columns to the requested output keys
                # and cast Age to int and Score to float.
                record = {
                    "Name": row['Name'],
                    "Age": int(row['Age']),
                    "Email": row['Email'],
                    "Score": float(row['Score'])
                }
                json_output.append(record)
        
        # Output the array as a valid JSON string to stdout
        # No extra text is printed to ensure only valid JSON is sent to stdout
        sys.stdout.write(json.dumps(json_output))

    except (FileNotFoundError, KeyError, ValueError):
        # In case of file error or parsing error, we do not output anything
        # to ensure we don't pollute stdout with non-JSON error messages.
        pass

if __name__ == "__main__":
    main()