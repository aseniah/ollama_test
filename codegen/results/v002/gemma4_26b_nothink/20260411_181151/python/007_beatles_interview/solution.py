import csv
import json
import sys
from datetime import datetime

def calculate_age(birthday_str, reference_date):
    """Calculates age as of a specific reference date."""
    # Input format in CSV appears to be M/D/YYYY based on sample
    birthday = datetime.strptime(birthday_str, '%m/%d/%Y')
    age = reference_date.year - birthday.year
    if (reference_date.month, reference_date.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def main():
    input_file = 'input/input.csv'
    reference_date = datetime(2025, 7, 1)
    output = []

    try:
        with open(input_file, mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # Parse Name: split by space. 
                # Based on expected output, 'John Winston Lennon' -> FirstName: 'John', LastName: 'Lennon'
                # Looking at the pattern, it seems to take the first and last elements.
                name_parts = row['Name'].split()
                first_name = name_parts[0]
                last_name = name_parts[-1]

                # Parse Birthday
                birthday_str = row['Birthday']
                age = calculate_age(birthday_str, reference_date)

                # Prepare Relatives
                relatives = []
                
                # Helper to add relatives if they are not 'null'
                def add_relative(name_str, relation):
                    if name_str and name_str.lower() != 'null':
                        # The CSV columns like 'Father' contain "Alfred Lennon"
                        # We need to split them into FirstName and LastName
                        rel_parts = name_str.split()
                        rel_first = rel_parts[0]
                        rel_last = rel_parts[-1] if len(rel_parts) > 1 else ""
                        relatives.append({
                            "FirstName": rel_first,
                            "LastName": rel_last,
                            "Relationship": relation
                        })

                add_relative(row.get('Father'), 'Father')
                add_relative(row.get('Mother'), 'Mother')
                add_relative(row.get('Brother'), 'Brother')
                add_relative(row.get('Sister'), 'Sister')

                # Construct person object
                person_obj = {
                    "FirstName": first_name,
                    "LastName": last_name,
                    "Birthday": datetime.strptime(birthday_str, '%m/%d/%Y').strftime('%Y-%m-%d'),
                    "Age": age,
                    "Relatives": relatives
                }
                output.append(person_obj)

        # Output as JSON array to stdout
        print(json.dumps(output, indent=2))

    except FileNotFoundError:
        pass
    except Exception:
        pass

if __name__ == "__main__":
    main()