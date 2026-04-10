import csv
import json
import sys
from datetime import datetime

def calculate_age(birth_date_str, target_date):
    """Calculates age as of a specific target date."""
    try:
        birth_date = datetime.strptime(birth_date_str, '%Y-%m-%d')
        age = target_date.year - birth_date.year - (
            (target_date.month, target_date.day) < (birth_date.month, birth_date.day)
        )
        return age
    except ValueError:
        return None

def transform_data(input_path, target_date):
    """
    Reads CSV and transforms it based on inferred logic.
    Since the expected_format.json is not provided, 
    this implementation assumes a standard mapping:
    CSV columns: 'name', 'birth_date', 'city'
    JSON keys: 'full_name', 'age', 'location'
    """
    results = []
    
    try:
        with open(input_path, mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # 1. Map name to full_name (example transformation)
                full_name = row.get('name', '').strip()
                
                # 2. Calculate age as of July 1, 2025
                age = calculate_age(row.get('birth_date', ''), target_date)
                
                # 3. Map city to location
                location = row.get('city', '').strip()
                
                # Construct the transformed object
                transformed_entry = {
                    "full_name": full_name,
                    "age": age,
                    "location": location
                }
                results.append(transformed_entry)
                
    except FileNotFoundError:
        # If file doesn't exist, we output nothing or an empty array
        return []
    except Exception:
        return []

    return results

def main():
    input_csv = 'input/input.csv'
    # Reference date provided in prompt: July 1, 2025
    reference_date = datetime(2025, 7, 1)
    
    output_data = transform_data(input_csv, reference_date)
    
    # Output only the JSON array to stdout
    sys.stdout.write(json.dumps(output_data, indent=2))
    sys.stdout.write('\n')

if __name__ == "__main__":
    main()