import csv
import json
from datetime import datetime
from io import StringIO

def calculate_age(birthday_str: str) -> int:
    """Calculate age as of July 1, 2025."""
    # Parse the birthday string in MM/DD/YYYY format
    birth_date = datetime.strptime(birthday_str, '%m/%d/%Y')
    # Reference date
    reference_date = datetime(2025, 7, 1)
    
    # Calculate age
    age = reference_date.year - birth_date.year
    # Check if birthday has occurred this year
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    
    return age

def parse_name(name_str: str) -> tuple:
    """Parse full name into FirstName and LastName.
    
    Handles names with middle names by taking the first and last non-empty parts.
    """
    parts = name_str.strip().split()
    if len(parts) >= 2:
        first_name = parts[0]
        last_name = parts[-1]
        return first_name, last_name
    elif len(parts) == 1:
        return parts[0], ""
    else:
        return "", ""

def get_relative(name_str: str, relationship: str) -> dict:
    """Create a relative dict from name and relationship."""
    if name_str.lower() == 'null' or not name_str.strip():
        return None
    
    first_name, last_name = parse_name(name_str)
    return {
        "FirstName": first_name,
        "LastName": last_name,
        "Relationship": relationship
    }

def main():
    # Read the CSV file
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        rows = list(reader)
    
    result = []
    
    for row in rows:
        # Parse the name
        first_name, last_name = parse_name(row['Name'])
        
        # Calculate age
        age = calculate_age(row['Birthday'])
        
        # Build relatives list
        relatives = []
        
        # Father
        father_rel = get_relative(row['Father'], 'Father')
        if father_rel:
            relatives.append(father_rel)
        
        # Mother
        mother_rel = get_relative(row['Mother'], 'Mother')
        if mother_rel:
            relatives.append(mother_rel)
        
        # Brother
        brother_rel = get_relative(row['Brother'], 'Brother')
        if brother_rel:
            relatives.append(brother_rel)
        
        # Sister
        sister_rel = get_relative(row['Sister'], 'Sister')
        if sister_rel:
            relatives.append(sister_rel)
        
        person = {
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": row['Birthday'],  # Keep original format
            "Age": age,
            "Relatives": relatives
        }
        
        result.append(person)
    
    # Output JSON array
    print(json.dumps(result, indent=2))

if __name__ == "__main__":
    main()