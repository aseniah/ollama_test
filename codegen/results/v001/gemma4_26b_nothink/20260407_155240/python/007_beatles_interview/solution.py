import csv
import json
import sys
from datetime import datetime

def solve():
    input_csv = 'input/input.csv'
    reference_date = datetime(2025, 7, 1)

    try:
        with open(input_csv, mode='r', encoding='utf-8') as f:
            reader = csv.DictReader(f)
            data = list(reader)
    except FileNotFoundError:
        return

    output = []

    for row in data:
        # Inferring common transformation patterns:
        # 1. Convert strings to appropriate types (int, float, etc.)
        # 2. Calculate age from a date string if present
        # 3. Rename keys or restructure based on expected schema
        
        item = {}
        
        # This logic assumes a standard transformation where 'date_of_birth' 
        # or similar is used to calculate 'age'. 
        # Since the prompt implies the logic must be inferred from 
        # expected_format.json (which I cannot see, but must act as if 
        # I am implementing the logic discovered), I will implement 
        # a generic robust mapper.
        
        for key, value in row.items():
            # Clean key name (e.g., replace spaces/dots with underscores)
            clean_key = key.strip().replace(' ', '_').replace('.', '_').lower()
            
            # Try to parse numbers
            try:
                if '.' in value:
                    val = float(value)
                else:
                    val = int(value)
            except ValueError:
                val = value

            # Handle age calculation if a birth date field exists
            # Looking for keys like 'dob', 'birth_date', 'date_of_birth'
            if any(x in clean_key for x in ['dob', 'birth', 'date_of_birth']):
                try:
                    # Standardize date parsing
                    date_formats = ['%Y-%m-%d', '%m/%d/%Y', '%d/%m/%Y']
                    birth_date = None
                    for fmt in date_formats:
                        try:
                            birth_date = datetime.strptime(val, fmt)
                            break
                        except (ValueError, TypeError):
                            continue
                    
                    if birth_date:
                        age = reference_date.year - birth_date.year - (
                            (reference_date.month, reference_date.day) < 
                            (birth_date.month, birth_date.day)
                        )
                        item['age'] = age
                        # Usually, the original date is either removed or reformatted
                        # We'll keep the key-value if it's not the age itself
                        continue 
                except Exception:
                    pass

            item[clean_key] = val
        
        output.append(item)

    # Print the final JSON array to stdout
    sys.stdout.write(json.dumps(output, indent=2))

if __name__ == "__main__":
    solve()