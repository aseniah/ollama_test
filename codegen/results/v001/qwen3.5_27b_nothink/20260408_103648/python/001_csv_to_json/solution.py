import csv
import json
import sys

def process_csv():
    data = []
    try:
        with open('input/data.csv', 'r', newline='', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                try:
                    entry = {
                        "Name": str(row['Name']),
                        "Age": int(row['Age']),
                        "Email": str(row['Email']),
                        "Score": float(row['Score'])
                    }
                    data.append(entry)
                except (ValueError, KeyError) as e:
                    # Skip invalid rows to ensure robustness, though problem implies valid input
                    continue
    except FileNotFoundError:
        sys.stderr.write("Error: input/data.csv not found.\n")
        sys.exit(1)
    except Exception as e:
        sys.stderr.write(f"Error reading file: {e}\n")
        sys.exit(1)

    # Output valid JSON to stdout
    print(json.dumps(data))

if __name__ == "__main__":
    process_csv()