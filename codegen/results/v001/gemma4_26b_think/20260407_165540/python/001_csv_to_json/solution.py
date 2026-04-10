import csv
import json
import sys

def main():
    """
    Reads 'input/data.csv', parses the content, and outputs a 
    JSON array of objects to stdout.
    """
    input_file = 'input/data.csv'
    output_data = []

    try:
        with open(input_file, mode='r', encoding='utf-8-sig') as csvfile:
            # DictReader uses the first row as field names
            reader = csv.DictReader(csvfile)
            
            for row in reader:
                try:
                    # Construct the object with explicit type casting
                    # We assume the CSV headers match the required field names
                    obj = {
                        "Name": str(row["Name"]),
                        "Age": int(row["Age"]),
                        "Email": str(row["Email"]),
                        "Score": float(row["Score"])
                    }
                    output_data.append(obj)
                except (ValueError, KeyError):
                    # Skip rows that are malformed or missing required columns
                    continue

        # Output the resulting list as a JSON array to stdout
        # Using json.dumps ensures the output is valid JSON
        sys.stdout.write(json.dumps(output_data, indent=2))
        
    except FileNotFoundError:
        # If the file does not exist, we output an empty JSON array 
        # to ensure the output remains valid JSON.
        sys.stdout.write("[]")
    except Exception:
        # In case of any other unexpected error, output an empty array 
        # to maintain the requirement of valid JSON output.
        sys.stdout.write("[]")

if __name__ == "__main__":
    main()