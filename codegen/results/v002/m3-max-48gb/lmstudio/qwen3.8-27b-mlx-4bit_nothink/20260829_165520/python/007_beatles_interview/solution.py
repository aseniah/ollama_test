import csv
import json
from datetime import date

def parse_date(mdy_str):
    """Parse MM/DD/YYYY to date object"""
    parts = mdy_str.strip().split('/')
    month = int(parts[0])
    day = int(parts[1])
    year = int(parts[2])
    return date(year, month, day)

def split_name(full_name):
    """Split a full name into (FirstName, LastName)"""
    full_name = full_name.strip()
    parts = full_name.split()
    if len(parts) == 1:
        return parts[0], ""
    first_name = parts[0]
    last_name = ' '.join(parts[1:])
    return first_name, last_name

def calculate_age(birth_date, died_date, reference_date):
    """Calculate age as of reference_date."""
    if died_date is not None:
        return died_date.year - birth_date.year
    else:
        # Person is alive
        age = reference_date.year - birth_date.year
        # If birthday hasn't occurred yet this year, subtract 1
        if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
            age -= 1
        return age

def main():
    reference_date = date(2025, 7, 1)
    
    relatives_keys = ['Father', 'Mother', 'Brother', 'Sister']
    
    results = []
    
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        for row in reader:
            # Parse name
            first_name, last_name = split_name(row['Name'])
            
            # Parse birthday
            birthday = parse_date(row['Birthday'])
            birthday_str = birthday.strftime('%Y-%m-%d')
            
            # Parse death date
            died_str = row['Died'].strip()
            if died_str == 'null' or died_str == '':
                died_date = None
            else:
                died_date = parse_date(died_str)
            
            # Calculate age
            age = calculate_age(birthday, died_date, reference_date)
            
            # Build relatives
            relatives = []
            for rel_key in relatives_keys:
                rel_name = row[rel_key].strip()
                if rel_name != 'null' and rel_name != '':
                    rel_first, rel_last = split_name(rel_name)
                    relatives.append({
                        "FirstName": rel_first,
                        "LastName": rel_last,
                        "Relationship": rel_key
                    })
            
            person = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday_str,
                "Age": age,
                "Relatives": relatives
            }
            results.append(person)
    
    print(json.dumps(results, indent=2))

if __name__ == '__main__':
    main()