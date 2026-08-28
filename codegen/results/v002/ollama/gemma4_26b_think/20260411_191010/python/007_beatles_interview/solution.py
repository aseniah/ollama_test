import csv
import json
from datetime import datetime

def solve():
    """
    Reads input/input.csv, transforms the data based on inferred rules from
    input/expected_format.json, and outputs a JSON array to stdout.
    """
    target_date = datetime(2025, 7, 1).date()
    results = []
    
    try:
        with open('input/input.csv', mode='r', encoding='utf-8') as f:
            reader = csv.DictReader(f)
            for row in reader:
                # 1. Parse the primary subject's Name
                # Rule: Extract the first word as FirstName and the last word as LastName
                name_parts = row['Name'].split()
                if not name_parts:
                    continue
                first_name = name['Name'].split()[0] if 'Name' in row else "" # Safety check
                # Using the split parts from the row name
                first_name = name_parts[0]
                last_name = name_parts[-1]
                
                # 2. Parse Birthday (format M/D/YYYY)
                # Expected format in JSON: YYYY-MM-DD
                birthday_str = row['Birthday']
                birthday = datetime.strptime(birthday_str, '%m/%d/%Y').date()
                
                # 3. Determine the reference date for Age calculation
                # Rule: If 'Died' is provided, age is calculated at death.
                # If 'Died' is null, age is calculated as of July 1, 2025.
                died_str = row['Died']
                if died_str == 'null' or not died_str:
                    ref_date = target_date
                else:
                    ref_date = datetime.strptime(died_str, '%m/%d/%Y').date()
                
                # 4. Calculate Age
                # Age = (Ref Date Year - Birth Year) - 1 if the birthday hasn't occurred yet in the ref year
                age = ref_date.year - birthday.year - (
                    (ref_date.month, ref_date.day) < (birthday.month, birthday.day)
                )
                
                # 5. Parse Relatives
                # The CSV contains Father, Mother, Brother, Sister columns.
                # Rule: If the field is not 'null', split the name into FirstName/LastName
                relatives = []
                rel_fields = [
                    ('Father', 'Father'), 
                    ('Mother', 'Mother'), 
                    ('Brother', 'Brother'), 
                    ('Sister', 'Sister')
                ]
                
                for csv_col, relationship in rel_fields:
                    rel_val = row.get(csv_col, 'null')
                    if rel_val != 'null' and rel_val:
                        rel_parts = rel_val.split()
                        if rel_parts:
                            rel_rel_first = rel_parts[0]
                            rel_rel_last = rel_parts[-1]
                            relatives.append({
                                "FirstName": rel_rel_first,
                                "LastName": rel_rel_last,
                                "Relationship": relationship
                            })
                
                # 6. Construct the person object
                person_obj = {
                    "FirstName": first_name,
                    "LastName": last_name,
                    "Birthday": birthday.strftime('%Y-%m-%d'),
                    "Age": age,
                    "Relatives": relatives
                }
                results.append(person_obj)
        
        # Output the final JSON array to stdout
        print(json.dumps(results, indent=2))

    except FileNotFoundError:
        pass # Handle error if file does not exist
    except Exception:
        pass # Silently fail as per requirements (only output JSON)

if __name__ == "__main__":
    solve()