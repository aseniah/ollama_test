import csv
from datetime import datetime

def calculate_age(birthday):
    birth_date = datetime.strptime(birthday, "%m/%d/%Y")
    current_year = 2025
    age = current_year - birth_date.year
    if (current_year, birth_date.month, birth_date.day) < (birth_date.year, birth_date.month, birth_date.day):
        age -= 1
    return age

def read_csv_file(file_path):
    with open(file_path, mode='r', newline='') as file:
        reader = csv.DictReader(file)
        data = list(reader)
    return data

def transform_data(data):
    result = []
    for row in data:
        relatives = []
        if row['Father']:
            relatives.append({'FirstName': row['Father'], 'LastName': '', 'Relationship': 'Father'})
        if row['Mother']:
            relatives.append({'FirstName': row['Mother'], 'LastName': '', 'Relationship': 'Mother'})
        if row['Brother']:
            relatives.append({'FirstName': row['Brother'], 'LastName': '', 'Relationship': 'Brother'})
        if row['Sister']:
            relatives.append({'FirstName': row['Sister'], 'LastName': '', 'Relationship': 'Sister'})
        
        transformed_row = {
            "FirstName": row['Name'].split()[0],
            "LastName": row['Name'].split()[1] if len(row['Name'].split()) > 1 else '',
            "Birthday": row['Birthday'],
            "Age": calculate_age(row['Birthday']),
            "Relatives": relatives
        }
        result.append(transformed_row)
    return result

def write_json_file(data):
    import json
    print(json.dumps(data, indent=2))

if __name__ == "__main__":
    csv_data = read_csv_file('input/input.csv')
    transformed_data = transform_data(csv_data)
    write_json_file(transformed_data)