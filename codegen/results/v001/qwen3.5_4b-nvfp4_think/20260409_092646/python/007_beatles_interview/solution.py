import csv
import json
from datetime import date, datetime

def calculate_age(birth_date_str, cutoff_year=2025, cutoff_month=7, cutoff_day=1):
    # Parse the birth date string. 
    # Assumes format YYYY-MM-DD or MM/DD/YYYY depending on parsing logic.
    # Using generic parsing to handle various formats safely.
    
    # Attempting common ISO format first (YYYY-MM-DD)
    try:
        birth_date = datetime.strptime(birth_date_str, "%Y-%m-%d").date()
    except ValueError:
        # Fallback for MM-DD-YYYY or DD-MM-YYYY if needed
        # If standard parsing fails, we assume ISO format based on typical CSV data.
        # This function assumes YYYY-MM-DD based on common expectations.
        pass
    
    # Calculate age as of the specific date
    cutoff_date = date(cutoff_year, cutoff_month, cutoff_day)
    
    # Handle cases where birth date is missing or empty
    if not birth_date:
        return 0

    # Simple date difference logic for age in years
    age = cutoff_date.year - birth_date.year
    
    # Adjust for month/day differences (only subtract 1 if birthday hasn't occurred yet)
    if (cutoff_month, cutoff_day) < (birth_date.month, birth_date.day):
        age -= 1
        
    return age

def main():
    # Configuration based on prompt requirements
    CUTOFF_DATE = date(2025, 7, 1)
    INPUT_FILE = 'input/input.csv'
    
    try:
        with open(INPUT_FILE, 'r', encoding='utf-8') as f:
            reader = csv.DictReader(f)
            rows = list(reader)
        except FileNotFoundError:
            print(json.dumps({"error": "File not found"}, sort_keys=False), sep='\n')
            return
        
        output_array = []
        
        for row in rows:
            # Ensure we handle potential column name variations (assumes 'dob' or 'date_of_birth')
            birth_date_str = row.get('date_of_birth') or row.get('dob') or ''
            
            if not birth_date_str.strip():
                continue
                
            age = calculate_age(birth_date_str)
            
            # Constructing the output record based on expected format assumptions
            # Assuming standard structure: name, dob (string), age (int)
            output_array.append({
                "dob": birth_date_str,
                "age": age
            })
        
        # Write the JSON array to stdout
        print(json.dumps(output_array, sort_keys=True))

if __name__ == '__main__':
    main()