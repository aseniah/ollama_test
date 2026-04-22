import csv
import json
from datetime import datetime, date

def main():
    ref_date = date(2025, 7, 1)
    result = []
    
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        for row in reader:
            name_parts = row['Name'].strip().split()
            first_name = name_parts[0]
            last_name = name_parts[-1]
            
            bday_date = datetime.strptime(row['Birthday'].strip(), "%m/%d/%Y").date()
            bday_str = bday_date.strftime('%Y-%m-%d')
            
            age = ref_date.year - bday_date.year
            if (ref_date.month, ref_date.day) < (bday_date.month, bday_date.day):
                age -= 1
                
            relatives = []
            for rel_type in ['Father', 'Mother', 'Brother', 'Sister']:
                rel_name = row[rel_type].strip()
                if rel_name.lower() != 'null':
                    rel_parts = rel_name.split()
                    relatives.append({
                        "FirstName": rel_parts[0],
                        "LastName": rel_parts[-1],
                        "Relationship": rel_type
                    })
                    
            result.append({
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": bday_str,
                "Age": age,
                "Relatives": relatives
            })
            
    print(json.dumps(result, indent=2))

if __name__ == '__main__':
    main()