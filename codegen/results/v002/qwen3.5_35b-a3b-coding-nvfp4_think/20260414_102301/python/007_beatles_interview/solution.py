import csv
import json
from datetime import date, datetime

def parse_name(name):
    """Parse 'FirstName MiddleName LastName' to extract FirstName and LastName"""
    parts = name.strip().split()
    if len(parts) >= 2:
        return parts[0], parts[-1]
    return name.strip(), ""

def parse_birthday(birthday_str):
    """Parse 'M/D/YYYY' to 'YYYY-MM-DD' format"""
    parts = birthday_str.strip().split('/')
    if len(parts) == 3:
        month = int(parts[0])
        day = int(parts[1])
        year = int(parts[2])
        return f"{year}-{month:02d}-{day:02d}"
    return ""

def calculate_age(birthday_str, reference_date=date(2025, 7, 1)):
    """Calculate age as of reference_date"""
    parts = birthday_str.strip().split('/')
    if len(parts) == 3:
        month = int(parts[0])
        day = int(parts[1])
        year = int(parts[2])
        birthday = date(year, month, day)
        age = reference_date.year - birthday.year
        if (reference_date.month, reference_date.day) < (birthday.month, birthday.day):
            age -= 1
        return age
    return 0

def parse_relative(relative_str):
    """Parse 'FirstName LastName' to dict with FirstName, LastName, Relationship"""
    if not relative_str or relative_str.strip().lower() == 'null':
        return None
    parts = relative_str.strip().split()
    if len(parts) >= 2:
        return {
            "FirstName": parts[0],
            "LastName": parts[-1]
        }
    return None

def get_relationship_type(column_name):
    """Map CSV column to relationship type"""
    mapping = {
        "Father": "Father",
        "Mother": "Mother",
        "Brother": "Brother",
        "Sister": "Sister"
    }
    return mapping.get(column_name, "")

def main():
    reference_date = date(2025, 7, 1)
    result = []
    
    with open('input/input.csv', 'r', newline='', encoding='utf-8') as csvfile:
        reader = csv.DictReader(csvfile)
        for row in reader:
            first_name, last_name = parse_name(row['Name'])
            birthday_str = parse_birthday(row['Birthday'])
            age = calculate_age(row['Birthday'], reference_date)
            
            relatives = []
            for col in ['Father', 'Mother', 'Brother', 'Sister']:
                relative_str = row.get(col, '')
                if relative_str and relative_str.strip().lower() != 'null':
                    relative = parse_relative(relative_str)
                    if relative:
                        relative["Relationship"] = get_relationship_type(col)
                        relatives.append(relative)
            
            person = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday_str,
                "Age": age,
                "Relatives": relatives
            }
            result.append(person)
    
    print(json.dumps(result, indent=2))

if __name__ == "__main__":
    main()