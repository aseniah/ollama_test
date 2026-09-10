import csv
import json
import sys

def main():
    input_file = 'input/data.csv'
    data = []

    try:
        with open(input_file, mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # Construct the object with specified type conversions
                item = {
                    "Name": row['Name'],
                    "Age": int(row['Age']),
                    "Email": row['Email'],
                    "Score": float(row['Score'])
                }
                data.append(item)
        
        # Output the resulting list as a JSON array to stdout
        sys.stdout.write(json.dumps(data))
        
    except FileNotFoundError:
        # If the file is missing, we output an empty array to remain valid JSON
        sys.stdout.write(json.dumps([]))
    except Exception:
        # Generic handler to ensure we don't output non-JSON text on crash
        sys.exit(1)

if __name__ == "__main__":
    main()