import csv
import json
import sys

def solve():
    input_file = 'input/data.csv'
    data_list = []

    try:
        with open(input_file, mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # Construct the object with specific type conversions
                entry = {
                    "Name": row['Name'],
                    "Age": int(row['Age']),
                    "Email": row['Email'],
                    "Score": float(row['Score'])
                }
                data_list.append(entry)
        
        # Output the list as a JSON array to stdout
        sys.stdout.write(json.dumps(data_list))
        
    except FileNotFoundError:
        # In a real scenario, you might handle this differently, 
        # but for this task, we assume the file exists as per the prompt.
        pass

if __name__ == "__main__":
    solve()