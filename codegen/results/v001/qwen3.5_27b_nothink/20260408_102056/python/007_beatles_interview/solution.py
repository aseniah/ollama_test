import csv
import json
import datetime
import sys

def calculate_age(birth_date_str, reference_date):
    """
    Calculates age as of a specific reference date.
    Handles formats: 'YYYY-MM-DD' or 'DD/MM/YYYY' or 'MM/DD/YYYY'.
    """
    try:
        # Try common date formats
        for fmt in ['%Y-%m-%d', '%d/%m/%Y', '%m/%d/%Y']:
            try:
                birth_date = datetime.datetime.strptime(birth_date_str, fmt).date()
                break
            except ValueError:
                continue
        else:
            # If format fails, return null or handle error
            return None

        # Calculate age
        age = reference_date.year - birth_date.year
        
        # Check if birthday has occurred this year
        if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
            age -= 1
            
        return age
    except Exception:
        return None

def main():
    input_file = 'input/input.csv'
    reference_date = datetime.date(2025, 7, 1)
    
    output_data = []

    try:
        with open(input_file, mode='r', encoding='utf-8') as f:
            # Use DictReader to handle headers automatically
            reader = csv.DictReader(f)
            
            # Infer column mappings based on common CSV structures 
            # since expected_format.json is not readable here.
            # We assume columns like: 'name', 'email', 'birth_date', etc.
            
            for row in reader:
                record = {}
                
                # Map specific fields, assuming standard naming or passing through all
                # Adjust these keys to match your specific CSV headers exactly
                for key, value in row.items():
                    key_lower = key.strip().lower().replace(' ', '_')
                    
                    if 'birth' in key_lower or 'dob' in key_lower or 'date_of_birth' in key_lower:
                        record['age'] = calculate_age(value.strip(), reference_date)
                    else:
                        # Keep the original key name (or normalize it)
                        # If you need specific key names from expected_format.json, map them here.
                        # For now, we preserve the key but normalize underscores/spaces for safety
                        record[key_lower] = value.strip()

                output_data.append(record)

        print(json.dumps(output_data, indent=2))

    except FileNotFoundError:
        print(f"Error: {input_file} not found.", file=sys.stderr)
        sys.exit(1)
    except Exception as e:
        print(f"Error processing file: {e}", file=sys.stderr)
        sys.exit(1)

if __name__ == "__main__":
    main()