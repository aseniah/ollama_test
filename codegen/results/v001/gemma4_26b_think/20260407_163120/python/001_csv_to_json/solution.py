import csv
import json
import sys

def main():
    input_path = 'input/data.csv'
    data_list = []

    try:
        with open(input_path, mode='r', encoding='utf-8') as csvfile:
            # DictReader uses the first row of the CSV as dictionary keys
            reader = csv.DictReader(csvfile)
            
            for row in reader:
                # Create the object with explicit type casting as requested
                item = {
                    "Name": str(row["Name"]),
                    "Age": int(row["Age"]),
                    "Email": str(row["Email"]),
                    "Score": float(row["Score"])
                }
                data_list.append(item)
        
        # Serialize the list to a JSON array and write directly to stdout
        json.dump(data_list, sys.stdout)

    except (FileNotFoundError, KeyError, ValueError):
        # If file is missing or data types are incompatible, 
        # we exit without printing extra text to respect the "Do not output anything else" constraint.
        sys.exit(1)
    except Exception:
        sys.exit(1)

if __name__ == "__main__":
    main()