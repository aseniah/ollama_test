import csv
import json
from datetime import datetime, date

def read_input_csv(filename):
    with open(filename, 'r') as f:
        reader = csv.DictReader(f)
        rows = []
        for row in reader:
            rows.append(row)
    return rows

def calculate_age(birthday_str, reference_date):
    """Calculate age as of reference_date"""
    # Parse birthday in format M/D/YYYY
    parts = birthday_str.split('/')
    month = int(parts[0])
    day = int(parts[1])
    year = int(parts[2])
    birthday = date(year, month, day)
    
    # Calculate age
    age = reference_date.year - birthday.year
    # Check if birthday hasn't occurred yet in the reference year
    if (reference_date.month, reference_date.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def parse_name(full_name):
    """Parse full name into first and last name"""
    parts = full_name.split()
    first_name = parts[0]
    last_name = parts[-1]
    return first_name, last_name

def process_data(rows):
    reference_date = date(2025, 7, 1)
    result = []
    
    for row in rows:
        first_name, last_name = parse_name(row['Name'])
        
        # Calculate age
        if row['Died'] and row['Died'] != 'null':
            # If person died, age is their age at death
            died_parts = row['Died'].split('/')
            died_date = date(int(died_parts[2]), int(died_parts[0]), int(died_parts[1]))
            age = calculate_age(row['Birthday'], died_date)
        else:
            # Age as of July 1, 2025
            age = calculate_age(row['Birthday'], reference_date)
        
        # Format birthday as YYYY-MM-DD
        bday_parts = row['Birthday'].split('/')
        birthday_formatted = f"{int(bday_parts[2]):04d}-{int(bday_parts[0]):02d}-{int(bday_parts[1]):02d}"
        
        # Build relatives list
        relatives = []
        
        # Father
        if row['Father'] and row['Father'] != 'null':
            father_first, father_last = parse_name(row['Father'])
            relatives.append({
                "FirstName": father_first,
                "LastName": father_last,
                "Relationship": "Father"
            })
        
        # Mother
        if row['Mother'] and row['Mother'] != 'null':
            mother_first, mother_last = parse_name(row['Mother'])
            relatives.append({
                "FirstName": mother_first,
                "LastName": mother_last,
                "Relationship": "Mother"
            })
        
        # Brother
        if row['Brother'] and row['Brother'] != 'null':
            brother_first, brother_last = parse_name(row['Brother'])
            relatives.append({
                "FirstName": brother_first,
                "LastName": brother_last,
                "Relationship": "Brother"
            })
        
        # Sister
        if row['Sister'] and row['Sister'] != 'null':
            sister_first, sister_last = parse_name(row['Sister'])
            relatives.append({
                "FirstName": sister_first,
                "LastName": sister_last,
                "Relationship": "Sister"
            })
        
        result.append({
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": birthday_formatted,
            "Age": age,
            "Relatives": relatives
        })
    
    return result

if __name__ == '__main__':
    rows = read_input_csv('input/input.csv')
    output = process_data(rows)
    print(json.dumps(output, indent=2))