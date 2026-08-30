import csv
import json
from datetime import datetime, date

def parse_name(name):
    """Split full name into first and last name."""
    parts = name.strip().split()
    first_name = parts[0] if parts else ""
    last_name = parts[-1] if len(parts) > 1 else ""
    return first_name, last_name

def parse_date(date_str):
    """Parse date string in MM/D/YYYY format and return YYYY-MM-DD."""
    if not date_str or date_str == 'null':
        return None
    try:
        dt = datetime.strptime(date_str, '%m/%d/%Y')
        return dt.strftime('%Y-%m-%d')
    except ValueError:
        try:
            dt = datetime.strptime(date_str, '%m/%Y')
            return dt.strftime('%Y-%m-%d')
        except ValueError:
            return date_str

def calculate_age(birthday_str, reference_date):
    """Calculate age as of reference_date."""
    if not birthday_str:
        return 0
    try:
        birth_date = datetime.strptime(birthday_str, '%Y-%m-%d').date()
        age = reference_date.year - birth_date.year
        # Check if birthday hasn't occurred yet this year
        if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
            age -= 1
        return age
    except ValueError:
        return 0

def main():
    reference_date = date(2025, 7, 1)
    result = []
    
    with open('input/input.csv', 'r', newline='', encoding='utf-8') as f:
        reader = csv.DictReader(f)
        for row in reader:
            first_name, last_name = parse_name(row['Name'])
            birthday = parse_date(row['Birthday'])
            age = calculate_age(birthday, reference_date)
            
            relatives = []
            
            # Father
            if row['Father'] and row['Father'] != 'null':
                fn, ln = parse_name(row['Father'])
                relatives.append({
                    "FirstName": fn,
                    "LastName": ln,
                    "Relationship": "Father"
                })
            
            # Mother
            if row['Mother'] and row['Mother'] != 'null':
                fn, ln = parse_name(row['Mother'])
                relatives.append({
                    "FirstName": fn,
                    "LastName": ln,
                    "Relationship": "Mother"
                })
            
            # Brother
            if row['Brother'] and row['Brother'] != 'null':
                fn, ln = parse_name(row['Brother'])
                relatives.append({
                    "FirstName": fn,
                    "LastName": ln,
                    "Relationship": "Brother"
                })
            
            # Sister
            if row['Sister'] and row['Sister'] != 'null':
                fn, ln = parse_name(row['Sister'])
                relatives.append({
                    "FirstName": fn,
                    "LastName": ln,
                    "Relationship": "Sister"
                })
            
            person = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday,
                "Age": age,
                "Relatives": relatives
            }
            result.append(person)
    
    print(json.dumps(result, indent=2))

if __name__ == "__main__":
    main()