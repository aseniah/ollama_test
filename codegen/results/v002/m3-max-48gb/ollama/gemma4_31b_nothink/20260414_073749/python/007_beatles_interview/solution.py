import csv
import json
from datetime import datetime

def calculate_age(birthday_str, reference_date):
    # Birthday format is M/D/YYYY based on input.csv
    birth_date = datetime.strptime(birthday_str, '%m/%d/%Y')
    
    # Age is calculated by comparing year and then adjusting if birth date 
    # hasn't occurred yet in the reference year
    age = reference_date.year - birth_date.year
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def main():
    ref_date = datetime(2025, 7, 1)
    results = []

    try:
        with open('input/input.csv', mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # Process Name (Full Name -> First and Last)
                name_parts = row['Name'].split(' ')
                first_name = name_parts[0]
                last_name = name_parts[-1]

                # Format Birthday to YYYY-MM-DD
                bday_obj = datetime.strptime(row['Birthday'], '%m/%d/%Y')
                birthday_iso = bday_obj.strftime('%Y-%m-%d')

                # Calculate Age
                age = calculate_age(row['Birthday'], ref_date)

                # Process Relatives
                relatives = []
                rel_mapping = {
                    'Father': row['Father'],
                    'Mother': row['Mother'],
                    'Brother': row['Brother'],
                    'Sister': row['Sister']
                }

                for rel_type, rel_name in rel_mapping.items():
                    if rel_name and rel_name.lower() != 'null':
                        rel_parts = rel_name.split(' ')
                        relatives.append({
                            "FirstName": rel_parts[0],
                            "LastName": rel_parts[-1],
                            "Relationship": rel_type
                        })

                # Construct Person object
                person = {
                    "FirstName": first_name,
                    "LastName": last_name,
                    "Birthday": birthday_iso,
                    "Age": age,
                    "Relatives": relatives
                }
                results.append(person)

        # Output to stdout
        print(json.dumps(results, indent=2))

    except FileNotFoundError:
        pass

if __name__ == "__main__":
    main()