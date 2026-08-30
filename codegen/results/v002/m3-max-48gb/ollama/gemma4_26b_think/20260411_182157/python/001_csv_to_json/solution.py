import csv
import json
import sys

def main():
    """
    Reads a CSV file from 'input/data.csv', processes the rows into
    a list of dictionaries with specific types, and outputs the 
    result as a JSON array to stdout.
    """
    input_file_path = 'input/data.csv'
    output_data = []

    try:
        with open(input_file_path, mode='r', encoding='utf-8') as csvfile:
            # Use DictReader to automatically map the header row to dictionary keys
            reader = csv.DictReader(csvfile)
            
            for row in reader:
                # Create the object with the required type conversions:
                # Name: string, Age: integer, Email: string, Score: float
                try:
                    record = {
                        "Name": str(row["Name"]),
                        "Age": int(row["Age"]),
                        "Email": str(row["Email"]),
                        "Score": float(row["Score"])
                    }
                    output_data.append(record)
                except (ValueError, KeyError) as e:
                    # If a row is malformed or missing a key, we skip it to maintain 
                    # the validity of the final JSON array output.
                    continue
        
        # Serialize the list of dictionaries to a JSON string
        # We output to stdout as requested.
        json_output = json.dumps(output_data, indent=4)
        sys.stdout.write(json_output + '\n')

    except FileNotFoundError:
        # If the file is not found, we output nothing to adhere to the 
        # "Do not output anything else" constraint.
        pass
    except Exception:
        # Catch-all for unexpected errors to ensure no non-JSON text is printed.
        pass

if __name__ == "__main__":
    main()