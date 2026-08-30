import csv
import json
from datetime import datetime

def calculate_age(birthday_str):
    return (datetime(2025, 7, 1) - datetime.strptime(birthday_str, "%d/%m/%Y")).days // 365

def parse_csv_to_json(csv_file):
    data = []
    with open(csv_file, mode='r') as file:
        csv_reader = csv.DictReader(file)
        for row in csv_reader:
            birth_date = row['Birthday']
            age = calculate_age(birth_date)
            data.append({
                "FirstName": row['Name'],
                "LastName": row['LastName'],
                "Birthday": row['Birthday'],
                "Age": age,
                "Relatives": []
            })
    return data

def add_relatives(person, relatives):
    for relative in relatives:
        relative['Person'] = person
        person['Relatives'].append(relative)

def main():
    csv_file = 'input/input.csv'
    json_file = 'output/output.json'

    # Parse CSV to JSON
    data = parse_csv_to_json(csv_file)
    
    # Assuming the CSV has a column 'Person' which links each relative back to a person
    for person in data:
        person_id = person['FirstName'] + ' ' + person['LastName']
        # Link each relative back to the correct person
        add_relatives(person, data)

    # Output the JSON data
    with open(json_file, 'w') as json_file:
        json.dump(data, json_file, indent=4)

if __name__ == "__main__":
    main()