import csv
import json
from datetime import datetime

def solve():
    input_file = 'input/input.csv'
    target_date = datetime(2025, 7, 1).date()
    
    people_list = []

    try:
        with open(input_file, mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # Parse Birthday
                birth_str = row['Birthday']
                birth_date = datetime.strptime(birth_str, '%m/%d/%Y').date()
                
                # Parse Died
                died_str = row['Died']
                died_date = None
                if died_str and died_str.lower() != 'null':
                    died_date = datetime.strptime(died_str, '%m/%d/%Y').date()
                
                # Calculate Age based on the inferred rule:
                # If died_date exists: Age = year_of_death - year_of_birth (adjusted for birthday)
                # If died_date is null: Age = 2025 - year_of_birth (adjusted for July 1st)
                if died_date:
                    # Compare death date to birthday to see if they had already passed it in their death year
                    if (died_date.month, died_date.day) < (birth_date.month, birth_date.day):
                        age = died_date.year - birth_date.year - 1
                    else:
                        age = died_date.year - birth_date.year
                else:
                    # Compare target date (July 1, 2025) to birthday
                    if (target_date.month, target_date.day) < (birth_date.month, birth_date.day):
                        age = target_date.year - birth_date.year - 1
                    else:
                        age = target_date.year - birth_date.year
                
                # Parse Main Person Name
                name_parts = row['Name'].split()
                first_name = name_parts[0]
                last_name = name_parts[-1]
                
                # Prepare the person object
                person_obj = {
                    "FirstName": first_name,
                    "LastName": last_name,
                    "Birthday": birth_date.strftime('%Y-%m-%d'),
                    "Age": age,
                    "Relatives": []
                }
                
                # Parse Relatives (Father, Mother, Brother, Sister)
                for relation in ['Father', 'Mother', 'Brother', 'Sister']:
                    rel_name = row.get(relation)
                    if rel_name and rel_name.lower() != 'null':
                        rel_parts = rel_name.split()
                        rel_first = rel_parts[0]
                        rel_last = rel_parts[-1]
                        person_obj["Relatives"].append({
                            "FirstName": rel_first,
                            "LastName": rel_last,
                            "Relationship": relation
                        })
                
                people_list.append(person_obj)
                
        # Output the JSON array to stdout
        print(json.dumps(people_list, indent=2))
        
    except FileNotFoundError:
        pass

if __name__ == "__main__":
    solve()