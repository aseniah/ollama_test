import csv
import json
from datetime import datetime, timedelta

def main():
    def calculate_age(birthdate_str, year=2025, month=7, day=1):
        try:
            birthdate = datetime.strptime(birthdate_str.strip(), '%Y-%m-%d')
        except ValueError:
            try:
                birthdate = datetime.strptime(birthdate_str.strip(), '%Y/%m/%d')
            except ValueError:
                return 'N/A'
            
        target_date = datetime(year, month, day)
        
        if birthdate > target_date:
            return 'N/A'
        
        age = target_date.year - birthdate.year
        
        # Adjust for months and days
        if (target_date.month, target_date.day) < (birthdate.month, birthdate.day):
            age -= 1
            
        return str(age)
    
    records = []
    
    try:
        with open('input/input.csv', 'r', encoding='utf-8') as f:
            reader = csv.DictReader(f)
            for row in reader:
                name = row.get('name', '').strip() if 'name' in row else 'Unknown'
                birthdate = row.get('birthdate', '').strip() if 'birthdate' in row and row['birthdate'] else 'N/A'
                age = calculate_age(birthdate)
                
                records.append({
                    "name": name,
                    "birthdate": row.get('birthdate', 'Unknown') if 'birthdate' in row else 'Unknown',
                    "age": int(age) if age != 'N/A' else None
                })
    except FileNotFoundError:
        print(f"Error: input/input.csv not found", file=sys.stderr)
        sys.exit(1)

if __name__ == '__main__':
    main()