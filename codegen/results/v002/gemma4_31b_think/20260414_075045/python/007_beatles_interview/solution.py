import csv
import json
from datetime import datetime

def parse_date(date_str):
    """Parses date strings in MM/DD/YYYY format."""
    if not date_str or date_str.lower() == 'null':
        return None
    return datetime.strptime(date_str, '%m/%d/%Y')

def calculate_age(birth_date, reference_date):
    """Calculates age relative to a reference date."""
    age = reference_date.year - birth_date.year
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def split_name(full_name):
    """Splits a full name into first and last name (first word and last word)."""
    if not full_name or full_name.lower() == 'null':
        return None, None
    parts = full_name.split()
    if not parts:
        return None, None
    # Based on expected_format.json, take the first and last word
    return parts[0], parts[-1]

def main():
    input_file = 'input/input.csv'
    target_date = datetime(2025, 7, 1)
    
    results = []
    
    try:
        with open(input_file, mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # Parse main person details
                full_name = row['Name']
                first_name, last_name = split_name(full_name)
                
                birth_dt = parse_date(row['Birthday'])
                died_dt = parse_date(row['Died'])
                
                # Age calculation: 
                # If person died, age at death. Otherwise, age as of July 1, 2025.
                ref_dt = died_dt if died_dt else target_date
                age = calculate_age(birth_dt, ref_dt) if birth_dt and ref_dt else None
                
                # Birthday format as YYYY-MM-DD
                birthday_iso = birth_dt.strftime('%Y-%m-%d') if birth_dt else None
                
                # Process relatives
                relatives = []
                rel_cols = ['Father', 'Mother', 'Brother', 'Sister']
                for col in rel_cols:
                    rel_name = row[col]
                    if rel_name and rel_name.lower() != 'null':
                        rel_first, rel_last = split_name(rel_name)
                        relatives.append({
                            "FirstName": rel_first,
                            "LastName": rel_last,
                            "Relationship": col
                        })
                
                # Construct the person object
                person = {
                    "FirstName": first_name,
                    "LastName": last_name,
                    "Birthday": birthday_iso,
                    "Age": age,
                    "Relatives": relatives
                }
                results.append(person)
                
    except FileNotFoundError:
        # In a real scenario we might handle this, but based on instructions 
        # we assume the file exists.
        pass

    # Output the resulting JSON array to stdout
    print(json.dumps(results, indent=2))

if __name__ == "__main__":
    main()