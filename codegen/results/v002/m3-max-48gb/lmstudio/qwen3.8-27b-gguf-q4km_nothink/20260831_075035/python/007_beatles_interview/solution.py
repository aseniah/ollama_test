import csv
import json
from datetime import date

# Reference date for age calculation
reference_date = date(2025, 7, 1)

def calculate_age(birthday_str):
    """Calculate age as of July 1, 2025"""
    # Parse birthday in M/D/YYYY format
    parts = birthday_str.split('/')
    month = int(parts[0])
    day = int(parts[1])
    year = int(parts[2])
    
    birthday = date(year, month, day)
    
    age = reference_date.year - birthday.year
    # Check if birthday hasn't occurred yet this year
    if (reference_date.month, reference_date.day) < (birthday.month, birthday.day):
        age -= 1
    
    return age

def parse_name(full_name):
    """Parse full name into first and last name"""
    parts = full_name.strip().split()
    if len(parts) == 1:
        return parts[0], ""
    # Assume last word is last name, everything before is first name
    last_name = parts[-1]
    first_name = ' '.join(parts[:-1])
    return first_name, last_name

def main():
    results = []
    
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        for row in reader:
            first_name, last_name = parse_name(row['Name'])
            birthday_str = row['Birthday']
            age = calculate_age(birthday_str)
            
            # Format birthday as YYYY-MM-DD
            parts = birthday_str.split('/')
            month = f"{int(parts[0]):02d}"
            day = f"{int(parts[1]):02d}"
            year = parts[2]
            birthday_formatted = f"{year}-{month}-{day}"
            
            # Build relatives list
            relatives = []
            if row['Father'] and row['Father'] != 'null':
                f_first, f_last = parse_name(row['Father'])
                relatives.append({
                    "FirstName": f_first,
                    "LastName": f_last,
                    "Relationship": "Father"
                })
            if row['Mother'] and row['Mother'] != 'null':
                m_first, m_last = parse_name(row['Mother'])
                relatives.append({
                    "FirstName": m_first,
                    "LastName": m_last,
                    "Relationship": "Mother"
                })
            if row['Brother'] and row['Brother'] != 'null':
                b_first, b_last = parse_name(row['Brother'])
                relatives.append({
                    "FirstName": b_first,
                    "LastName": b_last,
                    "Relationship": "Brother"
                })
            if row['Sister'] and row['Sister'] != 'null':
                s_first, s_last = parse_name(row['Sister'])
                relatives.append({
                    "FirstName": s_first,
                    "LastName": s_last,
                    "Relationship": "Sister"
                })
            
            results.append({
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday_formatted,
                "Age": age,
                "Relatives": relatives
            })
    
    print(json.dumps(results, indent=2))

if __name__ == '__main__':
    main()