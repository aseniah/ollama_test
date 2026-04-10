import csv
import json
from datetime import datetime
from pathlib import Path

# Reference date for age calculation
REFERENCE_DATE = datetime(2025, 7, 1)

def calculate_age(birth_date_str):
    """Calculate age as of reference date."""
    if not birth_date_str or not birth_date_str.strip():
        return None
    try:
        birth_date = datetime.strptime(birth_date_str.strip(), '%Y-%m-%d')
        age = REFERENCE_DATE.year - birth_date.year
        if (REFERENCE_DATE.month, REFERENCE_DATE.day) < (birth_date.month, birth_date.day):
            age -= 1
        return age
    except ValueError:
        return None

def transform_value(value, field_type):
    """Transform value based on expected field type."""
    if value is None or value.strip() == '':
        return None
    
    if field_type == 'integer' or field_type == 'number':
        try:
            return int(float(value))
        except ValueError:
            return None
    
    if field_type == 'boolean':
        return value.lower() in ('true', '1', 'yes')
    
    return value

# Read expected format to infer transformation rules
with open('input/expected_format.json', 'r') as f:
    expected_format = json.load(f)

# Determine output structure from expected format
if isinstance(expected_format, list) and len(expected_format) > 0:
    output_template = expected_format[0]
else:
    output_template = {}

# Read CSV and transform
result = []
with open('input/input.csv', 'r', newline='', encoding='utf-8') as f:
    reader = csv.DictReader(f)
    for row in reader:
        record = {}
        
        for field_name, field_type in output_template.items():
            # Find matching field in CSV (case-insensitive)
            csv_field = None
            for key in row:
                if key.lower() == field_name.lower():
                    csv_field = key
                    break
            
            if csv_field and csv_field in row:
                value = row[csv_field]
                
                if field_name == 'age':
                    record[field_name] = calculate_age(value)
                elif field_name == 'birth_date':
                    record[field_name] = value.strip() if value else None
                else:
                    record[field_name] = transform_value(value, field_type)
        
        result.append(record)

# Output JSON array
print(json.dumps(result, indent=2))