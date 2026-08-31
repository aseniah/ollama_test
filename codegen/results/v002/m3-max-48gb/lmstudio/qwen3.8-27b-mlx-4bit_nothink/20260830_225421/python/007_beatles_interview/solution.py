import csv
import json
from datetime import date

def parse_date(date_str):
    """Parse a date string in M/D/YYYY format"""
    if date_str.strip() == '' or date_str.strip() == 'null':
        return None
    parts = date_str.strip().split('/')
    if len(parts) != 3:
        return None
    month, day, year = int(parts[0]), int(parts[1]), int(parts[2])
    return date(year, month, day)

def calculate_age(birthday, died, reference_date):
    """Calculate age as of reference_date (July 1, 2025)"""
    if birthday is None:
        return None
    
    if died is not None and died < reference_date:
        # Person died before reference date, age is at time of death
        if died.month < birthday.month or (died.month == birthday.month and died.day < birthday.day):
            age = died.year - birthday.year - 1
        else:
            age = died.year - birthday.year
        return age
    else:
        # Person is alive (or died on/after reference date)
        if birthday.month > reference_date.month or (birthday.month == reference_date.month and birthday.day > reference_date.day):
            age = reference_date.year - birthday.year - 1
        else:
            age = reference_date.year - birthday.year
        return age

def extract_names(full_name):
    """Extract first and last name from full name"""
    if not full_name or full_name.strip() == '' or full_name.strip() == 'null':
        return None, None
    parts = full_name.strip().split()
    if len(parts) == 1:
        return parts[0], ''
    elif len(parts) == 2:
        return parts[0], parts[1]
    else:
        # If more than 2 parts, treat first as first name, rest as last name
        return parts[0], ' '.join(parts[1:])

reference_date = date(2025, 7, 1)

results = []

with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    for row in reader:
        name = row['Name'].strip()
        first_name, last_name = extract_names(name)
        
        birthday = parse_date(row['Birthday'])
        died = parse_date(row['Died'])
        
        age = calculate_age(birthday, died, reference_date)
        
        relatives = []
        
        # Father
        father = row.get('Father', 'null').strip()
        if father and father != 'null':
            f_first, f_last = extract_names(father)
            relatives.append({
                "FirstName": f_first,
                "LastName": f_last,
                "Relationship": "Father"
            })
        
        # Mother
        mother = row.get('Mother', 'null').strip()
        if mother and mother != 'null':
            m_first, m_last = extract_names(mother)
            relatives.append({
                "FirstName": m_first,
                "LastName": m_last,
                "Relationship": "Mother"
            })
        
        # Brother
        brother = row.get('Brother', 'null').strip()
        if brother and brother != 'null':
            b_first, b_last = extract_names(brother)
            relatives.append({
                "FirstName": b_first,
                "LastName": b_last,
                "Relationship": "Brother"
            })
        
        # Sister
        sister = row.get('Sister', 'null').strip()
        if sister and sister != 'null':
            s_first, s_last = extract_names(sister)
            relatives.append({
                "FirstName": s_first,
                "LastName": s_last,
                "Relationship": "Sister"
            })
        
        entry = {
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": row['Birthday'].strip(),
            "Age": age,
            "Relatives": relatives
        }
        results.append(entry)

print(json.dumps(results, indent=2))