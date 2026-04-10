import csv
import json
from datetime import date

def calculate_age(birth_date_str):
    birth_date = date.fromisoformat(birth_date_str)
    reference_date = date(2025, 7, 1)
    age = reference_date.year - birth_date.year
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def transform_csv_to_json():
    result = []
    
    with open('input/input.csv', 'r') as csvfile:
        reader = csv.DictReader(csvfile)
        for row in reader:
            person = {
                "id": int(row["id"]),
                "name": row["name"],
                "email": row["email"],
                "birthDate": row["birthDate"],
                "age": calculate_age(row["birthDate"])
            }
            result.append(person)
    
    print(json.dumps(result, indent=2))

transform_csv_to_json()