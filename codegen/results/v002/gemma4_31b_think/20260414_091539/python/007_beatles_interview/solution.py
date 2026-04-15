import csv
import json
from datetime import datetime

def calculate_age(birth_date, reference_date):
    """Calculates age based on birth date and a reference date."""
    age = reference_date.year - birth_date.year
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def parse_date(date_str):
    """Parses MM/DD/YYYY string to datetime.date object."""
    if not date_str or date_str.lower() == 'null':
        return None
    return datetime.strptime(date_str, '%m/%d/%Y').date()

def split_name(full_name):
    """Splits name into first and last components."""
    if not full_name or full_name.lower() == 'null':
        return None, None
    parts = full_name.split()
    # First name is the first part, Last name is the last part
    return parts[0], parts[-1]

def main():
    input_file = 'input/input.csv'
    ref_date = datetime(2025, 7, 1).date()
    
    results = []
    
    try:
        with open(input_file, mode='r', encoding='utf-8') as f:
            reader = csv.DictReader(f)
            for row in reader:
                # Name handling
                first_name, last_name = split_name(row['Name'])
                
                # Date handling
                birthday = parse_date(row['Birthday'])
                died_date = parse_date(row['Died'])
                
                # Age calculation: 
                # If person died, age is at time of death. Otherwise, as of July 1, 2025.
                target_date = died_date if died_date else ref_date
                age = calculate_age(birthday, target_date)
                
                # Relative handling
                relatives = []
                relationship_cols = ['Father', 'Mother', 'Brother', 'Sister']
                for col in relationship_cols:
                    rel_name = row[col]
                    if rel_name and rel_name.lower() != 'null':
                        rel_first, rel_last = split_name(rel_name)
                        relatives.append({
                            "FirstName": rel_first,
                            "LastName": rel_last,
                            "Relationship": col
                        })
                
                # Build object
                person = {
                    "FirstName": first_name,
                    "LastName": last_name,
                    "Birthday": birthday.strftime('%Y-%m-%d'),
                    "Age": age,
                    "Relatives": relatives
                }
                results.append(person)
                
    except FileNotFoundError:
        # In a real environment, handle error. Here we output empty list or exit.
        pass

    # Output JSON to stdout
    print(json.dumps(results, indent=2))

if __name__ == "__main__":
    main()