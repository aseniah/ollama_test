import csv
import json
import sys

def solve():
    data_list = []
    try:
        with open('input/data.csv', mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # Construct the object with specified types
                item = {
                    "Name": str(row["Name"]),
                    "Age": int(row["Age"]),
                    "Email": str(row["Email"]),
                    "Score": float(row["Score"])
                }
                data_list.append(item)
        
        # Output the JSON array to stdout
        sys.stdout.write(json.dumps(data_list))
    except FileNotFoundError:
        # In a real scenario, you might handle this, 
        # but the prompt assumes the file exists based on provided data.
        pass

if __name__ == "__main__":
    solve()