import csv
import json
from datetime import date

def parse_date(date_str):
    """Convert MM/DD/YYYY to YYYY-MM-DD"""
    parts = date_str.split('/')
    return f"{parts[2]}-{parts[1]}-{parts[0]}"

def calculate_age(birthday, reference_date=date(2025, 7, 1)):
    """Calculate age as of reference date (July 1, 2025)"""
    birth_date = date(
        int(birthday.split('-')[0]),
        int(birthday.split('-')[1]),
        int(birthday.split('-')[2])
    )
    
    age = reference_date.year - birth_date.year
    if reference_date < (birth_date.replace(year=reference_date.year)):
        age -= 1
    
    return age

def parse_name(name):
    """Parse full name into first and last name"""
    parts = name.split()
    if len(parts) >= 2:
        first = parts[0]
        last = parts[-1]
    else:
        first = parts[0]
        last = parts[-1]
    
    return first, last

def build_relatives(relatives_data):
    """Build relatives array from CSV columns"""
    relatives = []
    for col in ['Father', 'Mother', 'Brother', 'Sister']:
        if relatives_data.get(col) and relatives_data[col].lower() != 'null':
            first, last = parse_name(relatives_data[col])
            relatives.append({
                "FirstName": first,
                "LastName": last,
                "Relationship": col
            })
    return relatives

def main():
    relatives_data = {}
    
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        for row in reader:
            name = row['Name']
            birthday = parse_date(row['Birthday'])
            first, last = parse_name(name)
            
            relatives_data[name] = {
                "FirstName": first,
                "LastName": last,
                "Birthday": birthday,
                "Age": calculate_age(birthday),
                "Relatives": build_relatives(row)
            }
            
            # Process relatives from CSV columns
            if row.get('Father') and row['Father'].lower() != 'null':
                relatives_data[name]["Father"] = parse_name(row['Father'])
            if row.get('Mother') and row['Mother'].lower() != 'null':
                relatives_data[name]["Mother"] = parse_name(row['Mother'])
            if row.get('Brother') and row['Brother'].lower() != 'null':
                relatives_data[name]["Brother"] = parse_name(row['Brother'])
            if row.get('Sister') and row['Sister'].lower() != 'null':
                relatives_data[name]["Sister"] = parse_name(row['Sister'])
    
    output = []
    for name, data in relatives_data.items():
        output.append(data)
    
    print(json.dumps(output))

if __name__ == "__main__":
    main()