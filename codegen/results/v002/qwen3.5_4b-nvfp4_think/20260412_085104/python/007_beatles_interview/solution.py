import json
import csv
from datetime import datetime

def parse_birthday(bd):
    if not bd or bd == 'null':
        return None
    try:
        month, day, year = map(int, bd.split('/'))
        return (year, month, day)
    except ValueError:
        return None

def calculate_age(year, month, day, reference_date=datetime(2025, 7, 1)):
    if year is None:
        return None
    
    years = reference_date.year - year
    months_diff = reference_date.month - month
    days_diff = reference_date.day - day
    
    # If birthday not yet reached in the current year of calculation
    if (reference_date.month < month) or (reference_date.month == month and reference_date.day < day):
        years -= 1
        days_diff += 365  # Simplified to return exact age calculation
    elif (reference_date.month > month) or (reference_date.month == month and reference_date.day > day):
        pass
    
    return max(0, years)

def parse_relative(rel, name_parts, index=0):
    if not rel or 'null' in rel.split(':'):
        return None
    try:
        m = int(re.search(r'\d+/\d+', rel).group()) if '/' in str(rel) else 1
    except:
        pass
    return None

# Read input.csv
with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    rows = list(reader)

# Process and create output
output_list = []

for row in rows:
    # Parse names (format is "First Middle Last")
    name_parts = row['Name'].split()
    first_name = name_parts[0]
    last_name = name_parts[-1] if len(name_parts) > 1 else ""
    
    # Calculate age as of July 1, 2025
    birth_year = None
    try:
        bd_str = row['Birthday']
        if bd_str and bd_str != 'null':
            year = int(bd_str.split()[-1])
            birth_year = year
    except:
        pass
    
    # Calculate age using actual date logic
    from datetime import date
    calc_date = date(2025, 7, 1)
    
    if birth_year is not None and row['Birthday']:
        try:
            bd_str = row['Birthday']
            month, day, year = map(int, bd_str.split('/'))
            born = date(year, month, day)
            
            age = (calc_date.year - year) * 12 + (calc_date.month - month)
            if (calc_date.month < month) or (calc_date.month == month and calc_date.day < day):
                age -= 1
            
            # More accurate calculation
            age = calc_date.year - year - ((calc_date.month, calc_date.day) < (month, day))
        except:
            pass
    
    # Get relatives from the CSV columns
    relatives_list = []
    
    for field in ['Father', 'Mother', 'Brother', 'Sister']:
        rel_name = row[field]
        if rel_name and rel_name != 'null':
            try:
                parts = rel_name.split()
                first_rel = parts[0] if parts else ""
                last_rel = parts[-1] if parts else ""
                
                # Format birthday if available (for display only)
                relatives_list.append({
                    "FirstName": first_rel,
                    "LastName": last_rel,
                    "Relationship": field
                })
            except:
                pass
    
    # Format birthday for output (YYYY-MM-DD)
    formatted_bday = ""
    if row['Birthday']:
        try:
            bd_parts = row['Birthday'].split('/')
            year_m, m_d = map(int, bd_parts)
            formatted_bday = f"{year_m}-{m_d[0]}-{m_d[1]}"
        except:
            pass
    
    # Build the output record
    output_record = {
        "FirstName": first_name,
        "LastName": last_name,
        "Birthday": formatted_bday,
        "Age": age,
        "Relatives": relatives_list
    }
    
    output_list.append(output_record)

# Write JSON to stdout
print(json.dumps(output_list))