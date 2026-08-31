import csv
import json
from datetime import datetime, date

def parse_date(date_str):
    """Parse a date string in M/D/YYYY format to a date object."""
    if date_str and date_str != 'null':
        try:
            return datetime.strptime(date_str, '%m/%d/%Y').date()
        except ValueError:
            return None
    return None

def calculate_age(birth_date, reference_date):
    """Calculate age as of the reference date."""
    if birth_date is None:
        return 0
    # Age as of reference_date
    age = reference_date.year - birth_date.year
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def main():
    reference_date = date(2025, 7, 1)
    
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        result = []
        
        for row in reader:
            # Parse name into first and last name
            name = row['Name'].strip()
            parts = name.split()
            first_name = parts[0]
            # Handle cases where there might be a middle name
            if len(parts) == 2:
                last_name = parts[1]
            else:
                # Take everything after the first part as last name
                last_name = ' '.join(parts[1:])
            
            # Parse birthday
            birthday_str = row['Birthday'].strip()
            birth_date = parse_date(birthday_str)
            
            # Format birthday as YYYY-MM-DD
            if birth_date:
                birthday_formatted = birth_date.strftime('%Y-%m-%d')
            else:
                birthday_formatted = ''
            
            # Calculate age as of July 1, 2025
            age = calculate_age(birth_date, reference_date)
            
            # Collect relatives
            relatives = []
            
            # Father
            father_str = row['Father'].strip()
            if father_str and father_str != 'null':
                father_parts = father_str.split()
                relatives.append({
                    "FirstName": father_parts[0],
                    "LastName": ' '.join(father_parts[1:]),
                    "Relationship": "Father"
                })
            
            # Mother
            mother_str = row['Mother'].strip()
            if mother_str and mother_str != 'null':
                mother_parts = mother_str.split()
                relatives.append({
                    "FirstName": mother_parts[0],
                    "LastName": ' '.join(mother_parts[1:]),
                    "Relationship": "Mother"
                })
            
            # Brother
            brother_str = row['Brother'].strip()
            if brother_str and brother_str != 'null':
                brother_parts = brother_str.split()
                relatives.append({
                    "FirstName": brother_parts[0],
                    "LastName": ' '.join(brother_parts[1:]),
                    "Relationship": "Brother"
                })
            
            # Sister
            sister_str = row['Sister'].strip()
            if sister_str and sister_str != 'null':
                sister_parts = sister_str.split()
                relatives.append({
                    "FirstName": sister_parts[0],
                    "LastName": ' '.join(sister_parts[1:]),
                    "Relationship": "Sister"
                })
            
            # Create the entry
            entry = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday_formatted,
                "Age": age,
                "Relatives": relatives
            }
            
            result.append(entry)
    
    print(json.dumps(result, indent=2))

if __name__ == '__main__':
    main()