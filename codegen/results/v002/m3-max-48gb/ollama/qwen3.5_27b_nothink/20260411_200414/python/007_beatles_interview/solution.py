import csv
import json
from datetime import datetime

def calculate_age(birthday_str, reference_date):
    """Calculate age as of a specific date."""
    if not birthday_str:
        return None
    
    # Parse date from various formats (MM/DD/YYYY, M/D/YYYY, etc.)
    try:
        birth_date = datetime.strptime(birthday_str, "%m/%d/%Y")
    except ValueError:
        # Try alternative format if needed, though input seems consistent
        try:
            birth_date = datetime.strptime(birthday_str, "%m/%Y")
        except ValueError:
            return None
            
    # Calculate age
    age = reference_date.year - birth_date.year
    # Adjust if birthday hasn't occurred yet in the reference year
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def parse_name(full_name):
    """Extract first and last name from full name."""
    parts = full_name.strip().split()
    if len(parts) == 1:
        return parts[0], parts[0]
    elif len(parts) == 2:
        return parts[0], parts[1]
    else:
        # Assume first is first name, last is last name, middle names ignored
        return parts[0], parts[-1]

def parse_relative(name_str):
    """Parse relative name into first and last name components."""
    if not name_str or name_str.lower() == "null":
        return None
    
    parts = name_str.strip().split()
    if not parts:
        return None
    
    if len(parts) == 1:
        return {"FirstName": parts[0], "LastName": parts[0]}
    elif len(parts) == 2:
        return {"FirstName": parts[0], "LastName": parts[1]}
    else:
        return {"FirstName": parts[0], "LastName": parts[-1]}

def main():
    # Reference date: July 1, 2025
    reference_date = datetime(2025, 7, 1)
    
    result = []
    
    # Read the CSV file
    with open('input/input.csv', 'r') as csvfile:
        reader = csv.DictReader(csvfile)
        
        for row in reader:
            # Parse the main person's name
            first_name, last_name = parse_name(row['Name'])
            
            # Calculate age
            age = calculate_age(row['Birthday'], reference_date)
            
            # Build relatives list
            relatives = []
            
            # Check Father
            if row['Father'] and row['Father'].lower() != 'null':
                relative = parse_relative(row['Father'])
                if relative:
                    relative['Relationship'] = 'Father'
                    relatives.append(relative)
            
            # Check Mother
            if row['Mother'] and row['Mother'].lower() != 'null':
                relative = parse_relative(row['Mother'])
                if relative:
                    relative['Relationship'] = 'Mother'
                    relatives.append(relative)
            
            # Check Brother
            if row['Brother'] and row['Brother'].lower() != 'null':
                relative = parse_relative(row['Brother'])
                if relative:
                    relative['Relationship'] = 'Brother'
                    relatives.append(relative)
            
            # Check Sister
            if row['Sister'] and row['Sister'].lower() != 'null':
                relative = parse_relative(row['Sister'])
                if relative:
                    relative['Relationship'] = 'Sister'
                    relatives.append(relative)
            
            # Construct the person object
            person = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": row['Birthday'],
                "Age": age,
                "Relatives": relatives
            }
            
            result.append(person)
    
    # Output the JSON array
    print(json.dumps(result))

if __name__ == "__main__":
    main()