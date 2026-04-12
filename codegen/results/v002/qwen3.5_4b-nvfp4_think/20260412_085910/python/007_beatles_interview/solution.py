import csv
import json
from datetime import datetime, date

def main():
    # Read input file
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        data = list(reader)
    
    # Define reference date for age calculation
    ref_date = date(2025, 7, 1)
    
    result = []
    
    for row in data:
        name = row['Name']
        birthday_str = row['Birthday']
        
        # Split name into first and last name
        first_name, last_name = name.split(' ', 1)
        
        # Parse birthday from MM/DD/YYYY format
        birth_date = datetime.strptime(birthday_str, '%m/%d/%Y').date()
        
        # Calculate age as of July 1, 2025
        age = (ref_date - birth_date).days // 365 + (ref_date.year - birth_date.year)
        
        # Format birthday as YYYY-MM-DD
        formatted_birthday = birth_date.strftime('%Y-%m-%d')
        
        # Parse relatives
        relatives = []
        father = row.get('Father', '').split(' ') if row.get('Father') else []
        mother = row.get('Mother', '').split(' ') if row.get('Mother') else []
        brother = row.get('Brother', '').split(' ') if row.get('Brother') else []
        sister = row.get('Sister', '').split(' ') if row.get('Sister') else []
        
        # Add father
        if father:
            relatives.append({
                "FirstName": father[0],
                "LastName": father[-1] if len(father) > 1 else '',
                "Relationship": "Father"
            })
        
        # Add mother
        if mother:
            relatives.append({
                "FirstName": mother[0],
                "LastName": mother[-1] if len(mother) > 1 else '',
                "Relationship": "Mother"
            })
        
        # Add brother (if not null or empty)
        if brother and 'null' not in str(brother):
            for b in brother:
                relatives.append({
                    "FirstName": b.split(' ')[0],
                    "LastName": b.split(' ')[-1] if len(b.split(' ')) > 1 else '',
                    "Relationship": "Brother"
                })
        
        # Add sister (if not null or empty)
        if sister and 'null' not in str(sister):
            for s in sister:
                relatives.append({
                    "FirstName": s.split(' ')[0],
                    "LastName": s.split(' ')[-1] if len(s.split(' ')) > 1 else '',
                    "Relationship": "Sister"
                })
        
        # Prepare output object
        person = {
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": formatted_birthday,
            "Age": age,
            "Relatives": relatives
        }
        result.append(person)
    
    # Output JSON array to stdout
    print(json.dumps(result))

if __name__ == '__main__':
    main()