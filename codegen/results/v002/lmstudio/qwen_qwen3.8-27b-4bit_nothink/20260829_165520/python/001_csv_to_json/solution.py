import csv
import json

def main():
    with open('input/data.csv', 'r') as f:
        reader = csv.DictReader(f)
        data = []
        for row in reader:
            obj = {
                "Name": row["Name"].strip(),
                "Age": int(row["Age"]),
                "Email": row["Email"].strip(),
                "Score": float(row["Score"])
            }
            data.append(obj)
    print(json.dumps(data))

if __name__ == "__main__":
    main()