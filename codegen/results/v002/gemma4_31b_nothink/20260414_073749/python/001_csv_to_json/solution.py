import csv
import json
import sys

def solve():
    input_filename = 'input/data.csv'
    data_list = []

    try:
        with open(input_filename, mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # Construct the object with required type conversions
                obj = {
                    "Name": row["Name"],
                    "Age": int(row["Age"]),
                    "Email": row["Email"],
                    "Score": float(row["Score"])
                }
                data_list.append(obj)
    except FileNotFoundError:
        # Depending on requirements, you might want to output an empty array 
        # or handle the error. Given the prompt, we assume the file exists.
        pass

    # Output the JSON array to stdout
    sys.stdout.write(json.dumps(data_list))

if __name__ == "__main__":
    solve()