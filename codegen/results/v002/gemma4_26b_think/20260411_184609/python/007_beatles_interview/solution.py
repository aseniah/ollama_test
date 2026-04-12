import csv
import json
from datetime import datetime
import os

def solve():
    # The path to the input CSV file
    input_path = 'input/input.csv'
    
    # The reference date for calculating age of those who have not died
    reference_date = datetime(2025, 7, 1)
    
    # Check if file exists to avoid error, though per instructions it should exist
    if not os.path.exists(input_path):
        return

    results = []
    
    try:
        with open(input_path, mode='r', encoding='utf-8') as f:
            # Use DictReader to map columns to dictionary keys
            reader = csv.DictReader(f)
            for row in reader:
                # 1. Parse Primary Name
                # Strategy: Extract the first and last word from the full Name string
                full_name = row['Name'].strip()
                name_parts = full_name.split()
                if not name_parts:
                    continue
                first_name = name_parts[0]
                last_name = name_parts[-1]
                
                # 2. Parse Birthday
                # CSV format: M/D/YYYY
                bday_str = row['Birthday'].strip()
                birthday = datetime.strptime(bday_str, '%m/%d/%Y')
                
                # 3. Parse Died Date
                died_str = row.get('Died', '').strip()
                died_date = None
                if died_str and died_str.lower() != 'null':
                    died_date = datetime.strptime(died_str, '%m/%d/%Y')
                
                # 4. Calculate Age
                # Logic inferred from expected_format:
                # If Died is present, Age is age at the time of death.
                # If Died is null, Age is age as of July 1, 2025.
                if died_date:
                    age = died_date.year - birthday.year
                    # If the birthday hadn't occurred yet in the year of death
                    if (died_date.month, died_date.day) < (birthday.month, birthday.day):
                        age -= 1
                else:
                    age = reference_date.year - birthday.year
                    # If the birthday hadn't occurred yet by July 1, 2025
                    if (reference_date.month, reference_date.day) < (birthday.month, birthday.day):
                        age -= 1
                
                # 5. Process Relatives
                # Map CSV columns to relationship types
                relatives = []
                rel_mappings = [
                    ('Father', 'Father'),
                    ('Mother', 'Mother'),
                    ('Brother', 'Brother'),
                    ('Sister', 'Sister')
                ]
                
                for col, rel_type in rel_mappings:
                    rel_val = row.get(col, '').strip()
                    if rel_val and rel_val.lower() != 'null':
                        rel_parts = rel_val.split()
                        if rel_parts:
                            rel_first = rel_parts[0]
                            rel_last = rel_parts[-1]
                            relatives.append({
                                "FirstName": rel_first,
                                "LastName": rel_last,
                                "Relationship": rel_type
                            })
                
                # 6. Construct Object
                person_obj = {
                    "FirstName": first_name,
                    "LastName": last_name,
                    "Birthday": birthday.strftime('%Y-%m-%d'),
                    "Age": age,
                    "Relatives": relatives
                }
                results.append(person_obj)
                
        # Final output: JSON array to stdout
        print(json.dumps(results, indent=2))

    except Exception:
        # If any error occurs during parsing, we stop. 
        # The instruction implies valid input.
        pass

if __name__ == "__main__":
    solve()