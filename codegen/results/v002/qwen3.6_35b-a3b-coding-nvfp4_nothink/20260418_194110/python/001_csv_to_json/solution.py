import csv
import json
import sys

def main():
    input_file = 'input/data.csv'
    output = []
    
    with open(input_file, 'r', newline='') as f:
        reader = csv.DictReader(f)
        for row in reader:
            obj = {
                "Name": row['Name'],
                "Age": int(row['Age']),
                "Email": row['Email'],
                "Score": float(row['Score'])
            }
            output.append(obj)
    
    print(json.dumps(output))

if __name__ == '__main__':
    main()