import csv
import json

# Read the CSV file and output JSON array to stdout
try:
    with open('input/data.csv', 'r') as f:
        reader = csv.DictReader(f)
        rows = []
        
        for row in reader:
            obj = {
                'Name': row['Name'],
                'Age': int(row['Age']),
                'Email': row['Email'],
                'Score': float(row['Score'])
            }
            rows.append(obj)
        
        json_output = json.dumps(rows, ensure_ascii=False)
        print(json_output)
except FileNotFoundError:
    print('Error: input/data.csv not found')
except Exception as e:
    print(f'Error: {str(e)}')