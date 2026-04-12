import csv
import json
from datetime import datetime, date

def parse_birthday(birthday_str):
    """Convert MM/D/YYYY to YYYY-MM-DD format"""
    parts = birthday_str.split('/')
    month, day, year = int(parts[0]), int(parts[1]), int(parts[2])
    return f"{year:04d}-{month:02d}-{day:02d}"

def calculate_age(birthday_str, reference_date):
    """Calculate age as of reference_date"""
    parts = birthday_str.split('/')
    birth_month = int(parts[0])
    birth_day = int(parts[1])
    birth_year = int(parts[2])
    
    age = reference_date.year - birth_year
    # If birthday hasn't occurred yet this year, subtract 1
    if (reference_date.month, reference_date.day) < (birth_month, birth_day):
        age -= 1
    return age

def split_name(full_name):
    """Split full name into first and last name"""
    parts = full_name.strip().split()
    if len(parts) >= 2:
        return parts[0], parts[-1]
    elif len(parts) == 1:
        return parts[0], parts[0]
    return "", ""

def process_relative(name_str, relationship):
    """Create a relative object if name is not null/empty"""
    if not name_str or name_str.lower() in ('null', 'none', ''):
        return None
    
    first_name, last_name = split_name(name_str)
    return {
        "FirstName": first_name,
        "LastName": last_name,
        "Relationship": relationship
    }

def main():
    reference_date = date(2025, 7, 1)
    result = []
    
    with open('input/input.csv', 'r', newline='', encoding='utf-8') as csvfile:
        reader = csv.DictReader(csvfile)
        
        for row in reader:
            # Parse name
            full_name = row['Name']
            first_name, last_name = split_name(full_name)
            
            # Parse birthday
            birthday = parse_birthday(row['Birthday'])
            
            # Calculate age
            age = calculate_age(row['Birthday'], reference_date)
            
            # Process relatives
            relatives = []
            
            # Father
            father_rel = process_relative(row.get('Father', ''), 'Father')
            if father_rel:
                relatives.append(father_rel)
            
            # Mother
            mother_rel = process_relative(row.get('Mother', ''), 'Mother')
            if mother_rel:
                relatives.append(mother_rel)
            
            # Brother
            brother_rel = process_relative(row.get('Brother', ''), 'Brother')
            if brother_rel:
                relatives.append(brother_rel)
            
            # Sister
            sister_rel = process_relative(row.get('Sister', ''), 'Sister')
            if sister_rel:
                relatives.append(sister_rel)
            
            # Create person record
            person = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday,
                "Age": age,
                "Relatives": relatives
            }
            result.append(person)
    
    # Output JSON array to stdout
    print(json.dumps(result, indent=2))

if __name__ == "__main__":
    main()