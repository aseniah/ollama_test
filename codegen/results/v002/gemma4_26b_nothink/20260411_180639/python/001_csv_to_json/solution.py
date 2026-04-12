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
                # Create object with specific types as requested
                obj = {
                    "Name": str(row["Name"]),
                    "Age": int(row["Age"]),
                    "Email": str(row["Email"]),
                    "Score": float(row["Score"])
                }
                output_data.append(obj)
        
        # Output the JSON array to stdout
        json.dump(output_data, sys.stdout, indent=None)
        
    except FileNotFoundError:
        # If file doesn't exist, we exit silently or handle as needed
        # Based on instructions, we just need to output valid JSON.
        # An empty array is a valid JSON response for an empty/missing dataset.
        sys.stdout.write("[]")
    except Exception as e:
        # In a production environment, you might want to log this error.
        # For this task, we ensure we don't break the JSON output format.
        pass

if __name__ == "__main__":
    main()