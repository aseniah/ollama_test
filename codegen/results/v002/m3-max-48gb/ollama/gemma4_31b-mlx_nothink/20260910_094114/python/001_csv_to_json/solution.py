import csv
import json
import sys

def solve():
    input_file = 'input/data.csv'
    data_list = []

    try:
        with open(input_file, mode='r', encoding='utf-8') as f:
            reader = csv.DictReader(f)
            for row in reader:
                # Convert types as specified
                item = {
                    "Name": str(row["Name"]),
                    "Age": int(row["Age"]),
                    "Email": str(row["Email"]),
                    "Score": float(row["Score"])
                }
                data_list.append(item)
        
        # Output valid JSON array to stdout
        sys.stdout.write(json.dumps(data_list, indent=2))
    except FileNotFoundError:
        pass

if __name__ == "__main__":
    solve()