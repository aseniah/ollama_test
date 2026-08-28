import csv
import json
from datetime import datetime
from datetime import date

def parse_birthday(birthday_str):
    """Parse birthday string in M/D/YYYY format to date object"""
    if not birthday_str or birthday_str == 'null':
        return None
    date_obj = datetime.strptime(birthday_str, '%m/%d/%Y').date()
    return date_obj

def calculate_age(birthday, reference_date):
    """Calculate age as of reference_date"""
    if not birthday:
        return 0
    age = reference_date.year - birthday.year
    if (reference_date.month, reference_date.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def parse_name(name_str):
    """Parse name into first and last name"""
    parts = name_str.strip().split()
    if len(parts) >= 2:
        return parts[0], ' '.join(parts[1:])
    return parts[0] if parts else '', ''

def format_date(date_obj):
    """Format date as YYYY-MM-DD"""
    return date_obj.strftime('%Y-%m-%d')

def build_relatives(row):
    """Build relatives list from row"""
    relatives = []
    
    father = row.get('Father', '').strip()
    if father and father != 'null':
        parts = father.strip().split()
        if len(parts) >= 2:
            relatives.append({
                "FirstName": parts[0],
                "LastName": parts[1],
                "Relationship": "Father"
            })
    
    mother = row.get('Mother', '').strip()
    if mother and mother != 'null':
        parts = mother.strip().split()
        if len(parts) >= 2:
            relatives.append({
                "FirstName": parts[0],
                "LastName": parts[1],
                "Relationship": "Mother"
            })
    
    brother = row.get('Brother', '').strip()
    if brother and brother != 'null':
        parts = brother.strip().split()
        if len(parts) >= 2:
            relatives.append({
                "FirstName": parts[0],
                "LastName": parts[1],
                "Relationship": "Brother"
            })
    
    sister = row.get('Sister', '').strip()
    if sister and sister != 'null':
        parts = sister.strip().split()
        if len(parts) >= 2:
            relatives.append({
                "FirstName": parts[0],
                "LastName": parts[1],
                "Relationship": "Sister"
            })
    
    return relatives

reference_date = date(2025, 7, 1)

with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    result = []
    
    for row in reader:
        name = row.get('Name', '').strip()
        first_name, last_name = parse_name(name)
        
        birthday_str = row.get('Birthday', '').strip()
        birthday = parse_birthday(birthday_str)
        
        relatives = build_relatives(row)
        
        entry = {
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": format_date(birthday) if birthday else None,
            "Age": calculate_age(birthday, reference_date) if birthday else 0,
            "Relatives": relatives
        }
        result.append(entry)

print(json.dumps(result))