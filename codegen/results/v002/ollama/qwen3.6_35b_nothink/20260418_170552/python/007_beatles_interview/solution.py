import csv
import json
from datetime import datetime

def calculate_age(birthday_str, reference_date):
    """Calculate age in years as of reference_date from birthday_str (MM/DD/YYYY)."""
    birth_date = datetime.strptime(birthday_str, "%m/%d/%Y")
    age = reference_date.year - birth_date.year
    # Check if the birthday hasn't occurred yet in the reference year
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def main():
    reference_date = datetime(2025, 7, 1)
    
    # Read the CSV file
    input_csv_path = 'input/input.csv'
    output_json = []
    
    with open(input_csv_path, mode='r', newline='') as csvfile:
        reader = csv.DictReader(csvfile)
        for row in reader:
            # Parse the Name field into First and Last Name
            name_parts = row['Name'].strip().split()
            if len(name_parts) >= 2:
                first_name = name_parts[0]
                last_name = ' '.join(name_parts[1:])
            elif len(name_parts) == 1:
                first_name = name_parts[0]
                last_name = ''
            else:
                first_name = ''
                last_name = ''
            
            birthday_str = row['Birthday'].strip()
            age = calculate_age(birthday_str, reference_date)
            
            relatives = []
            
            # Father
            father_name = row['Father'].strip()
            if father_name and father_name.lower() != 'null':
                parts = father_name.split()
                if parts:
                    relatives.append({
                        "FirstName": parts[0],
                        "LastName": ' '.join(parts[1:]) if len(parts) > 1 else '',
                        "Relationship": "Father"
                    })
            
            # Mother
            mother_name = row['Mother'].strip()
            if mother_name and mother_name.lower() != 'null':
                parts = mother_name.split()
                if parts:
                    relatives.append({
                        "FirstName": parts[0],
                        "LastName": ' '.join(parts[1:]) if len(parts) > 1 else '',
                        "Relationship": "Mother"
                    })
            
            # Brother
            brother_name = row['Brother'].strip()
            if brother_name and brother_name.lower() != 'null':
                parts = brother_name.split()
                if parts:
                    relatives.append({
                        "FirstName": parts[0],
                        "LastName": ' '.join(parts[1:]) if len(parts) > 1 else '',
                        "Relationship": "Brother"
                    })
            
            # Sister
            sister_name = row['Sister'].strip()
            if sister_name and sister_name.lower() != 'null':
                parts = sister_name.split()
                if parts:
                    relatives.append({
                        "FirstName": parts[0],
                        "LastName": ' '.join(parts[1:]) if len(parts) > 1 else '',
                        "Relationship": "Sister"
                    })
            
            person = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday_str,
                "Age": age,
                "Relatives": relatives
            }
            
            output_json.append(person)
    
    # Output the JSON array to stdout
    print(json.dumps(output_json, indent=2))

if __name__ == "__main__":
    main()