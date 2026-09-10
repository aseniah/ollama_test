import csv
import json
from datetime import datetime

def calculate_age(birth_date, reference_date):
    """
    Calculates age based on birth date and a reference date.
    """
    age = reference_date.year - birth_date.year
    # Subtract 1 if the birthday has not occurred yet in the reference year
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def parse_name(name_string):
    """
    Splits a name into first and last name.
    Assumes first word is FirstName and last word is LastName.
    """
    parts = name_string.strip().split()
    if not parts:
        return "", ""
    first = parts[0]
    last = parts[-1] if len(parts) > 1 else ""
    return first, last

def main():
    input_file = 'input/input.csv'
    target_date = datetime(2025, 7, 1)
    
    results = []
    
    try:
        with open(input_file, mode='r', encoding='utf-8') as f:
            reader = csv.DictReader(f)
            for row in reader:
                # Parse primary person's name
                first_name, last_name = parse_name(row['Name'])
                
                # Parse birthday (Format: M/D/YYYY)
                birthday_str = row['Birthday']
                birthday_dt = datetime.strptime(birthday_str, '%m/%d/%Y')
                
                # Determine reference date for age calculation
                # Use death date if available, otherwise use July 1, 2025
                died_str = row.get('Died')
                if died_str and died_str.lower() != 'null':
                    ref_date = datetime.strptime(died_str, '%m/%d/%Y')
                else:
                    ref_date = target_date
                
                age = calculate_age(birthday_dt, ref_date)
                
                # Process relatives
                relatives = []
                relationship_cols = ['Father', 'Mother', 'Brother', 'Sister']
                for col in relationship_cols:
                    rel_name_str = row.get(col)
                    if rel_name_str and rel_name_str.lower() != 'null':
                        r_first, r_last = parse_name(rel_name_str)
                        relatives.append({
                            "FirstName": r_first,
                            "LastName": r_last,
                            "Relationship": col
                        })
                
                # Build person object
                person = {
                    "FirstName": first_name,
                    "LastName": last_name,
                    "Birthday": birthday_dt.strftime('%Y-%m-%d'),
                    "Age": age,
                    "Relatives": relatives
                }
                results.append(person)
                
    except FileNotFoundError:
        # The prompt implies the file exists, but good for robustness
        pass

    # Output the JSON array to stdout
    print(json.dumps(results, indent=2))

if __name__ == "__main__":
    main()