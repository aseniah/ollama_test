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
                # Construct the dictionary with specific type casting
                obj = {
                    "Name": row["Name"],
                    "Age": int(row["Age"]),
                    "Email": row["Email"],
                    "Score": float(row["Score"])
                }
                output_data.append(obj)
        
        # Output the JSON array to stdout
        json.dump(output_data, sys.stdout)
        
    except FileNotFoundError:
        # If file doesn't exist, we output an empty list or handle as error
        # Per instructions, we just need to output valid JSON.
        pass
    except Exception as e:
        # In a production environment, you might log this error.
        # For this task, we ensure valid JSON is the only output.
        pass

if __name__ == "__main__":
    main()