import csv
import json
import sys

def main():
    """
    Reads a CSV file from 'input/data.csv', parses the content into a list of 
    dictionaries with specific types, and outputs a JSON array to stdout.
    """
    input_file_path = 'input/data.csv'
    json_output = []

    try:
        with open(input_file_path, mode='r', encoding='utf-8') as csvfile:
            # DictReader uses the first row as field names (keys)
            reader = csv.DictReader(csvfile)
            
            for row in reader:
                # Extract and convert types as specified in the requirements
                try:
                    processed_row = {
                        "Name": row["Name"],
                        "Age": int(row["Age"]),
                        "Email": row["Email"],
                        "Score": float(row["Score"])
                    }
                    json_output.append(processed_row)
                except (ValueError, KeyError) as e:
                    # If a row has invalid data or missing columns, we skip it
                    # to ensure the final JSON remains valid according to specs.
                    continue
        
        # Output the final JSON array to stdout
        # indent=4 is used for readability, but valid JSON remains standard.
        print(json.dumps(json_output, indent=4))

    except FileNotFoundError:
        # If the file is not found, we exit silently to avoid polluting stdout 
        # with error messages that would break the "valid JSON only" constraint.
        sys.exit(1)
    except Exception:
        sys.exit(1)

if __name__ == "__main__":
    main()