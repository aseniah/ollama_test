import csv
import json
from datetime import date
import os

def main():
    input_file = "input/input.csv"
    expected_file = "input/expected_format.json"
    
    # Determine reference date for age calculation
    reference_date = date(2025, 7, 1)
    
    output_data = []
    
    if not os.path.exists(input_file):
        print(json.dumps(output_data))
        return

    with open(input_file, mode='r', encoding='utf-8') as file:
        reader = csv.DictReader(file)
        
        # Try to infer key mapping from the first row or expected format
        # If we can read the expected file, we use it to map columns strictly
        headers_in_expected = []
        if os.path.exists(expected_file):
            try:
                with open(expected_file, mode='r', encoding='utf-8') as exp_f:
                    sample = json.load(exp_f)
                    if isinstance(sample, list) and len(sample) > 0:
                        headers_in_expected = list(sample[0].keys())
            except Exception:
                pass
        
        # Common heuristics for mapping CSV columns to the expected keys
        # We assume the CSV has common names like 'name', 'birthdate' or 'dob', 'id', etc.
        # If the expected format exists, we try to match by similarity or standard conventions.
        
        # Fallback headers if no sample provided
        target_keys = ['id', 'first_name', 'last_name', 'age'] 
        if headers_in_expected:
            target_keys = headers_in_expected
        
        for row in reader:
            record = {}
            
            # Map 'id'
            if 'id' in target_keys:
                possible_ids = ['id', 'ID', 'employee_id', 'user_id']
                val = next((row[k] for k in possible_ids if k in row), None)
                if val:
                    record['id'] = int(val) if val.isdigit() else val
            
            # Map Names (Heuristic: assume single name string or split by space if no separate columns)
            if 'first_name' in target_keys and 'last_name' in target_keys:
                possible_firsts = ['first_name', 'FirstName', 'fname']
                possible_lasts = ['last_name', 'LastName', 'lname']
                
                fname_val = next((row[k] for k in possible_firsts if k in row), None)
                lname_val = next((row[k] for k in possible_lasts if k in row), None)
                
                # If only one name column exists (e.g., 'full_name'), split it
                if not fname_val and not lname_val:
                    possible_full = ['name', 'full_name', 'Name']
                    full_name = next((row[k] for k in possible_full if k in row), None)
                    if full_name:
                        parts = full_name.strip().split()
                        fname_val = parts[0] if len(parts) > 0 else ""
                        lname_val = " ".join(parts[1:]) if len(parts) > 1 else ""
                
                record['first_name'] = fname_val or ""
                record['last_name'] = lname_val or ""
            elif 'name' in target_keys and ('first_name' not in target_keys or 'last_name' not in target_keys):
                 possible_full = ['name', 'full_name', 'Name', 'first_name'] # catch-all
                 val = next((row[k] for k in possible_full if k in row), None)
                 record['name'] = val or ""

            # Map Date of Birth and calculate Age
            date_keys = []
            # Define potential keys found in CSV that might be a date
            possible_date_cols = ['birthdate', 'dob', 'date_of_birth', 'birth_date', 'BirthDate', 'DOB', 'date']
            
            birth_str = None
            for col in possible_date_cols:
                if col in row and row[col]:
                    birth_str = row[col]
                    break
            
            if birth_str:
                try:
                    # Try parsing ISO format YYYY-MM-DD first, then common US/Intl formats
                    # Common formats: YYYY-MM-DD, DD/MM/YYYY, MM/DD/YYYY
                    birth_date_obj = None
                    
                    # Attempt 1: YYYY-MM-DD
                    if '-' in birth_str and '/' not in birth_str:
                        parts = birth_str.split('-')
                        if len(parts) == 3:
                            birth_date_obj = date(int(parts[0]), int(parts[1]), int(parts[2]))
                    
                    # Attempt 2: DD/MM/YYYY or MM/DD/YYYY (Ambiguous without context, assume ISO if parsed above, else try specific)
                    elif '/' in birth_str:
                        parts = birth_str.split('/')
                        if len(parts) == 3:
                            d, m, y = int(parts[0]), int(parts[1]), int(parts[2])
                            # Simple heuristic: if first part > 12, it must be DD/MM/YYYY
                            if d > 12:
                                birth_date_obj = date(y, m, d)
                            else:
                                # Assume MM/DD/YYYY as default US format if ambiguous, 
                                # or try to see if M/D/Y is valid.
                                # Since no context, we try MM/DD/YYYY first for US standard.
                                try:
                                    birth_date_obj = date(y, m, d) # Assuming MM/DD/YYYY logic failed? 
                                    # Actually standard Python date parsing is robust. 
                                    # Let's try to guess based on typical dataset constraints (Year is usually 4 digits).
                                    # If parts[2] < 100, it's likely YYYY/MM/DD or similar.
                                    if y < 100:
                                        # Assume YYYY-MM-DD split logic handled above? No, we are in / branch.
                                        # If y is small (e.g. 25), it might be the year? Or day? 
                                        # Let's try: if len(parts[2]) == 4 -> YYYY at end.
                                        if len(parts[2]) == 4:
                                            birth_date_obj = date(int(parts[2]), int(parts[1]), int(parts[0])) # DD/MM/YYYY
                                        else:
                                            birth_date_obj = date(y, m, d) # MM/DD/YYYY
                                    else:
                                        birth_date_obj = date(y, m, d)
                                except ValueError:
                                    try:
                                         # Fallback to YYYY-MM-DD logic if year was at start but delimiters were wrong? Unlikely.
                                         pass 
                                    except:
                                        pass
                    
                    if birth_date_obj:
                        # Calculate age
                        age = reference_date.year - birth_date_obj.year
                        if (reference_date.month, reference_date.day) < (birth_date_obj.month, birth_date_obj.day):
                            age -= 1
                        
                        record['age'] = age
                
                except (ValueError, IndexError):
                    pass # Skip if date parsing fails
            
            if record:
                output_data.append(record)

    print(json.dumps(output_data))

if __name__ == "__main__":
    main()