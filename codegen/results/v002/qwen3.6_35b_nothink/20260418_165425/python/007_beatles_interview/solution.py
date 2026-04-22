import csv
import json
from datetime import datetime

def calculate_age(birthday_str, reference_date_str):
    """
    Calculate age as of the reference date.
    """
    # Parse the birthday from the CSV format (MM/DD/YYYY)
    birthday = datetime.strptime(birthday_str, "%m/%d/%Y")
    # Reference date: July 1, 2025
    reference_date = datetime.strptime(reference_date_str, "%m/%d/%Y")
    
    age = reference_date.year - birthday.year
    # Adjust if the birthday hasn't occurred yet in the reference year
    if (reference_date.month, reference_date.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def main():
    reference_date_str = "07/01/2025"
    result = []
    
    with open('input/input.csv', 'r') as csvfile:
        reader = csv.DictReader(csvfile)
        for row in reader:
            # Parse the full name into first and last name
            name_parts = row['Name'].strip().split()
            first_name = name_parts[0]
            last_name = name_parts[-1]
            
            # Parse birthday and calculate age
            birthday = row['Birthday'].strip()
            # Format birthday as YYYY-MM-DD for output
            birthday_dt = datetime.strptime(birthday, "%m/%d/%Y")
            birthday_formatted = birthday_dt.strftime("%Y-%m-%d")
            age = calculate_age(birthday, reference_date_str)
            
            # Build the relatives list
            relatives = []
            father = row['Father'].strip()
            mother = row['Mother'].strip()
            brother = row['Brother'].strip()
            sister = row['Sister'].strip()
            
            if father != 'null' and father != '':
                f_parts = father.split()
                relatives.append({
                    "FirstName": f_parts[0],
                    "LastName": f_parts[-1],
                    "Relationship": "Father"
                })
            
            if mother != 'null' and mother != '':
                m_parts = mother.split()
                relatives.append({
                    "FirstName": m_parts[0],
                    "LastName": m_parts[-1],
                    "Relationship": "Mother"
                })
            
            if brother != 'null' and brother != '':
                b_parts = brother.split()
                relatives.append({
                    "FirstName": b_parts[0],
                    "LastName": b_parts[-1],
                    "Relationship": "Brother"
                })
            
            if sister != 'null' and sister != '':
                s_parts = sister.split()
                relatives.append({
                    "FirstName": s_parts[0],
                    "LastName": s_parts[-1],
                    "Relationship": "Sister"
                })
            
            person = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday_formatted,
                "Age": age,
                "Relatives": relatives
            }
            result.append(person)
    
    # Output the JSON array
    print(json.dumps(result, indent=2))

if __name__ == "__main__":
    main()