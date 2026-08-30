import csv
import json
from datetime import datetime

def calculate_age(birth_date, death_date=None):
    # Target date is July 1, 2025
    reference_date = datetime(2025, 7, 1)
    end_date = death_date if death_date else reference_date
    
    age = end_date.year - birth_date.year
    # Adjust if birthday hasn't occurred yet in the end_date year
    if (end_date.month, end_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def parse_name(full_name):
    if not full_name or full_name == "null":
        return None
    parts = full_name.split()
    return {
        "FirstName": parts[0],
        "LastName": parts[-1]
    }

def main():
    input_file = 'input/input.csv'
    
    people = []
    
    try:
        with open(input_file, mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # Parse Name
                name_parts = parse_name(row['Name'])
                
                # Parse Birthday
                bday_str = row['Birthday']
                bday_dt = datetime.strptime(bday_str, "%m/%d/%Y")
                
                # Parse Death Date
                died_str = row['Died']
                died_dt = None
                if died_str and died_str != "null":
                    died_dt = datetime.strptime(died_str, "%m/%d/%Y")
                
                # Calculate Age
                age = calculate_age(bday_dt, died_dt)
                
                # Process Relatives
                relatives = []
                for rel_type in ['Father', 'Mother', 'Brother', 'Sister']:
                    col_name = rel_type
                    rel_full_name = row.get(col_name)
                    if rel_full_name and rel_full_name != "null":
                        rel_parts = parse_name(rel_full_name)
                        relatives.append({
                            "FirstName": rel_parts["FirstName"],
                            "LastName": rel_parts["LastName"],
                            "Relationship": rel_type
                        })
                
                # Build person object
                person = {
                    "FirstName": name_parts["FirstName"],
                    "LastName": name_parts["LastName"],
                    "Birthday": bday_dt.strftime("%Y-%m-%d"),
                    "Age": age,
                    "Relatives": relatives
                }
                people.append(person)
                
    except FileNotFoundError:
        pass

    # Output result as JSON array to stdout
    print(json.dumps(people, indent=2))

if __name__ == "__main__":
    main()