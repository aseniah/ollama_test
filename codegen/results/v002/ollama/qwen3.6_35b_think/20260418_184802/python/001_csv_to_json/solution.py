import csv
import json
import os

def main():
    file_path = 'input/data.csv'
    
    # Check if file exists to prevent errors during testing
    if not os.path.exists(file_path):
        print("Error: File not found", file=sys.stderr)
        return

    data = []

    with open(file_path, mode='r', encoding='utf-8') as f:
        reader = csv.DictReader(f)
        for row in reader:
            obj = {
                "Name": row["Name"],
                "Age": int(row["Age"]),
                "Email": row["Email"],
                "Score": float(row["Score"])
            }
            data.append(obj)

    # Output valid JSON to stdout
    print(json.dumps(data))

if __name__ == "__main__":
    main()