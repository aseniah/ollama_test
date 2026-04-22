import csv
import json
from datetime import date

def main():
    ref_date = date(2025, 7, 1)
    result = []
    
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        for row in reader:
            name_parts = row['Name'].strip().split()
            first_name = name_parts[0]
            last_name = name_parts[-1]
            
            b_parts = row['Birthday'].strip().split('/')
            birthday_str = f"{b_parts[2]}-{b_parts[0].zfill(2)}-{b_parts[1].zfill(2)}"
            
            b_month, b_day, b_year = map(int, b_parts)
            b_date = date(b_year, b_month, b_day)
            age = ref_date.year - b_date.year
            if (ref_date.month, ref_date.day) < (b_date.month, b_date.day):
                age -= 1
                
            relatives = []
            for col, rel in [('Father', 'Father'), ('Mother', 'Mother'), ('Brother', 'Brother'), ('Sister', 'Sister')]:
                val = row[col].strip()
                if val.lower() != 'null':
                    parts = val.split()
                    relatives.append({
                        'FirstName': parts[0],
                        'LastName': parts[-1],
                        'Relationship': rel
                    })
                    
            result.append({
                'FirstName': first_name,
                'LastName': last_name,
                'Birthday': birthday_str,
                'Age': age,
                'Relatives': relatives
            })
            
    print(json.dumps(result, indent=2))

if __name__ == '__main__':
    main()