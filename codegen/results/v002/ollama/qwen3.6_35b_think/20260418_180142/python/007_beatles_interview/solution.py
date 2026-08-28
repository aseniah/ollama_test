import csv
import json

def main():
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        result = []
        ref_year, ref_month, ref_day = 2025, 7, 1
        
        for row in reader:
            # Extract first and last name
            name_parts = row['Name'].strip().split()
            first_name = name_parts[0]
            last_name = name_parts[-1]
            
            # Parse and format birthday
            b_parts = row['Birthday'].strip().split('/')
            b_year, b_month, b_day = int(b_parts[2]), int(b_parts[0]), int(b_parts[1])
            birthday_str = f"{b_year:04d}-{b_month:02d}-{b_day:02d}"
            
            # Calculate age as of July 1, 2025
            age = ref_year - b_year
            if (ref_month, ref_day) < (b_month, b_day):
                age -= 1
                
            # Build relatives list
            relatives = []
            for rel_type in ['Father', 'Mother', 'Brother', 'Sister']:
                rel_val = row.get(rel_type, '').strip()
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

if __name__ == "__main__":
    main()