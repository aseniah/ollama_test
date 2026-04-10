import csv
import json
import sys

def main():
    input_file = 'input/data.csv'
    data_list = []

    try:
        with open(input_file, mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                try:
                    # Construct the object with explicit type casting
                    obj = {
                        "Name": str(row["Name"]),
                        "Age": int(row["Age"]),
                        "Email": str(row["Email"]),
                        "Score": float(row["Score"])
                    }
                    data_list.append(obj)
                except (ValueError, KeyError) as e:
                    # Skip rows with malformed data or missing columns
                    continue

        # Output the JSON array to stdout
        json.dump(data_list, sys.stdout, indent=None)
        
    except FileNotFoundError:
        # If file is missing, output an empty array to maintain valid JSON format
        sys.stdout.write("[]")
    except Exception:
        # In case of any other error, output an empty array
        sys.stdout.write("[]")

if __name__ == "__main__":
    main()