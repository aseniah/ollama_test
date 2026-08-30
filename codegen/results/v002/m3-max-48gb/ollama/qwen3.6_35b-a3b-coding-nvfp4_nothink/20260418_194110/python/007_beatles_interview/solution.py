import csv
import json
from datetime import datetime
from dateutil.relativedelta import relativedelta

def calculate_age(birthday_str, reference_date):
    """
    Calculate age as of reference_date given a birthday string.
    """
    # Parse the birthday string. The format seems to be M/D/YYYY or MM/DD/YYYY.
    # Let's try common formats.
    try:
        # Try standard formats
        birthday = datetime.strptime(birthday_str, "%m/%d/%Y")
    except ValueError:
        try:
            birthday = datetime.strptime(birthday_str, "%d/%m/%Y")
        except ValueError:
            raise ValueError(f"Could not parse date: {birthday_str}")
    
    age = reference_date.year - birthday.year - \
          ((reference_date.month, reference_date.day) < (birthday.month, birthday.day))
    return age

def main():
    reference_date = datetime(2025, 7, 1)
    
    results = []
    
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        for row in reader:
            name_parts = row['Name'].strip().split()
            if len(name_parts) >= 2:
                first_name = name_parts[0]
                last_name = ' '.join(name_parts[1:])
            elif len(name_parts) == 1:
                first_name = name_parts[0]
                last_name = ''
            else:
                first_name = ''
                last_name = ''
            
            birthday_str = row['Birthday'].strip()
            birthday_formatted = datetime.strptime(birthday_str, "%m/%d/%Y").strftime("%Y-%m-%d")
            age = calculate_age(birthday_str, reference_date)
            
            relatives = []
            
            father = row.get('Father', '').strip()
            if father and father.lower() != 'null':
                parts = father.split()
                if len(parts) >= 2:
                    relatives.append({
                        "FirstName": parts[0],
                        "LastName": ' '.join(parts[1:]),
                        "Relationship": "Father"
                    })
                else:
                    relatives.append({
                        "FirstName": parts[0],
                        "LastName": "",
                        "Relationship": "Father"
                    })
            
            mother = row.get('Mother', '').strip()
            if mother and mother.lower() != 'null':
                parts = mother.split()
                if len(parts) >= 2:
                    relatives.append({
                        "FirstName": parts[0],
                        "LastName": ' '.join(parts[1:]),
                        "Relationship": "Mother"
                    })
                else:
                    relatives.append({
                        "FirstName": parts[0],
                        "LastName": "",
                        "Relationship": "Mother"
                    })
            
            brother = row.get('Brother', '').strip()
            if brother and brother.lower() != 'null':
                parts = brother.split()
                if len(parts) >= 2:
                    relatives.append({
                        "FirstName": parts[0],
                        "LastName": ' '.join(parts[1:]),
                        "Relationship": "Brother"
                    })
                else:
                    relatives.append({
                        "FirstName": parts[0],
                        "LastName": "",
                        "Relationship": "Brother"
                    })
            
            sister = row.get('Sister', '').strip()
            if sister and sister.lower() != 'null':
                parts = sister.split()
                if len(parts) >= 2:
                    relatives.append({
                        "FirstName": parts[0],
                        "LastName": ' '.join(parts[1:]),
                        "Relationship": "Sister"
                    })
                else:
                    relatives.append({
                        "FirstName": parts[0],
                        "LastName": "",
                        "Relationship": "Sister"
                    })
            
            entry = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday_formatted,
                "Age": age,
                "Relatives": relatives
            }
            results.append(entry)
    
    print(json.dumps(results, indent=2))

if __name__ == "__main__":
    main()