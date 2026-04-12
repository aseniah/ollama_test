import csv
import json
from datetime import datetime, date

def calculate_age(birthday_str, reference_date):
    """Calculate age as of the reference date."""
    try:
        # Parse M/D/YYYY format
        birth_date = datetime.strptime(birthday_str, "%m/%d/%Y").date()
    except ValueError:
        return 0

    if birth_date > reference_date:
        return 0

    age = reference_date.year - birth_date.year
    # Check if birthday has occurred in the current year of reference date
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def process_csv(input_file, output_format_file=None):
    # Reference date: July 1, 2025
    reference_date = date(2025, 7, 1)
    
    result = []
    
    with open(input_file, 'r', newline='', encoding='utf-8') as csvfile:
        reader = csv.DictReader(csvfile)
        
        for row in reader:
            # Parse Name: "John Winston Lennon" -> First: "John", Last: "Lennon"
            full_name = row['Name'].strip()
            name_parts = full_name.split()
            
            if len(name_parts) >= 2:
                first_name = name_parts[0]
                last_name = ' '.join(name_parts[1:])
            elif len(name_parts) == 1:
                first_name = name_parts[0]
                last_name = ""
            else:
                continue # Skip malformed rows

            # Parse Birthday: "10/9/1940" -> ISO format "1940-09-09"
            try:
                birth_date_obj = datetime.strptime(row['Birthday'].strip(), "%m/%d/%Y").date()
                iso_birthday = birth_date_obj.isoformat()
            except ValueError:
                iso_birthday = None

            # Calculate Age
            age = calculate_age(row['Birthday'], reference_date) if row['Birthday'] else 0

            # Build Relatives list
            relatives = []
            
            # Helper to add relative
            def add_relative(name, relationship):
                name = name.strip()
                if name and name.lower() != "null":
                    parts = name.split(maxsplit=1)
                    if len(parts) >= 2:
                        fname = parts[0]
                        lname = parts[1]
                    elif len(parts) == 1:
                        fname = parts[0]
                        lname = ""
                    else:
                        return # Empty name
                    
                    relatives.append({
                        "FirstName": fname,
                        "LastName": lname,
                        "Relationship": relationship
                    })

            add_relative(row.get('Father', ''), 'Father')
            add_relative(row.get('Mother', ''), 'Mother')
            add_relative(row.get('Brother', ''), 'Brother')
            add_relative(row.get('Sister', ''), 'Sister')

            # Construct output object
            record = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": iso_birthday,
                "Age": age,
                "Relatives": relatives
            }
            result.append(record)

    # Output as JSON to stdout
    print(json.dumps(result, indent=3))

if __name__ == "__main__":
    process_csv("input/input.csv")