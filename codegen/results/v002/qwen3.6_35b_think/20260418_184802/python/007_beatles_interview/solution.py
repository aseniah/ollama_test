import csv
import json
from datetime import date

def parse_birthday(birthday_str):
    """Parse MM/DD/YYYY to YYYY-MM-DD"""
    if not birthday_str or birthday_str.strip() == '':
        return None
    parts = birthday_str.strip().split('/')
    month, day, year = int(parts[0]), int(parts[1]), int(parts[2])
    return date(year, month, day)

def calculate_age(birthday, as_of=date(2025, 7, 1)):
    """Calculate age as of July 1, 2025"""
    age = as_of.year - birthday.year
    if (as_of.month, as_of.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def split_name(full_name):
    """Split full name into first and last name"""
    parts = full_name.strip().split()
    if len(parts) >= 2:
        first_name = parts[0]
        last_name = ' '.join(parts[1:])
    elif len(parts) == 1:
        first_name = parts[0]
        last_name = ''
    else:
        first_name = ''
        last_name = ''
    return first_name, last_name

def build_relatives(father, mother, brother, sister):
    """Build relatives list from CSV columns"""
    relatives = []
    
    relationships = {
        'Father': father,
        'Mother': mother,
        'Brother': brother,
        'Sister': sister
    }
    
    for rel_name in relationships.values():
        if rel_name and rel_name.strip() != 'null' and rel_name.strip() != '':
            first_name, last_name = split_name(rel_name)
            relatives.append({
                "FirstName": first_name,
                "LastName": last_name,
                "Relationship": None  # We'll fill this in
            })
    
    # Map relationships in order
    rel_list = [
        ("Father", father),
        ("Mother", mother),
        ("Brother", brother),
        ("Sister", sister)
    ]
    
    result = []
    for relationship, name in rel_list:
        if name and name.strip() != 'null' and name.strip() != '':
            first_name, last_name = split_name(name)
            result.append({
                "FirstName": first_name,
                "LastName": last_name,
                "Relationship": relationship
            })
    
    return result

def main():
    results = []
    
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        for row in reader:
            full_name = row['Name']
            birthday_str = row['Birthday']
            died = row['Died']
            father = row['Father']
            mother = row['Mother']
            brother = row['Brother']
            sister = row['Sister']
            
            first_name, last_name = split_name(full_name)
            birthday = parse_birthday(birthday_str)
            age = calculate_age(birthday) if birthday else None
            relatives = build_relatives(father, mother, brother, sister)
            
            person = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday.strftime('%Y-%m-%d') if birthday else None,
                "Age": age,
                "Relatives": relatives
            }
            
            results.append(person)
    
    print(json.dumps(results, indent=2))

if __name__ == '__main__':
    main()