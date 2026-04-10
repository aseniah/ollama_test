import csv
import json
from datetime import datetime

# Constants
CUTOFF_DATE = datetime(2025, 7, 1)
INPUT_FILE = 'input/input.csv'
EXPECTED_FORMAT_FILE = 'input/expected_format.json'

def calculate_age(birth_date_str):
    """
    Calculates age as of CUTOFF_DATE.
    Assumes input birth_date_str is in YYYY-MM-DD format.
    """
    try:
        birth_date = datetime.strptime(birth_date_str, '%Y-%m-%d')
        age = CUTOFF_DATE.year - birth_date.year
        # If birthday hasn't occurred yet in the current year, subtract 1
        if (CUTOFF_DATE.month, CUTOFF_DATE.day) < (birth_date.month, birth_date.day):
            age -= 1
        return age
    except ValueError:
        return None

def load_expected_format():
    """
    Loads the expected format file to infer structure (field names, types).
    Returns the list of rows from the JSON file to analyze keys.
    """
    try:
        with open(EXPECTED_FORMAT_FILE, 'r') as f:
            return json.load(f)
    except FileNotFoundError:
        # Fallback if file doesn't exist, though prompt implies it does.
        return []

def main():
    output_data = []
    
    # Load expected format to get keys and structure if needed
    # In a real scenario with unknown mapping, we'd inspect this to map CSV columns to JSON keys.
    # For this task, we assume the CSV headers might match or we map standard fields.
    expected = load_expected_format()
    
    # Read CSV
    try:
        with open(INPUT_FILE, 'r', newline='', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            
            for row in reader:
                record = {}
                
                # Extract and transform fields
                # We infer keys from the CSV headers or the expected format.
                # Assuming the CSV has headers like 'id', 'name', 'date_of_birth', etc.
                
                # 1. ID
                # Try common ID field names
                id_val = None
                for key in ['id', 'ID', 'employee_id', 'EmployeeID', 'user_id']:
                    if key in row and row[key]:
                        id_val = row[key]
                        break
                # If expected format shows a specific key, we prioritize that mapping logic if possible.
                # Since we don't know exact CSV headers without seeing the file, we rely on common conventions 
                # or map directly if headers match expected keys.
                
                # Let's try to map based on the first entry of expected if available
                if expected and len(expected) > 0:
                    # Get keys from expected format
                    target_keys = expected[0].keys()
                    
                    # Map CSV 'id' to JSON 'id' (or whatever the expected key is)
                    # This is a heuristic: if expected has 'id', look for it in row.
                    for target_key in target_keys:
                        # Heuristic: if target_key is 'age', calculate it.
                        if target_key == 'age':
                            # Find DOB column
                            dob_val = None
                            for csv_key in row:
                                if 'birth' in csv_key.lower() or 'dob' in csv_key.lower():
                                    dob_val = row[csv_key]
                                    break
                            if dob_val:
                                record[target_key] = calculate_age(dob_val)
                            continue
                        
                        # If target_key is 'name' or 'first_name' or 'last_name'
                        if 'name' in target_key.lower():
                            for csv_key in row:
                                if 'name' in csv_key.lower():
                                    if target_key.lower() == 'name':
                                        record[target_key] = row[csv_key]
                                        break
                                    # Handle split names if necessary
                            continue
                        
                        # Default: direct mapping if header matches, otherwise try fuzzy match
                        found = False
                        for csv_key, csv_val in row.items():
                            if csv_key.lower() == target_key.lower():
                                record[target_key] = csv_val
                                found = True
                                break
                        
                        if not found:
                            # Try to find a column that matches the target key name loosely
                            for csv_key, csv_val in row.items():
                                if csv_key.lower().replace('_', '').replace(' ', '') == target_key.lower().replace('_', '').replace(' ', ''):
                                    record[target_key] = csv_val
                                    break

                else:
                    # If no expected format, just dump the row with age calculation
                    # But we must produce JSON array, let's just pass through keys
                    for key, value in row.items():
                        if 'birth' in key.lower() or 'dob' in key.lower():
                            record['age'] = calculate_age(value)
                        else:
                            record[key] = value

                # If we used the expected format logic, we might have missed the direct copy if keys matched exactly
                # Let's refine: If expected format exists, we strictly follow its keys.
                # If not, we just output the row data + age.
                
                if not expected:
                    # Fallback: simple passthrough + age
                    final_record = {}
                    for k, v in row.items():
                        if 'birth' in k.lower() or 'dob' in k.lower():
                            # Don't include raw DOB in output if not in expected format?
                            # Prompt says "infer transformation rules". Usually means output matches expected keys.
                            pass 
                        else:
                            final_record[k] = v
                    # Add age if we found a DOB
                    for k, v in row.items():
                        if 'birth' in k.lower() or 'dob' in k.lower():
                            final_record['age'] = calculate_age(v)
                            break
                    output_data.append(final_record)
                else:
                    output_data.append(record)

    except FileNotFoundError:
        print("[]")
        return

    # Ensure output is a JSON array
    print(json.dumps(output_data))

if __name__ == "__main__":
    main()