import csv
import json
from datetime import date

def parse_date(date_str):
    """Parse a date string in M/D/YYYY or MM/DD/YYYY format to a date object."""
    parts = date_str.split('/')
    return date(int(parts[2]), int(parts[0]), int(parts[1]))

def calculate_age(birth_date, reference_date):
    """Calculate age in full years."""
    age = reference_date.year - birth_date.year
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def format_date(d):
    """Format a date as YYYY-MM-DD."""
    return d.strftime('%Y-%m-%d')

def split_name(full_name):
    """Split a full name into first and last name (last word is last name, first word is first name)."""
    parts = full_name.split()
    if len(parts) == 1:
        return parts[0], ""
    return parts[0], parts[-1]

def parse_relative(name_str, relationship):
    """Parse a relative's name into a dict, or return None if 'null'."""
    if name_str == 'null' or name_str.strip() == '':
        return None
    first, last = split_name(name_str)
    return {
        "FirstName": first,
        "LastName": last,
        "Relationship": relationship
    }

def main():
    reference_date = date(2025, 7, 1)
    results = []

    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        for row in reader:
            name = row['Name'].strip()
            first_name, last_name = split_name(name)
            birth_date = parse_date(row['Birthday'].strip())
            
            # Determine reference date for age calculation
            died_str = row['Died'].strip()
            if died_str and died_str != 'null':
                death_date = parse_date(died_str)
                ref_date = death_date
            else:
                ref_date = reference_date
            
            age = calculate_age(birth_date, ref_date)

            # Build relatives list
            relatives = []
            for col, rel_type in [('Father', 'Father'), ('Mother', 'Mother'), ('Brother', 'Brother'), ('Sister', 'Sister')]:
                rel = parse_relative(row[col].strip(), rel_type)
                if rel is not None:
                    relatives.append(rel)

            results.append({
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": format_date(birth_date),
                "Age": age,
                "Relatives": relatives
            })

    print(json.dumps(results, indent=2))

if __name__ == '__main__':
    main()