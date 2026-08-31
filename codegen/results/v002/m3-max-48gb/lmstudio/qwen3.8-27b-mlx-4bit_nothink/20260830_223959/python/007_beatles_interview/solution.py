<tool_call>
function=Write
<parameter=file_path
/solution/solution.py
</parameter>
<parameter=content>
import csv
import json
from datetime import date

def parse_date(date_str):
    """Parse a date string in MM/DD/YYYY format into a date object."""
    if not date_str or date_str == 'null':
        return None
    parts = date_str.split('/')
    month, day, year = int(parts[0]), int(parts[1]), int(parts[2])
    return date(year, month, day)

def calculate_age(birthday, reference_date):
    """Calculate age as of the reference date (July 1, 2025)."""
    if birthday is None:
        return None
    age = reference_date.year - birthday.year
    if (reference_date.month, reference_date.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def main():
    reference_date = date(2025, 7, 1)
    
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        results = []
        
        for row in reader:
            # Parse name into first and last name
            name_parts = row['Name'].strip().split()
            first_name = name_parts[0]
            last_name = ' '.join(name_parts[1:])
            
            # Parse dates
            birthday = parse_date(row['Birthday'])
            died = parse_date(row['Died'])
            
            # Calculate age
            age = calculate_age(birthday, reference_date)
            
            # Format birthday as YYYY-MM-DD
            birthday_str = birthday.strftime('%Y-%m-%d') if birthday else ""
            
            # Build relatives list
            relatives = []
            
            if row['Father'] and row['Father'] != 'null':
                rel_name = row['Father'].strip()
                rel_parts = rel_name.split()
                rel_first = rel_parts[0]
                rel_last = ' '.join(rel_parts[1:])
                relatives.append({
                    "FirstName": rel_first,
                    "LastName": rel_last,
                    "Relationship": "Father"
                })
            
            if row['Mother'] and row['Mother'] != 'null':
                rel_name = row['Mother'].strip()
                rel_parts = rel_name.split()
                rel_first = rel_parts[0]
                rel_last = ' '.join(rel_parts[1:])
                relatives.append({
                    "FirstName": rel_first,
                    "LastName": rel_last,
                    "Relationship": "Mother"
                })
            
            if row['Brother'] and row['Brother'] != 'null':
                rel_name = row['Brother'].strip()
                rel_parts = rel_name.split()
                rel_first = rel_parts[0]
                rel_last = ' '.join(rel_parts[1:])
                relatives.append({
                    "FirstName": rel_first,
                    "LastName": rel_last,
                    "Relationship": "Brother"
                })
            
            if row['Sister'] and row['Sister'] != 'null':
                rel_name = row['Sister'].strip()
                rel_parts = rel_name.split()
                rel_first = rel_parts[0]
                rel_last = ' '.join(rel_parts[1:])
                relatives.append({
                    "FirstName": rel_first,
                    "LastName": rel_last,
                    "Relationship": "Sister"
                })
            
            entry = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday_str,
                "Age": age,
                "Relatives": relatives
            }
            
            results.append(entry)
    
    print(json.dumps(results, indent=2))

if __name__ == '__main__':
    main()

</parameter>
</function>
</tool_call>