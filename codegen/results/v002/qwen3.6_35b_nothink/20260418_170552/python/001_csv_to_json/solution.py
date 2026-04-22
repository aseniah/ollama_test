import csv
import json
import sys

def main():
    data = []
    try:
        with open('input/data.csv', 'r') as f:
            reader = csv.DictReader(f)
            for row in reader:
                obj = {
                    "Name": row['Name'],
                    "Age": int(row['Age']),
                    "Email": row['Email'],
                    "Score": float(row['Score'])
                }
                data.append(obj)
    except FileNotFoundError:
        sys.exit(1)
    
    print(json.dumps(data))

if __name__ == '__main__':
    main()