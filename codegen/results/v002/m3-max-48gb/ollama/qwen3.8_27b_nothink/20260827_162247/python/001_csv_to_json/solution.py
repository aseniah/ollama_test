import csv
import json

def main():
    data = []
    with open('input/data.csv', 'r') as f:
        reader = csv.DictReader(f)
        for row in reader:
            data.append({
                "Name": row['Name'],
                "Age": int(row['Age']),
                "Email": row['Email'],
                "Score": float(row['Score'])
            })
    print(json.dumps(data))

if __name__ == '__main__':
    main()