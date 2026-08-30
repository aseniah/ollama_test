import csv
import json

def main():
    data = []
    with open('input/data.csv', mode='r', encoding='utf-8') as f:
        reader = csv.DictReader(f)
        for row in reader:
            data.append({
                'Name': row['Name'].strip(),
                'Age': int(row['Age'].strip()),
                'Email': row['Email'].strip(),
                'Score': float(row['Score'].strip())
            })
    print(json.dumps(data))

if __name__ == '__main__':
    main()