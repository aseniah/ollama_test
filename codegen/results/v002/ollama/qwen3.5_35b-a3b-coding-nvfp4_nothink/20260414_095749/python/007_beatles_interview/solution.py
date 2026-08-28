import csv
import json
from datetime import date

def calculate_age(birthday_str, reference_date=date(2025, 7, 1)):
    # Parse the birthday string in format M/D/YYYY or MM/DD/YYYY
    parts = birthday_str.split('/')
    year = int(parts[0])
    month = int(parts[1])
    day = int(parts[2])
    birth_date = date(year, month, day)
    
    # Calculate age
    age = reference_date.year - birth_date.year
    # Adjust if birthday hasn't occurred yet in the reference year
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def main():
    input_file = 'input/input.csv'
    
    results = []
    
    with open(input_file, mode='r', newline='', encoding='utf-8') as csvfile:
        reader = csv.DictReader(csvfile)
        for row in reader:
            # Parse name into First and Last
            full_name = row['Name']
            name_parts = full_name.split()
            first_name = name_parts[0]
            last_name = name_parts[-1] # Assuming last name is the last part
            
            birthday = row['Birthday']
            age = calculate_age(birthday)
            
            relatives = []
            
            # Father
            father_name = row['Father']
            if father_name and father_name != 'null':
                f_parts = father_name.split()
                relatives.append({
                    "FirstName": f_parts[0],
                    "LastName": f_parts[-1],
                    "Relationship": "Father"
                })
            
            # Mother
            mother_name = row['Mother']
            if mother_name and mother_name != 'null':
                m_parts = mother_name.split()
                relatives.append({
                    "FirstName": m_parts[0],
                    "LastName": m_parts[-1],
                    "Relationship": "Mother"
                })
            
            # Brother
            brother_name = row['Brother']
            if brother_name and brother_name != 'null':
                b_parts = brother_name.split()
                relatives.append({
                    "FirstName": b_parts[0],
                    "LastName": b_parts[-1],
                    "Relationship": "Brother"
                })
            
            # Sister
            sister_name = row['Sister']
            if sister_name and sister_name != 'null':
                s_parts = sister_name.split()
                relatives.append({
                    "FirstName": s_parts[0],
                    "LastName": s_parts[-1],
                    "Relationship": "Sister"
                })
            
            person = {
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday, # Keep format consistent with expected? Expected is YYYY-MM-DD. Input is M/D/YYYY.
                                      # The expected output shows "Birthday": "1940-10-09" for input "10/9/1940".
                                      # So we need to reformat the date string.
                "Age": age,
                "Relatives": relatives
            }
            
            # Reformat Birthday to YYYY-MM-DD
            parts = birthday.split('/')
            person["Birthday"] = f"{int(parts[0]):02d}-{int(parts[1]):02d}-{int(parts[2]):04d}"
            
            results.append(person)

    print(json.dumps(results, indent=2))

if __name__ == "__main__":
    main()