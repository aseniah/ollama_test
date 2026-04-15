import csv
import json
import sys

def main():
    try:
        with open('input/data.csv', mode='r', newline='', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            data = []
            for row in reader:
                record = {
                    "Name": row['Name'],
                    "Age": int(row['Age']),
                    "Email": row['Email'],
                    "Score": float(row['Score'])
                }
                data.append(record)
            print(json.dumps(data))
    except FileNotFoundError:
        # If file doesn't exist, output empty array as per robust handling, 
        # though the prompt implies the file exists.
        print("[]")
    except Exception as e:
        # In case of other parsing errors, output empty array
        print("[]")

if __name__ == "__main__":
    main()