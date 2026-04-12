import csv
import json
from datetime import datetime, date

def calculate_age(birth_date_str, target_date):
    """
    Calculates the age of a person as of the target_date.
    """
    # Parse the birth date (format: MM/DD/YYYY)
    birth_date = datetime.strptime(birth_date_str, '%m/%d/%Y').date()
    
    # Calculate age
    age = target_date.year - birth_date.year
    # Check if birthday has not occurred yet in the target year
    if (target_date.month, target_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def parse_name(full_name):
    """
    Splits a full name string into First Name and Last Name.
    Handles cases with middle names by taking the first word as First Name
    and the last word as Last Name (based on the expected output pattern).
    """
    parts = full_name.strip().split()
    first_name = parts[0]
    last_name = parts[-1]
    return first_name, last_name

def parse_relative_info(relative_name, relationship_type):
    """
    Parses a relative's name and returns a dictionary with FirstName, LastName, and Relationship.
    Returns None if the name is 'null' or empty.
    """
    if relative_name is None or relative_name.lower() == 'null' or relative_name.strip() == '':
        return None
    
    first, last = parse_name(relative_name)
    return {
        "FirstName": first,
        "LastName": last,
        "Relationship": relationship_type
    }

def main():
    target_date = date(2025, 7, 1)
    output_data = []

    # Read the CSV file
    with open('input/input.csv', 'r') as csvfile:
        reader = csv.DictReader(csvfile)
        
        for row in reader:
            # Parse Name
            first_name, last_name = parse_name(row['Name'])
            
            # Calculate Age
            age = calculate_age(row['Birthday'], target_date)
            
            # Parse Birthday to ISO format (YYYY-MM-DD)
            # Input format is MM/DD/YYYY
            birth_date_obj = datetime.strptime(row['Birthday'], '%m/%d/%Y')
            birthday_iso = birth_date_obj.strftime('%Y-%m-%d')
            
            # Build Relatives list
            relatives = []
            
            # Father
            if row.get('Father'):
                rel = parse_relative_info(row['Father'], 'Father')
                if rel:
                    relatives.append(rel)
            
            # Mother
            if row.get('Mother'):
                rel = parse_relative_info(row['Mother'], 'Mother')
                if rel:
                    relatives.append(rel)
            
            # Brother
            if row.get('Brother'):
                rel = parse_relative_info(row['Brother'], 'Brother')
                if rel:
                    relatives.append(rel)
            
            # Sister
            if row.get('Sister'):
                rel = parse_relative_info(row['Sister'], 'Sister')
                if rel:
                    relatives.append(rel)
            
            # Construct the record object
            record = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday_iso,
                "Age": age,
                "Relatives": relatives
            }
            
            output_data.append(record)

    # Output the result as JSON array to stdout
    print(json.dumps(output_data, indent=2))

if __name__ == "__main__":
    main()