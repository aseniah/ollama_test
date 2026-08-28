import csv
import json
from datetime import datetime

def calculate_age(birthday_str, reference_date):
    """Calculates age based on birthday and a reference date."""
    # Input format is M/D/YYYY
    birth_date = datetime.strptime(birthday_str, "%m/%d/%Y")
    age = reference_date.year - birth_date.year - ((reference_date.month, reference_date.day) < (birth_date.month, birth_date.day))
    return age

def parse_name(full_name):
    """Splits full name into First and Last. Assumes Last is the last word."""
    if not full_name or full_name.lower() == 'null':
        return None
    parts = full_name.split()
    if len(parts) == 0:
        return None
    if len(parts) == 1:
        return {"FirstName": parts[0], "LastName": ""}
    return {"FirstName": parts[0], "LastName": parts[-1]}

def main():
    reference_date = datetime(2025, 7, 1)
    output_data = []

    try:
        with open('input/input.csv', mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # Parse primary person
                full_name_parts = row['Name'].split()
                first_name = full_name_parts[0]
                last_name = full_name_parts[-1]
                
                birthday_raw = row['Birthday']
                # Convert M/D/YYYY to YYYY-MM-DD for output
                b_date = datetime.strptime(birthday_raw, "%m/%d/%Y")
                birthday_formatted = b_date.strftime("%Y-%m-%d")
                
                age = calculate_age(birthday_raw, reference_date)
                
                # Handle Relatives
                relatives = []
                relative_mapping = {
                    "Father": row['Father'],
                    "Mother": row['Mother'],
                    "Brother": row['Brother'],
                    "Sister": row['Sister']
                }
                
                for rel_type, rel_name in relative_mapping.items():
                    if rel_name and rel_name.lower() != 'null':
                        name_parts = rel_name.split()
                        relatives.append({
                            "FirstName": name_parts[0],
                            "LastName": name_parts[-1],
                            "Relationship": rel_type
                        })
                
                person = {
                    "FirstName": first_name,
                    "LastName": last_name,
                    "Birthday": birthday_formatted,
                    "Age": age,
                    "Relatives": relatives
                }
                output_data.append(person)
                
    except FileNotFoundError:
        pass

    print(json.dumps(output_data, indent=2))

if __name__ == "__main__":
    main()