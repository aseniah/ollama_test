import csv
import json
from datetime import datetime

# Reference date for age calculation for living people
REFERENCE_DATE = datetime(2025, 7, 1)

def calculate_age(birth_date, death_date=None):
    """Calculate age based on death date if applicable, otherwise use reference date."""
    if death_date:
        target_date = death_date
    else:
        target_date = REFERENCE_DATE
    
    age = target_date.year - birth_date.year
    # If birthday hasn't occurred yet in the target year, subtract 1
    if (target_date.month, target_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def parse_date(date_str):
    """Parse date string in MM/DD/YYYY format, return None for 'null' values."""
    if date_str and str(date_str).lower() != 'null':
        return datetime.strptime(str(date_str), '%m/%d/%Y')
    return None

def parse_name(name_str):
    """Split name into first name (first word) and last name (last word)."""
    parts = name_str.strip().split()
    if len(parts) >= 2:
        return parts[0], parts[-1]
    return parts[0], ""

def process_relative(value, relationship):
    """Create relative object if value is not 'null', else return None."""
    if value and str(value).lower() != 'null':
        fname, lname = parse_name(str(value))
        return {
            "FirstName": fname,
            "LastName": lname,
            "Relationship": relationship
        }
    return None

# Read and process CSV file
with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    
    result = []
    for row in reader:
        # Parse main person's name and dates
        first_name, last_name = parse_name(row['Name'])
        birth_date = parse_date(row['Birthday'])
        death_date = parse_date(row['Died'])
        
        # Calculate age (using death date if applicable, else reference date)
        age = calculate_age(birth_date, death_date)
        
        # Build relatives list from parent/sibling fields
        relatives = []
        for field_name, rel_type in [
            ('Father', 'Father'),
            ('Mother', 'Mother'),
            ('Brother', 'Brother'),
            ('Sister', 'Sister')
        ]:
            relative = process_relative(row[field_name], rel_type)
            if relative is not None:
                relatives.append(relative)
        
        # Build output record
        record = {
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": birth_date.strftime('%Y-%m-%d'),
            "Age": age,
            "Relatives": relatives
        }
        
        result.append(record)

# Output JSON array to stdout (matching expected format style)
print(json.dumps(result, indent=2))