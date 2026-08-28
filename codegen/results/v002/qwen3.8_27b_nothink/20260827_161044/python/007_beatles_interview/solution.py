import csv
import json
from datetime import datetime

def calculate_age(birthday_str, reference_date):
    birthday = datetime.strptime(birthday_str, "%m/%d/%Y")
    if birthday > reference_date:
        return -1
    age = reference_date.year - birthday.year
    if (reference_date.month, reference_date.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def main():
    reference_date = datetime(2025, 7, 1)
    
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        records = list(reader)
    
    result = []
    for record in records:
        # Parse name
        name_parts = record['Name'].split()
        first_name = name_parts[0]
        last_name = name_parts[-1]
        
        # Parse birthday
        birthday_str = record['Birthday']
        
        # Calculate age (ignore if died, as per expected output which still has ages)
        age = calculate_age(birthday_str, reference_date)
        
        # Build relatives list
        relatives = []
        
        father = record.get('Father', 'null')
        if father != 'null' and father != '':
            father_parts = father.split()
            relatives.append({
                "FirstName": father_parts[0],
                "LastName": father_parts[-1],
                "Relationship": "Father"
            })
        
        mother = record.get('Mother', 'null')
        if mother != 'null' and mother != '':
            mother_parts = mother.split()
            relatives.append({
                "FirstName": mother_parts[0],
                "LastName": mother_parts[-1],
                "Relationship": "Mother"
            })
        
        brother = record.get('Brother', 'null')
        if brother != 'null' and brother != '':
            brother_parts = brother.split()
            relatives.append({
                "FirstName": brother_parts[0],
                "LastName": brother_parts[-1],
                "Relationship": "Brother"
            })
        
        sister = record.get('Sister', 'null')
        if sister != 'null' and sister != '':
            sister_parts = sister.split()
            relatives.append({
                "FirstName": sister_parts[0],
                "LastName": sister_parts[-1],
                "Relationship": "Sister"
            })
        
        entry = {
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": datetime.strptime(birthday_str, "%m/%d/%Y").strftime("%Y-%m-%d"),
            "Age": age,
            "Relatives": relatives
        }
        result.append(entry)
    
    print(json.dumps(result, indent=2))

if __name__ == '__main__':
    main()