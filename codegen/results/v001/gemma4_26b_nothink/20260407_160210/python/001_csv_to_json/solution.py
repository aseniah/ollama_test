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
                    # Transform and type-cast fields according to requirements
                    obj = {
                        "Name": str(row["Name"]),
                        "Age": int(row["Age"]),
                        "Email": str(row["Email"]),
                        "Score": float(row["Score"])
                    }
                    output_data.append(obj)
                except (ValueError, KeyError) as e:
                    # Skip rows that do not match the expected types or schema
                    continue

        # Write the JSON array to stdout
        json.dump(output_data, sys.stdout)
        
    except FileNotFoundError:
        # If file doesn't exist, output an empty array to maintain valid JSON requirement
        sys.stdout.write("[]")
    except Exception:
        # Fallback for any other unexpected errors
        sys.stdout.write("[]")

if __name__ == "__main__":
    main()