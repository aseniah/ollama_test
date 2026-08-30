import csv
import json
import sys

def solve():
    data_list = []
    try:
        with open('input/data.csv', mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # Construct the object with correct types
                entry = {
                    "Name": row["Name"],
                    "Age": int(row["Age"]),
                    "Email": row["Email"],
                    "Score": float(row["Score"])
                }
                data_list.append(entry)
        
        # Output the JSON array to stdout
        sys.stdout.write(json.dumps(data_list))
    except Exception as e:
        # Since the prompt asks for valid JSON and nothing else, 
        # we typically avoid printing errors to stdout unless required.
        pass

if __name__ == "__main__":
    solve()