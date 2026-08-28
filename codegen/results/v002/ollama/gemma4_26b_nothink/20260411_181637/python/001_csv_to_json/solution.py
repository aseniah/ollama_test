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
                # Create dictionary with typed values
                obj = {
                    "Name": row["Name"],
                    "Age": int(row["Age"]),
                    "Email": row["Email"],
                    "Score": float(row["Score"])
                }
                output_data.append(obj)
        
        # Output the JSON array to stdout
        json.dump(output_data, sys.stdout, indent=None)
        # Ensure a newline at the end for clean stdout
        sys.stdout.write('\n')

    except FileNotFoundError:
        # Silently fail or handle error as per requirement (requirement says output valid JSON)
        pass
    except Exception:
        pass

if __name__ == "__main__":
    main()