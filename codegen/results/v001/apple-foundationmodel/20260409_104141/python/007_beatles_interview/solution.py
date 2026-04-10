import csv
import json
from datetime import datetime

def calculate_age(birth_date):
    """Calculate age as of July 1, 2025."""
    today = datetime(2025, 7, 1)
    birth_date = datetime.strptime(birth_date, '%Y-%m-%d')
    return today.year - birth_date.year - ((today.month, today.day) < (birth_date.month, birth_date.day))

def process_csv_to_json(csv_file_path, json_file_path):
    with open(csv_file_path, mode='r') as csv_file, \
         open(json_file_path, mode='w') as json_file:
        csv_reader = csv.DictReader(csv_file)
        data = []

        for row in csv_reader:
            birth_date = row['birth_date']
            age = calculate_age(birth_date)
            entry = {
                'name': row['name'],
                'age': age
            }
            data.append(entry)

        json.dump(data, json_file, indent=4)

if __name__ == "__main__":
    csv_file_path = 'input/input.csv'
    json_file_path = 'output/output.json'
    process_csv_to_json(csv_file_path, json_file_path)