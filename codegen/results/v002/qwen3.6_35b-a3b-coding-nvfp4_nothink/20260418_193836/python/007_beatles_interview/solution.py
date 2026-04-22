import csv
import json
from datetime import date
from dateutil.relativedelta import relativedelta

def calculate_age(birthday_str, reference_date):
    """
    Calculate age as of reference_date given a birthday string in MM/DD/YYYY format.
    """
    # Parse the birthday
    parts = birthday_str.split('/')
    month = int(parts[0])
    day = int(parts[1])
    year = int(parts[2])
    birthday = date(year, month, day)
    
    # Calculate the difference
    delta = relativedelta(reference_date, birthday)
    return delta.years

def main():
    reference_date = date(2025, 7, 1)
    
    results = []
    
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        for row in reader:
            name_parts = row['Name'].split()
            first_name = name_parts[0]
            last_name = name_parts[-1]
            
            birthday_str = row['Birthday']
            age = calculate_age(birthday_str, reference_date)
            
            relatives = []
            
            # Father
            father_name = row['Father'].strip()
            if father_name and father_name != 'null':
                father_parts = father_name.split()
                father_first = father_parts[0]
                father_last = father_parts[-1]
                relatives.append({
                    "FirstName": father_first,
                    "LastName": father_last,
                    "Relationship": "Father"
                })
            
            # Mother
            mother_name = row['Mother'].strip()
            if mother_name and mother_name != 'null':
                mother_parts = mother_name.split()
                mother_first = mother_parts[0]
                mother_last = mother_parts[-1]
                relatives.append({
                    "FirstName": mother_first,
                    "LastName": mother_last,
                    "Relationship": "Mother"
                })
            
            # Brother
            brother_name = row['Brother'].strip()
            if brother_name and brother_name != 'null':
                brother_parts = brother_name.split()
                brother_first = brother_parts[0]
                brother_last = brother_parts[-1]
                relatives.append({
                    "FirstName": brother_first,
                    "LastName": brother_last,
                    "Relationship": "Brother"
                })
            
            # Sister
            sister_name = row['Sister'].strip()
            if sister_name and sister_name != 'null':
                sister_parts = sister_name.split()
                sister_first = sister_parts[0]
                sister_last = sister_parts[-1]
                relatives.append({
                    "FirstName": sister_first,
                    "LastName": sister_last,
                    "Relationship": "Sister"
                })
            
            entry = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday_str,
                "Age": age,
                "Relatives": relatives
            }
            results.append(entry)
    
    print(json.dumps(results, indent=2))

if __name__ == '__main__':
    main()