import csv
import json
import sys

def main():
    input_file = 'input/data.csv'
    try:
        with open(input_file, mode='r', newline='', encoding='utf-8') as file:
            reader = csv.DictReader(file)
            result = []
            
            for row in reader:
                # Convert Age to integer and Score to float
                try:
                    person = {
                        "Name": row["Name"],
                        "Age": int(row["Age"]),
                        "Email": row["Email"],
                        "Score": float(row["Score"])
                    }
                    result.append(person)
                except (ValueError, KeyError) as e:
                    # In a real-world scenario, you might log this or handle it differently.
                    # For strict compliance with the prompt (only valid JSON), we skip bad rows 
                    # or the script would fail if data is malformed. 
                    # Assuming input data matches the schema described.
                    pass
            
            # Output strictly to stdout as a valid JSON array
            print(json.dumps(result, indent=2))

    except FileNotFoundError:
        sys.exit(1)
    except Exception as e:
        # If any other error occurs, ensure we don't output garbage text
        sys.exit(1)

if __name__ == "__main__":
    main()