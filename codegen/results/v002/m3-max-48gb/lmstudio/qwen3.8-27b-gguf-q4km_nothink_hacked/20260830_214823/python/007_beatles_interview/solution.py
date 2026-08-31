import csv
import json
from datetime import date, datetime

REFERENCE_DATE = date(2025, 7, 1)

def parse_date(date_str):
    return datetime.strptime(date_str, "%m/%d/%Y").date()

def calculate_age(birthday_str, reference_date):
    birth = parse_date(birthday_str)
    age = reference_date.year - birth.year
    if (reference_date.month, reference_date.day) < (birth.month, birth.day):
        age -= 1
    return age

def is_null(val):
    return val is None or val.strip() == "" or val.strip().lower() == "null"

def main():
    relatives_map = {}
    records = []

    with open("input/input.csv", "r") as f:
        reader = csv.DictReader(f)
        for row in reader:
            name = row["Name"].strip()
            parts = name.split()
            first_name = parts[0]
            last_name = parts[-1] if len(parts) > 1 else ""
            birthday = row["Birthday"].strip()
            age = calculate_age(birthday, REFERENCE_DATE)
            relatives = []

            if not is_null(row.get("Father", "")):
                father_parts = row["Father"].strip().split()
                relatives.append({
                    "FirstName": father_parts[0],
                    "LastName": father_parts[-1] if len(father_parts) > 1 else "",
                    "Relationship": "Father"
                })
            if not is_null(row.get("Mother", "")):
                mother_parts = row["Mother"].strip().split()
                relatives.append({
                    "FirstName": mother_parts[0],
                    "LastName": mother_parts[-1] if len(mother_parts) > 1 else "",
                    "Relationship": "Mother"
                })
            if not is_null(row.get("Brother", "")):
                brother_parts = row["Brother"].strip().split()
                relatives.append({
                    "FirstName": brother_parts[0],
                    "LastName": brother_parts[-1] if len(brother_parts) > 1 else "",
                    "Relationship": "Brother"
                })
            if not is_null(row.get("Sister", "")):
                sister_parts = row["Sister"].strip().split()
                relatives.append({
                    "FirstName": sister_parts[0],
                    "LastName": sister_parts[-1] if len(sister_parts) > 1 else "",
                    "Relationship": "Sister"
                })

            records.append({
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday,
                "Age": age,
                "Relatives": relatives
            })

    print(json.dumps(records, indent=2))

if __name__ == "__main__":
    main()