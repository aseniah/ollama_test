import csv
import json
from datetime import datetime

REFERENCE_DATE = datetime(2025, 7, 1)

def calculate_age(birthday_str, reference_date):
    birth_date = datetime.strptime(birthday_str, "%m/%d/%Y")
    age = reference_date.year - birth_date.year
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def parse_name(full_name):
    parts = full_name.split()
    first_name = parts[0]
    last_name = parts[-1]
    return first_name, last_name

def parse_relative(rel_name):
    if rel_name is None or rel_name.strip() == "null" or rel_name.strip() == "":
        return None
    parts = rel_name.strip().split()
    first_name = parts[0]
    last_name = parts[-1]
    return {"FirstName": first_name, "LastName": last_name}

def format_date(birthday_str):
    birth_date = datetime.strptime(birthday_str, "%m/%d/%Y")
    return birth_date.strftime("%Y-%m-%d")

with open("input/input.csv", "r") as f:
    reader = csv.DictReader(f)
    result = []
    for row in reader:
        first_name, last_name = parse_name(row["Name"])
        birthday = format_date(row["Birthday"])
        age = calculate_age(row["Birthday"], REFERENCE_DATE)
        
        relatives = []
        for rel_type in ["Father", "Mother", "Brother", "Sister"]:
            rel_name = row[rel_type]
            parsed = parse_relative(rel_name)
            if parsed:
                relatives.append({
                    "FirstName": parsed["FirstName"],
                    "LastName": parsed["LastName"],
                    "Relationship": rel_type
                })
        
        result.append({
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": birthday,
            "Age": age,
            "Relatives": relatives
        })

print(json.dumps(result, indent=2))