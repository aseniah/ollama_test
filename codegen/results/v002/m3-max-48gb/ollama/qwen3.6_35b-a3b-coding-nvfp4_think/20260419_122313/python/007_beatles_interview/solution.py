import csv
import json
from datetime import date

def main():
    result = []
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        for row in reader:
            # Clean keys and values
            row = {k.strip(): v.strip() for k, v in row.items()}
            
            # Parse Name into First and Last
            name_parts = row['Name'].split()
            first_name = name_parts[0]
            last_name = name_parts[-1]
            
            # Parse Birthday to YYYY-MM-DD format
            b_parts = row['Birthday'].split('/')
            birthday_str = f"{int(b_parts[2]):04d}-{int(b_parts[0]):02d}-{int(b_parts[1]):02d}"
            
            # Calculate Age as of July 1, 2025
            birth_date = date(int(b_parts[2]), int(b_parts[0]), int(b_parts[1]))
            ref_date = date(2025, 7, 1)
            age = ref_date.year - birth_date.year
            if (ref_date.month, ref_date.day) < (birth_date.month, birth_date.day):
                age -= 1
                
            # Parse Relatives
            relatives = []
            for rel_type in ['Father', 'Mother', 'Brother', 'Sister']:
                rel_val = row.get(rel_type, '')
                if rel_val and rel_val.lower() != 'null':
                    rel_parts = rel_val.split()
                    relatives.append({
                        "FirstName": rel_parts[0],
                        "LastName": rel_parts[-1],
                        "Relationship": rel_type
                    })
                    
            result.append({
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday_str,
                "Age": age,
                "Relatives": relatives
            })
            
    print(json.dumps(result, indent=2))

if __name__ == '__main__':
    main()