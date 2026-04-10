import pandas as pd
import json
from datetime import date
import sys

# Define reference date for age calculation
reference_date = date(2025, 7, 1)

try:
    # Read the CSV file
    df = pd.read_csv('input/input.csv')

    # Read the expected format to infer structure if specific columns aren't obvious
    # However, standard practice is to map common names. 
    # We will assume the input has at least 'name' and 'birth_date' or similar.
    # Let's try to detect date columns based on content if headers are generic, 
    # but usually expected_format.json helps confirm keys.
    
    with open('input/expected_format.json', 'r') as f:
        expected_sample = json.load(f)
    
    # The task says "infer transformation rules" from expected_format.json.
    # Usually this means the JSON structure (keys) and data types.
    # Let's assume the input CSV has columns matching the keys in expected_format.json,
    # or we need to map 'Name' -> 'name', 'Date of Birth' -> 'dob', etc.
    
    output_list = []

    # Helper to parse date
    def safe_date_parse(date_str):
        if pd.isna(date_str):
            return None
        try:
            # Try parsing various formats common in CSVs
            for fmt in ['%Y-%m-%d', '%d/%m/%Y', '%m/%d/%Y', '%Y/%m/%d']:
                try:
                    return pd.to_datetime(date_str, format=fmt).date()
                except ValueError:
                    continue
            # Fallback to pandas default parsing if explicit formats fail
            return pd.to_datetime(date_str).date()
        except Exception:
            return None

    # Iterate through rows and construct the output based on expected_format structure
    # We assume the keys in the first element of expected_sample are the target keys.
    # We map common source column names to these target keys if necessary.
    
    target_keys = list(expected_sample.keys()) if isinstance(expected_sample, dict) else []
    
    # Heuristic mapping for common column names to target keys
    # This logic attempts to match CSV headers to the expected JSON keys by name similarity
    csv_columns = df.columns.tolist()
    key_mapping = {}

    for target_key in target_keys:
        # Direct match
        if target_key in csv_columns:
            key_mapping[target_key] = target_key
        else:
            # Fuzzy/Case-insensitive match or common aliases
            matched_source = None
            for src_col in csv_columns:
                # Check if src_col is a variation of target_key (e.g., "Date of Birth" -> "birth_date")
                normalized_target = target_key.replace("_", " ").lower()
                normalized_src = src_col.lower()
                
                if target_key == "name":
                    if "name" in normalized_src or "full_name" in normalized_src:
                        matched_source = src_col
                        break
                elif target_key == "age":
                    # Age is calculated, usually mapped from a birth_date column
                    pass 
                elif target_key == "birth_date" or "dob" in target_key.lower():
                    if "birth" in normalized_src or "dob" in normalized_src:
                        matched_source = src_col
                        break
                else:
                    # Simple containment check
                    if target_key.replace("_", "") in "".join(normalized_src.split()):
                         matched_source = src_col
                         break
            
            key_mapping[target_key] = matched_source

    for idx, row in df.iterrows():
        record = {}
        
        for target_key in target_keys:
            source_col = key_mapping.get(target_key)
            
            if target_key == "age":
                # Age needs a birth date source to calculate
                # Identify which column holds the birth date if not explicitly mapped as 'birth_date'
                dob_val = None
                for src_col in csv_columns:
                    if "birth" in src_col.lower() or "dob" in src_col.lower():
                        val = row.get(src_col)
                        parsed = safe_date_parse(val)
                        if parsed:
                            dob_val = parsed
                            break
                
                if dob_val:
                    # Calculate age as of July 1, 2025
                    # Age = Reference Year - Birth Year - (1 if birthday hasn't occurred yet this year else 0)
                    years = reference_date.year - dob_val.year
                    months_diff = reference_date.month - dob_val.month
                    days_diff = reference_date.day - dob_val.day
                    
                    if months_diff < 0 or (months_diff == 0 and days_diff < 0):
                        years -= 1
                    record[target_key] = years
                else:
                    record[target_key] = None # Or handle as per expected format logic
            
            elif source_col:
                val = row[source_col]
                if pd.isna(val):
                    val = None
                # Convert types if necessary based on expected_format.json type
                # But since we output JSON, strings and numbers are handled automatically by json.dumps
                record[target_key] = val
        
        output_list.append(record)

    print(json.dumps(output_list))

except FileNotFoundError as e:
    sys.stderr.write(f"Error: {e}\n")
    sys.exit(1)
except Exception as e:
    sys.stderr.write(f"Error processing file: {e}\n")
    sys.exit(1)