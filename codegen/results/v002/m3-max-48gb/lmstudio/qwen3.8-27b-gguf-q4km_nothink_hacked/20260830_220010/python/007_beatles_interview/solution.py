import csv
import json
from datetime import date

def calculate_age(birth_date, reference_date):
    """Calculate full years completed between birth_date and reference_date."""
    if reference_date.year > birth_date.year:
        age = reference_date.year - birth_date.year
        # If reference date is before birthday this year, subtract 1
        if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
            age -= 1
    else:
        age = 0
    return age

def parse_date(date_str):
    """Parse M/D/YYYY or MM/DD/YYYY date string to date object."""
    parts = date_str.split('/')
    month = int(parts[0])
    day = int(parts[1])
    year = int(parts[2])
    return date(year, month, day)

def parse_name(full_name):
    """Split full name into first and last name."""
    parts = full_name.strip().split()
    if len(parts) == 2:
        return parts[0], parts[1]
    elif len(parts) > 2:
        # If there's a middle name, use first and last
        return parts[0], parts[-1]
    else:
        return full_name.strip(), ""

def main():
    input_file = 'input/input.csv'
    reference_date = date(2025, 7, 1)
    
    people = []
    
    with open(input_file, 'r', newline='') as f:
        reader = csv.DictReader(f)
        for row in reader:
            # Parse name
            first_name, last_name = parse_name(row['Name'])
            
            # Parse birthday
            birth_date = parse_date(row['Birthday'])
            birthday_str = birth_date.strftime('%Y-%m-%d')
            
            # Determine reference date for age
            died_str = row['Died'].strip()
            if died_str and died_str.lower() != 'null':
                death_date = parse_date(died_str)
                reference = death_date
            else:
                reference = reference_date
            
            age = calculate_age(birth_date, reference)
            
            # Build relatives
            relatives = []
            
            for field, relationship in [('Father', 'Father'), ('Mother', 'Mother'), ('Brother', 'Brother'), ('Sister', 'Sister')]:
                value = row[field].strip()
                if value and value.lower() != 'null':
                    rel_first, rel_last = parse_name(value)
                    relatives.append({
                        "FirstName": rel_first,
                        "LastName": rel_last,
                        "Relationship": relationship
                    })
            
            person = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday_str,
                "Age": age,
                "Relatives": relatives
            }
            people.append(person)
    
    print(json.dumps(people, indent=2))

if __name__ == '__main__':
    main()