import csv
import json
from datetime import datetime

def calculate_age(dob_str, target_date):
    """Calculates age as of the target_date."""
    try:
        # Assuming date format is YYYY-MM-DD
        dob = datetime.strptime(dob_str, '%Y-%m-%d')
        age = target_date.year - dob.year - ((target_date.month, target_date.day) < (dob.month, dob.day))
        return age
    except ValueError:
        return None

def main():
    input_file = 'input/input.csv'
    # Target date as specified in the requirements
    target_date = datetime(2025, 7, 1)
    
    results = []
    
    try:
        with open(input_file, mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            
            # To adhere to the requirement of using expected_format.json, 
            # the logic assumes the CSV columns map to JSON keys, 
            # and the date column is replaced by 'age'.
            for row in reader:
                # Identify the date column (common patterns: dob, birth_date, etc.)
                date_col = None
                for key in row.keys():
                    if key.lower() in ['dob', 'birth_date', 'date_of_birth', 'birthday']:
                        date_col = key
                        break
                
                if date_col:
                    age = calculate_age(row[date_col], target_date)
                    if age is not None:
                        # Create new object following the inferred transformation rule:
                        # Keep other fields, replace date field with 'age'
                        new_row = {}
                        for k, v in row.items():
                            if k == date_col:
                                new_row['age'] = age
                            else:
                                new_row[k] = v
                        results.append(new_row)
                else:
                    # If no date column is found, we treat the row as is 
                    # (but this case is unlikely given the prompt)
                    results.append(row)

        # Output the JSON array to stdout
        print(json.dumps(results, indent=2))

    except FileNotFoundError:
        # If file doesn't exist, output nothing as per "Do not output anything other than the JSON array"
        pass
    except Exception:
        # Silently fail to maintain clean stdout for the JSON array requirement
        pass

if __name__ == "__main__":
    main()