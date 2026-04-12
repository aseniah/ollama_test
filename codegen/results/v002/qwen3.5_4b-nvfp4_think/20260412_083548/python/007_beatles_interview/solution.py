import csv
from datetime import datetime, timedelta
import json

def parse_date(date_str):
    """Parse date string in MM/DD/YY format to datetime"""
    if not date_str:
        return None
    try:
        m, d, y = map(int, date_str.split('/'))
        # Handle year - for years before 100, assume 19yy or 20yy
        y = int(y)
        if y < 100:
            year = (y + 2000) % 100  # Convert to correct century
        return datetime(year, m, d)
    except ValueError:
        return None

def calculate_age(birthday_str, reference_date):
    """Calculate age as of the reference date"""
    birthday = parse_date(birthday_str)
    if not birthday:
        return 0
    
    # Calculate total days between birthday and reference date
    delta = reference_date - birthday
    years = delta.days // 365.25
    return int(years)

def format_relative_date(date_str):
    """Format birth year in MM/DD/YY to YYYY-MM-DD"""
    if not date_str:
        return None
    parts = date_str.split('/')
    month = f"{int(parts[0]):02d}"
    day = f"{int(parts[1]):02d}"
    year = int(parts[2])
    # Assume full 4-digit year
    if len(year) == 2:
        year = (year - 1900) % 100 + 2000
    else:
        year = f"{year}"
    return f"{year}-{month}-{day}"

def parse_csnm():
    """Read and process CSV file"""
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        data = list(reader)
    return data

def calculate_age_from_reference(birthday, reference_year):
    """Calculate age from birthday year (as of given year)"""
    bday_date = datetime(int(birthday.split('-')[0]), int(birthday.split('-')[1]), int(birthday.split('-')[2]))
    ref_date = datetime(reference_year, 7, 1)
    days_diff = (ref_date - bday_date).days
    
    # Simple calculation: if birthday has passed in reference year, age is difference
    birth_year = bday_date.year
    age = reference_year - birth_year
    return age

def transform_data():
    """Transform CSV data to expected JSON format"""
    
    # Reference date for age calculation
    reference_date = datetime(2025, 7, 1)
    
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        rows = list(reader)
    
    output = []
    
    for row in rows:
        # Get personal info and birthday
        first_name = row.get('Name', '').split()[0]
        last_name = ' '.join(row['Name'].split()[1:]) if len(row['Name'].split()) > 1 else ''
        
        # Parse birth date from MM/DD/YY format
        birthday_str = row.get('Birthday', '')
        
        # Convert to YYYY-MM-DD format for display
        bday_parts = birthday_str.split('/')
        if bday_parts:
            year, month, day = map(int, bday_parts)
            formatted_bday = f"{year}-{month:02d}-{day:02d}"
        else:
            formatted_bday = ''
        
        # Calculate age
        age = calculate_age_from_reference(birthday_str, 2025)
        
        # Get relatives data (skip if 'null')
        relatives = []
        def add_relative(rel_type, name_val, first, last):
            rel_name = f"{first} {last}"
            return {"FirstName": rel_name, "LastName": " ".join(last.split()) if len(last.split()) > 1 else last, "Relationship": rel_type}
        
        relatives_data = [
            {'Name': row.get('Father', ''), 'Rel': 'Father'},
            {'Name': row.get('Mother', ''), 'Rel': 'Mother'},
            {'Name': row.get('Brother', ''), 'Rel': 'Brother'},
            {'Name': row.get('Sister', ''), 'Rel': 'Sister'}
        ]
        
        for rel in relatives_data:
            name = rel['Name']
            if not name or name == 'null':
                continue
            
            parts = name.split()
            if len(parts) < 1:
                continue
                
            first_name = parts[0]
            last_name = parts[1] if len(parts) > 1 else ''
            
            relatives.append({
                "FirstName": first_name,
                "LastName": last_name or '',
                "Relationship": rel['Rel']
            })
        
        # Build output object
        person = {
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": formatted_bday,
            "Age": age,
            "Relatives": relatives
        }
        output.append(person)
    
    return output

def main():
    try:
        result = transform_data()
        print(json.dumps(result))
    except FileNotFoundError:
        print("Error: input/input.csv not found")
        return 1
    except Exception as e:
        print(f"Error: {e}")
        return 1

if __name__ == "__main__":
    main()