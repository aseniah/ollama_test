import csv
import json
from datetime import datetime

def calculate_age(birthday_str, reference_date):
    # Birthday format in CSV is M/D/YYYY
    birth_date = datetime.strptime(birthday_str, '%m/%d/%Y')
    
    # If person died, age is calculated as of death date
    # However, the prompt specifically asks for ages as of July 1, 2025.
    # Looking at expected_format.json, John Lennon (died 1980) has age 40.
    # 1980-12-08 minus 1940-10-09 is 40. 
    # This implies: If Died date exists, age = Died Date - Birthday. 
    # Otherwise, age = Reference Date - Birthday.
    return birth_date

def main():
    reference_date = datetime(2025, 7, 1)
    results = []

    try:
        with open('input/input.csv', mode='r', encoding='utf-8') as f:
            reader = csv.DictReader(f)
            for row in reader:
                # Split full name into First and Last
                full_name_parts = row['Name'].split(' ')
                first_name = full_name_parts[0]
                last_name = full_name_parts[-1]
                
                # Parse Birthday
                bday_dt = datetime.strptime(row['Birthday'], '%m/%d/%Y')
                birthday_iso = bday_dt.strftime('%Y-%m-%d')
                
                # Determine Age
                if row['Died'] and row['Died'].lower() != 'null':
                    death_dt = datetime.strptime(row['Died'], '%m/%d/%Y')
                    # Calculate age at time of death
                    age = death_dt.year - bday_dt.year - ((death_dt.month, death_dt.day) < (bday_dt.month, bday_dt.day))
                else:
                    # Calculate age as of July 1, 2025
                    age = reference_date.year - bday_dt.year - ((reference_date.month, reference_date.day) < (bday_dt.month, bday_dt.day))
                
                # Process Relatives
                relatives = []
                rel_cols = {
                    'Father': 'Father',
                    'Mother': 'Mother',
                    'Brother': 'Brother',
                    'Sister': 'Sister'
                }
                
                for col, relationship in rel_cols.items():
                    val = row[col]
                    if val and val.lower() != 'null':
                        rel_name_parts = val.split(' ')
                        relatives.append({
                            "FirstName": rel_name_parts[0],
                            "LastName": rel_name_parts[-1],
                            "Relationship": relationship
                        })
                
                results.append({
                    "FirstName": first_name,
                    "LastName": last_name,
                    "Birthday": birthday_iso,
                    "Age": age,
                    "Relatives": relatives
                })
                
        print(json.dumps(results, indent=2))
    except FileNotFoundError:
        pass

if __name__ == "__main__":
    main()