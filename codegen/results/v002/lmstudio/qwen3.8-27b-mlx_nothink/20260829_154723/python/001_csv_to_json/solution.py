import csv
import json

def main():
    with open('input/data.csv', newline='') as f:
        reader = csv.DictReader(f)
        result = []
        for row in reader:
            result.append({
                "Name": row["Name"],
                "Age": int(row["Age"]),
                "Email": row["Email"],
                "Score": float(row["Score"])
            })
    print(json.dumps(result))

if __name__ == "__main__":
    main()