import csv
import json
from datetime import datetime
from dateutil.relativedelta import relativedelta

def parse_date(date_str):
    if not date_str or date_str.lower() == 'null':
        return None
    # Handle various date formats: MM/DD/YYYY or YYYY-MM-DD
    try:
        return datetime.strptime(date_str, "%m/%d/%Y")
    except ValueError:
        try:
            return datetime.strptime(date_str, "%Y-%m-%d")
        except ValueError:
            return None

def calculate_age(birthday_str):
    birthday = parse_date(birthday_str)
    if not birthday:
        return 0
    
    reference_date = datetime(2025, 7, 1)
    age = reference_date.year - birthday.year
    if (reference_date.month, reference_date.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def process_person(row):
    # Parse Name: "First Last" or "First Middle Last" -> First Last
    name_parts = row['Name'].strip().split()
    first_name = name_parts[0]
    last_name = name_parts[-1] # Take the last part as surname
    
    birthday = row['Birthday']
    age = calculate_age(birthday)
    
    relatives = []
    
    father_name = row['Father']
    if father_name and father_name.lower() != 'null':
        parts = father_name.strip().split()
        relatives.append({
            "FirstName": parts[0],
            "LastName": parts[-1],
            "Relationship": "Father"
        })
        
    mother_name = row['Mother']
    if mother_name and mother_name.lower() != 'null':
        parts = mother_name.strip().split()
        relatives.append({
            "FirstName": parts[0],
            "LastName": parts[-1],
            "Relationship": "Mother"
        })
        
    brother_name = row['Brother']
    if brother_name and brother_name.lower() != 'null':
        parts = brother_name.strip().split()
        relatives.append({
            "FirstName": parts[0],
            "LastName": parts[-1],
            "Relationship": "Brother"
        })
        
    sister_name = row['Sister']
    if sister_name and sister_name.lower() != 'null':
        parts = sister_name.strip().split()
        relatives.append({
            "FirstName": parts[0],
            "LastName": parts[-1],
            "Relationship": "Sister"
        })
        
    return {
        "FirstName": first_name,
        "LastName": last_name,
        "Birthday": birthday,
        "Age": age,
        "Relatives": relatives
    }

def main():
    results = []
    
    with open('input/input.csv', 'r', newline='', encoding='utf-8') as csvfile:
        reader = csv.DictReader(csvfile)
        for row in reader:
            person = process_person(row)
            # Format birthday as ISO string
            person['Birthday'] = person['Birthday'].strftime('%Y-%m-%d')
            results.append(person)
            
    print(json.dumps(results, indent=2))

if __name__ == "__main__":
    main()