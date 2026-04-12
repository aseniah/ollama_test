import csv
import json
from datetime import date

def parse_date(date_str):
    """Convert MM/DD/YYYY string to a date object."""
    if not date_str or date_str.lower() == 'null':
        return None
    parts = date_str.split('/')
    if len(parts) != 3:
        return None
    month, day, year = int(parts[0]), int(parts[1]), int(parts[2])
    try:
        return date(year, month, day)
    except ValueError:
        return None

def calculate_age(birth_date, reference_date):
    """Calculate age as of the reference date."""
    if birth_date is None:
        return 0
    
    age = reference_date.year - birth_date.year
    
    # Check if birthday has already occurred in the reference year
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
        
    return age

def main():
    input_file = 'input/input.csv'
    reference_date = date(2025, 7, 1)
    
    records = []
    
    with open(input_file, mode='r', newline='', encoding='utf-8') as file:
        reader = csv.DictReader(file)
        
        for row in reader:
            # Extract name parts (split on first space)
            full_name = row.get('Name', '')
            name_parts = full_name.split(None, 1)  # Split into max 2 parts
            
            first_name = name_parts[0] if len(name_parts) > 0 else ""
            last_name = name_parts[1] if len(name_parts) > 1 else ""
            
            birth_date = parse_date(row.get('Birthday', ''))
            age = calculate_age(birth_date, reference_date)
            
            # Build relatives list
            relatives = []
            
            # Helper function to add relative if name exists and is not null/empty
            def add_relative(first, last, relation_type):
                full_rel_name = (first + " " + last).strip()
                if full_rel_name and full_rel_name.lower() != 'null':
                    # Split relative name similarly
                    rel_parts = full_rel_name.split(None, 1)
                    rel_first = rel_parts[0] if len(rel_parts) > 0 else ""
                    rel_last = rel_parts[1] if len(rel_parts) > 1 else ""
                    
                    if rel_first: # Only add if we have a first name
                        relatives.append({
                            "FirstName": rel_first,
                            "LastName": rel_last,
                            "Relationship": relation_type
                        })

            # Process family members
            father_name = row.get('Father', '')
            mother_name = row.get('Mother', '')
            brother_name = row.get('Brother', '')
            sister_name = row.get('Sister', '')
            
            add_relative(father_name, "", "Father")
            add_relative(mother_name, "", "Mother")
            add_relative(brother_name, "", "Brother")
            add_relative(sister_name, "", "Sister")
            
            # Format birthday for output (YYYY-MM-DD)
            birthday_str = ""
            if birth_date:
                birthday_str = birth_date.strftime("%Y-%m-%d")
            
            record = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday_str,
                "Age": age,
                "Relatives": relatives
            }
            
            records.append(record)
    
    # Output JSON array to stdout
    print(json.dumps(records))

if __name__ == "__main__":
    main()