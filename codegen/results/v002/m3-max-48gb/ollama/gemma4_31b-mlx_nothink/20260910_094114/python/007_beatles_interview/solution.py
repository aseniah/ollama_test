import csv
import json
from datetime import datetime

def calculate_age(birthday_str, target_date):
    # Birthday format in CSV is M/D/YYYY
    birth_date = datetime.strptime(birthday_str, "%m/%d/%Y")
    
    # Calculate age
    age = target_date.year - birth_date.year
    # Subtract one if the birthday hasn't occurred yet in the target year
    if (target_date.month, target_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def main():
    target_date = datetime(2025, 7, 1)
    results = []

    try:
        with open('input/input.csv', mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # Process Name into First and Last
                full_name = row['Name']
                name_parts = full_name.split(' ')
                # Based on expected_format.json: 
                # "John Winston Lennon" -> FirstName: "John", LastName: "Lennon"
                # "James Paul McCartney" -> FirstName: "James", LastName: "McCartney"
                first_name = name_parts[0]
                last_name = name_parts[-1]

                # Birthday format conversion: M/D/YYYY -> YYYY-MM-DD
                bday_raw = row['Birthday']
                bday_dt = datetime.strptime(bday_raw, "%m/%d/%Y")
                birthday_iso = bday_dt.strftime("%Y-%m-%d")

                # Calculate Age
                age = calculate_age(bday_raw, target_date)

                # Process Relatives
                relatives = []
                rel_cols = [
                    ('Father', 'Father'),
                    ('Mother', 'Mother'),
                    ('Brother', 'Brother'),
                    ('Sister', 'Sister')
                ]
                
                for col_name, relationship in rel_cols:
                    val = row.get(col_name)
                    if val and val.lower() != 'null':
                        rel_name_parts = val.split(' ')
                        relatives.append({
                            "FirstName": rel_name_parts[0],
                            "LastName": rel_name_parts[-1] if len(rel_name_parts) > 1 else "",
                            "Relationship": relationship
                        })

                results.append({
                    "FirstName": first_name,
                    "LastName": last_name,
                    "Birthday": birthday_iso,
                    "Age": age,
                    "Relatives": relatives
                })

        print(json.dumps(results, indent=2))

    except FileNotFoundError:
        pass

if __name__ == "__main__":
    main()