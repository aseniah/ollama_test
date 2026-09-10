import csv
import json
from datetime import datetime

def calculate_age(birth_date, reference_date):
    """Calculates age based on a reference date."""
    age = reference_date.year - birth_date.year
    # Subtract one if the reference date is before the birthday in that year
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def parse_name(full_name):
    """Splits name into FirstName and LastName."""
    if not full_name or full_name.lower() == 'null':
        return None
    parts = full_name.split()
    if len(parts) == 0:
        return None
    return {
        "FirstName": parts[0],
        "LastName": parts[-1] if len(parts) > 1 else ""
    }

def main():
    input_file = 'input/input.csv'
    ref_date = datetime(2025, 7, 1)
    
    results = []
    
    try:
        with open(input_file, mode='r', encoding='utf-8') as f:
            reader = csv.DictReader(f)
            for row in reader:
                # Parse dates
                # CSV format is M/D/YYYY
                birth_dt = datetime.strptime(row['Birthday'], '%m/%d/%Y')
                
                # Determine reference date for age (Death date if exists, otherwise July 1, 2025)
                died_val = row.get('Died')
                if died_val and died_val.lower() != 'null':
                    effective_ref_date = datetime.strptime(died_val, '%m/%d/%Y')
                else:
                    effective_ref_date = ref_date
                
                age = calculate_age(birth_dt, effective_ref_date)
                
                # Parse person name
                person_name = parse_name(row['Name'])
                
                # Build relatives list
                relatives = []
                relative_cols = ['Father', 'Mother', 'Brother', 'Sister']
                for col in relative_cols:
                    rel_val = row.get(col)
                    if rel_val and rel_val.lower() != 'null':
                        name_data = parse_name(rel_val)
                        if name_data:
                            name_data['Relationship'] = col
                            relatives.append(name_data)
                
                # Construct final object
                entry = {
                    "FirstName": person_name["FirstName"],
                    "LastName": person_name["LastName"],
                    "Birthday": birth_dt.strftime('%Y-%m-%d'),
                    "Age": age,
                    "Relatives": relatives
                }
                results.append(entry)
                
    except FileNotFoundError:
        pass

    # Output the resulting JSON array to stdout
    print(json.dumps(results, indent=2))

if __name__ == "__main__":
    main()