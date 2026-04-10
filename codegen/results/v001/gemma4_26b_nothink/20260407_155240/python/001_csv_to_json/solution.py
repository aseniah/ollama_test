import csv
import json
import sys

def main():
    input_file = 'input/data.csv'
    output_data = []

    try:
        with open(input_file, mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                try:
                    # Create the object with specific type casting
                    obj = {
                        "Name": str(row["Name"]),
                        "Age": int(row["Age"]),
                        "Email": str(row["Email"]),
                        "Score": float(row["Score"])
                    }
                    output_data.append(obj)
                except (ValueError, KeyError) as e:
                    # Skip rows with malformed data or missing columns
                    continue

        # Output the JSON array to stdout
        json.dump(output_data, sys.stdout)

    except FileNotFoundError:
        # If file doesn't exist, output an empty array to maintain valid JSON output requirement
        sys.stdout.write("[]")
    except Exception:
        # In case of any other error, output an empty array
        sys.stdout.write("[]")

if __name__ == "__main__":
    main()