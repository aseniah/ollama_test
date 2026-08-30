import csv
import json
from datetime import datetime

def calculate_age(birthday_str, reference_date=datetime(2025, 7, 1)):
    """Calculate age as of the reference date."""
    birthday = datetime.strptime(birthday_str, "%m/%d/%Y")
    age = reference_date.year - birthday.year
    # Adjust if birthday hasn't occurred yet in the reference year
    if reference_date < birthday.replace(year=reference_date.year):
        age -= 1
    return age

def process_input():
    try:
        with open('input/input.csv', 'r') as f:
            reader = csv.DictReader(f)
            result = []
            
            for row in reader:
                name_parts = row['Name'].split()
                first_name = name_parts[0]
                last_name = name_parts[1]
                birthday = row['Birthday']
                
                # Parse relatives
                relatives = []
                if row['Father'] != 'null':
                    relatives.append({
                        "FirstName": row['Father'].split()[0],
                        "LastName": row['Father'].split()[1] if len(row['Father'].split()) > 1 else row['Father'].split()[0],
                        "Relationship": "Father"
                    })
                
                if row['Mother'] != 'null':
                    relatives.append({
                        "FirstName": row['Mother'].split()[0],
                        "LastName": row['Mother'].split()[1] if len(row['Mother'].split()) > 1 else row['Mother'].split()[0],
                        "Relationship": "Mother"
                    })
                
                if row['Brother'] != 'null':
                    relatives.append({
                        "FirstName": row['Brother'].split()[0],
                        "LastName": row['Brother'].split()[1] if len(row['Brother'].split()) > 1 else row['Brother'].split()[0],
                        "Relationship": "Brother"
                    })
                
                if row['Sister'] != 'null':
                    relatives.append({
                        "FirstName": row['Sister'].split()[0],
                        "LastName": row['Sister'].split()[1] if len(row['Sister'].split()) > 1 else row['Sister'].split()[0],
                        "Relationship": "Sister"
                    })
                
                result.append({
                    "FirstName": first_name,
                    "LastName": last_name,
                    "Birthday": birthday.replace('/', '-'),
                    "Age": calculate_age(birthday),
                    "Relatives": relatives
                })
            
            return result
    
    except FileNotFoundError:
        return None

if __name__ == "__main__":
    data = process_input()
    if data:
        print(json.dumps(data, indent=2))
    else:
        print(json.dumps([]))