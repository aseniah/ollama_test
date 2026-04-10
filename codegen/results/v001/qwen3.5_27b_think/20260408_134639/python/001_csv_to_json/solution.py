import csv
import json

def convert_csv_to_json(csv_file='input/data.csv'):
    result = []
    
    with open(csv_file, 'r', newline='', encoding='utf-8') as f:
        reader = csv.DictReader(f)
        for row in reader:
            obj = {
                'Name': row['Name'],
                'Age': int(row['Age']),
                'Email': row['Email'],
                'Score': float(row['Score'])
            }
            result.append(obj)
    
    print(json.dumps(result))

if __name__ == '__main__':
    convert_csv_to_json()