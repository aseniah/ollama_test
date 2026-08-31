import csv
import json

def main():
    with open('input/data.csv', 'r', newline='') as f:
        reader = csv.DictReader(f)
        rows = []
        for row in reader:
            rows.append({
                'Name': row['Name'],
                'Age': int(row['Age']),
                'Email': row['Email'],
                'Score': float(row['Score'])
            })
    print(json.dumps(rows))

if __name__ == '__main__':
    main()