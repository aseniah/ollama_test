import csv
import json
from datetime import datetime

def calculate_age(birth_date, reference_date, death_date_str):
    """
    Calculates age based on the minimum of the death date (if provided) 
    and the reference date.
    """
    # Determine the end date for age calculation
    if death_date_str and death_date_str.lower() != 'null':
        end_date = datetime.strptime(death_date_str, "%m/%d/%Y")
    else:
        end_date = reference_date

    age = end_date.year - birth_date.year
    # Subtract one year if the birthday hasn't occurred yet in the end_date year
    if (end_date.month, end_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def split_name(full_name):
    """
    Splits full name into first and last name. 
    Takes the first word as FirstName and the last word as LastName.
    """
    parts = full_name.strip().split()
    if not parts:
        return "", ""
    if len(parts) == 1:
        return parts[0], ""
    return parts[0], parts[-1]

def main():
    input_csv_path = 'input/input.csv'
    reference_date = datetime(2025, 7, 1)
    
    results = []

    try:
        with open(input_csv_path, mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # Process Main Person
                full_name = row['Name']
                first_name, last_name = split_name(full_name)
                
                birth_date = datetime.strptime(row['Birthday'], "%m/%d/%Y")
                age = calculate_age(birth_date, reference_date, row['Died'])
                
                # Process Relatives
                relatives = []
                relationship_cols = ['Father', 'Mother', 'Brother', 'Sister']
                for col in relationship_cols:
                    rel_name_full = row[col]
                    if rel_name_full and rel_name_full.lower() != 'null':
                        rel_first, rel_last = split_name(rel_name_full)
                        relatives.append({
                            "FirstName": rel_first,
                            "LastName": rel_last,
                            "Relationship": col
                        })
                
                # Build Object
                person_obj = {
                    "FirstName": first_name,
                    "LastName": last_name,
                    "Birthday": birth_date.strftime("%Y-%m-%d"),
                    "Age": age,
                    "Relatives": relatives
                }
                results.append(person_obj)

        # Output the result as JSON array to stdout
        print(json.dumps(results, indent=2))

    except FileNotFoundError:
        pass

if __name__ == "__main__":
    main()