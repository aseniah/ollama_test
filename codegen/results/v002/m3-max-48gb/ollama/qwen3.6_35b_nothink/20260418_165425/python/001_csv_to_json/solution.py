import csv
import json
import sys

def main():
    data = []
    with open('input/data.csv', 'r', newline='') as f:
        reader = csv.DictReader(f)
        for row in reader:
            record = {
                "Name": row["Name"],
                "Age": int(row["Age"]),
                "Email": row["Email"],
                "Score": float(row["Score"])
            }
            data.append(record)
    
    print(json.dumps(data, indent=2))

if __name__ == "__main__":
    main()