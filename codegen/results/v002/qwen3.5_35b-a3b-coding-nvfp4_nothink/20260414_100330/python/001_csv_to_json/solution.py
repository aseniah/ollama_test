import csv
import json
import sys

def main():
    data = []
    try:
        with open('input/data.csv', mode='r', newline='', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                record = {
                    "Name": row["Name"],
                    "Age": int(row["Age"]),
                    "Email": row["Email"],
                    "Score": float(row["Score"])
                }
                data.append(record)
    except FileNotFoundError:
        pass
    except Exception as e:
        sys.stderr.write(f"Error reading file: {e}\n")
    
    print(json.dumps(data))

if __name__ == "__main__":
    main()