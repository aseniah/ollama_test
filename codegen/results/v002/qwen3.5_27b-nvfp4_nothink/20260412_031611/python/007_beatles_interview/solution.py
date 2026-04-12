import csv
import json
from datetime import datetime

def calculate_age(birthday_str, reference_date):
    try:
        # Parse the birthday string (MM/DD/YYYY)
        bday = datetime.strptime(birthday_str, "%m/%d/%Y")
        
        # Calculate age as of July 1, 2025
        years = reference_date.year - bday.year
        # Check if birthday has occurred this year yet
        if (reference_date.month, reference_date.day) < (bday.month, bday.day):
            years -= 1
        
        return max(0, years)
    except ValueError:
        return None

def main():
    input_file = 'input/input.csv'
    reference_date = datetime(2025, 7, 1)
    
    output_records = []
    
    with open(input_file, mode='r', encoding='utf-8') as file:
        reader = csv.DictReader(file)
        
        for row in reader:
            # Parse full name into First and Last
            full_name = row['Name'].strip()
            parts = full_name.split()
            
            first_name = parts[0]
            last_name = ' '.join(parts[1:]) if len(parts) > 1 else ""
            
            # Calculate age
            age = calculate_age(row['Birthday'], reference_date)
            
            relatives = []
            
            # Process Father
            father_name = row['Father'].strip() if row['Father'] and row['Father'].lower() != 'null' else None
            if father_name:
                f_parts = father_name.split()
                rel_f_name = f_parts[0]
                rel_l_name = ' '.join(f_parts[1:]) if len(f_parts) > 1 else ""
                relatives.append({
                    "FirstName": rel_f_name,
                    "LastName": rel_l_name,
                    "Relationship": "Father"
                })
            
            # Process Mother
            mother_name = row['Mother'].strip() if row['Mother'] and row['Mother'].lower() != 'null' else None
            if mother_name:
                m_parts = mother_name.split()
                rel_f_name = m_parts[0]
                rel_l_name = ' '.join(m_parts[1:]) if len(m_parts) > 1 else ""
                relatives.append({
                    "FirstName": rel_f_name,
                    "LastName": rel_l_name,
                    "Relationship": "Mother"
                })
            
            # Process Brother
            brother_name = row['Brother'].strip() if row['Brother'] and row['Brother'].lower() != 'null' else None
            if brother_name:
                b_parts = brother_name.split()
                rel_f_name = b_parts[0]
                rel_l_name = ' '.join(b_parts[1:]) if len(b_parts) > 1 else ""
                relatives.append({
                    "FirstName": rel_f_name,
                    "LastName": rel_l_name,
                    "Relationship": "Brother"
                })
            
            # Process Sister
            sister_name = row['Sister'].strip() if row['Sister'] and row['Sister'].lower() != 'null' else None
            if sister_name:
                s_parts = sister_name.split()
                rel_f_name = s_parts[0]
                rel_l_name = ' '.join(s_parts[1:]) if len(s_parts) > 1 else ""
                relatives.append({
                    "FirstName": rel_f_name,
                    "LastName": rel_l_name,
                    "Relationship": "Sister"
                })
            
            record = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": datetime.strptime(row['Birthday'], "%m/%d/%Y").strftime("%Y-%m-%d"),
                "Age": age,
                "Relatives": relatives
            }
            
            output_records.append(record)
    
    print(json.dumps(output_records, indent=2))

if __name__ == "__main__":
    main()