import csv
import json
from datetime import date

# Reference date for age calculation
REFERENCE_DATE = date(2025, 7, 1)

def calculate_age(birthday_str: str) -> int:
    """
    Calculate age as of REFERENCE_DATE given a birthday string 'M/D/YYYY'.
    """
    try:
        month, day, year = map(int, birthday_str.split('/'))
        birthday = date(year, month, day)
    except ValueError:
        return 0
    
    age = REFERENCE_DATE.year - birthday.year
    if (REFERENCE_DATE.month, REFERENCE_DATE.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def parse_name(name_str: str) -> str:
    """
    Extract the first name from a full name string.
    The CSV contains "First Middle Last" or "First Last".
    The expected output uses just the first word as FirstName.
    """
    return name_str.strip().split()[0]

def main():
    input_file = 'input/input.csv'
    
    people = []
    
    with open(input_file, 'r', newline='') as f:
        reader = csv.DictReader(f)
        for row in reader:
            full_name = row['Name']
            birthday_str = row['Birthday']
            father_name = row['Father']
            mother_name = row['Mother']
            brother_name = row['Brother']
            sister_name = row['Sister']
            
            # Calculate age
            age = calculate_age(birthday_str)
            
            # Parse first name
            first_name = parse_name(full_name)
            
            # Parse last name: The expected output seems to take the last word of the name.
            # John Winston Lennon -> Lennon
            # James Paul McCartney -> McCartney
            # Ringo Starr -> Starr
            # George Harrison -> Harrison
            name_parts = full_name.strip().split()
            last_name = name_parts[-1]
            
            relatives = []
            
            # Process Father
            if father_name and father_name.strip().lower() != 'null':
                fname = parse_name(father_name)
                lname = father_name.strip().split()[-1]
                relatives.append({
                    "FirstName": fname,
                    "LastName": lname,
                    "Relationship": "Father"
                })
            
            # Process Mother
            if mother_name and mother_name.strip().lower() != 'null':
                fname = parse_name(mother_name)
                lname = mother_name.strip().split()[-1]
                relatives.append({
                    "FirstName": fname,
                    "LastName": lname,
                    "Relationship": "Mother"
                })
            
            # Process Brother
            if brother_name and brother_name.strip().lower() != 'null':
                fname = parse_name(brother_name)
                lname = brother_name.strip().split()[-1]
                relatives.append({
                    "FirstName": fname,
                    "LastName": lname,
                    "Relationship": "Brother"
                })
            
            # Process Sister
            if sister_name and sister_name.strip().lower() != 'null':
                fname = parse_name(sister_name)
                lname = sister_name.strip().split()[-1]
                relatives.append({
                    "FirstName": fname,
                    "LastName": lname,
                    "Relationship": "Sister"
                })
            
            person = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday_str if '/' in birthday_str else birthday_str,
                "Age": age,
                "Relatives": relatives
            }
            
            people.append(person)
            
    # Output JSON array
    print(json.dumps(people, indent=2))

if __name__ == '__main__':
    main()