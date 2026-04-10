import csv
import json
import sys
from datetime import datetime

def calculate_age(birth_date_str, reference_date):
    """Calculates age given a birth date string and a reference date."""
    # Common date formats to attempt parsing
    formats = ['%Y-%m-%d', '%m/%d/%Y', '%d/%m/%Y', '%Y/%m/%d']
    dob = None
    
    for fmt in formats:
        try:
            dob = datetime.strptime(birth_date_str, fmt)
            break
        except ValueError:
            continue
            
    if not dob:
        return None
        
    age = reference_date.year - dob.year - ((reference_date.month, reference_date.day) < (dob.month, dob.day))
    return age

def main():
    reference_date = datetime(2025, 7, 1)
    csv_path = 'input/input.csv'
    json_ref_path = 'input/expected_format.json'
    
    try:
        # 1. Load the expected format to understand the target schema
        with open(json_ref_path, 'r', encoding='='utf-8') as f:
            expected_data = json.load(f)
        
        # Use the first object in the array to determine required keys
        target_keys = []
        if isinstance(expected_data, list) and len(expected_data) > 0:
            target_keys = list(expected_data[0].keys())
        elif isinstance(expected_data, dict):
            target_keys = list(expected_data.keys())
            
        # 2. Read the CSV data
        output_array = []
        with open(csv_path, mode='r', encoding='utf-8-sig') as f:
            reader = csv.DictReader(f)
            # Standardize CSV headers: lowercase and replace spaces with underscores
            csv_headers = {h.strip().lower().replace(' ', '_'): h for h in reader.fieldnames}
            
            for row in reader:
                transformed_row = {}
                
                for key in target_keys:
                    # Normalize target key for matching
                    normalized_key = key.lower().replace(' ', '_')
                    
                    # Logic for 'age' field
                    if normalized_key == 'age':
                        # Look for potential date columns in CSV
                        date_col_key = None
                        possible_date_matches = ['birth_date', 'dob', 'date_of_birth', 'birthday', 'birthdate']
                        
                        # Check if any of our known date columns exist in the CSV
                        for match in possible_date_matches:
                            if match in csv_headers:
                                date_col_key = csv_headers[match]
                                break
                        
                        if date_col_key and row.get(date_col_key):
                            age = calculate_age(row[date_col_key], reference_date)
                            transformed_row[key] = age
                        else:
                            transformed_row[key] = None
                    
                    # Logic for standard fields
                    else:
                        # Try direct match, then normalized match
                        if normalized_key in csv_headers:
                            original_col_name = csv_headers[normalized_key]
                            val = row.get(original_col_name)
                            # Attempt to convert numeric strings to int/float if needed
                            if val is not None:
                                try:
                                    if '.' in val:
                                        val = float(val)
                                    else:
                                        val = int(val)
                                except ValueError:
                                    pass
                                transformed_row[key] = val
                        else:
                            transformed_row[key] = None
                            
                output_array.append(transformed_row)

        # 3. Output only the JSON array to stdout
        sys.stdout.write(json.dumps(output_array, indent=2))

    except FileNotFoundError as e:
        # If files are missing, we cannot produce the JSON array as requested.
        # However, per instructions, we shouldn't output anything else.
        pass
    except Exception:
        pass

if __name__ == "__main__":
    main()