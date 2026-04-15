import csv
import json
import sys

def solve():
    input_file = 'input/data.csv'
    data = []

    try:
        with open(input_file, mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # Construct the object with specific types
                entry = {
                    "Name": str(row["Name"]),
                    "Age": int(row["Age"]),
                    "Email": str(row["Email"]),
                    "Score": float(row["Score"])
                }
                data.append(entry)
        
        # Output the JSON array to stdout
        sys.stdout.write(json.dumps(data, indent=2))
    except FileNotFoundError:
        # In a real scenario, we might handle this, 
        # but per requirements, we focus on the successful conversion.
        pass

if __name__ == "__main__":
    solve()