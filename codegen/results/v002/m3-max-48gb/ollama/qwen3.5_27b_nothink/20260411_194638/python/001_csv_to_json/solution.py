import csv
import json
import os
import sys

def main():
    input_file = 'input/data.csv'

    # Ensure the file exists (optional safety check, though problem guarantees it)
    if not os.path.exists(input_file):
        sys.stderr.write(f"Error: File '{input_file}' not found.\n")
        sys.exit(1)

    result = []

    with open(input_file, mode='r', newline='', encoding='utf-8') as csvfile:
        reader = csv.DictReader(csvfile)
        for row in reader:
            record = {
                "Name": row["Name"],
                "Age": int(row["Age"]),
                "Email": row["Email"],
                "Score": float(row["Score"])
            }
            result.append(record)

    # Output valid JSON to stdout only
    print(json.dumps(result))

if __name__ == "__main__":
    main()