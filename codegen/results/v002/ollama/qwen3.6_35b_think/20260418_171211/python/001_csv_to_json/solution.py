import csv
import json

def main():
    records = []
    with open('input/data.csv', 'r', newline='') as f:
        reader = csv.DictReader(f)
        for row in reader:
            records.append({
                "Name": row["Name"],
                "Age": int(row["Age"]),
                "Email": row["Email"],
                "Score": float(row["Score"])
            })
    print(json.dumps(records))

if __name__ == "__main__":
    main()