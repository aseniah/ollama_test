import csv
import json
import sys

def main():
    """
    Reads data from input/data.csv and prints a JSON array to stdout.
    The CSV columns are mapped to specific types: Name (str), Age (int), Email (str), Score (float).
    """
    input_file = 'input/data.csv'
    json_output = []

    try:
        with open(input_file, mode='r', encoding='utf-8') as csvfile:
            # DictReader uses the first row as field names (headers)
            reader = csv.DictReader(csvfile)
            
            for row in reader:
                # Create a dictionary with the required types
                record = {
                    "Name": str(row["Name"]),
                    "Age": int(row["Age"]),
                    "Email": str(row["Email"]),
                    "Score": float(row["Score"])
                }
                json_output.append(record)
        
        # Output the list as a JSON array to stdout
        # indent=4 is used for readability, but valid JSON without it
        print(json.dumps(json_output, indent=4))

    except FileNotFoundError:
        # If the file is not found, we output nothing or handle as required.
        # Given the prompt, we assume the file exists.
        pass
    except KeyError as e:
        # Handle cases where CSV headers might be missing
        pass
    except ValueError as e:
        # Handle cases where type conversion (int/float) might fail
        pass

if __name__ == "__main__":
    main()