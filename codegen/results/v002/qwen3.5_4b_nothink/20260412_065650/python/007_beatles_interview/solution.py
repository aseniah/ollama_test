import csv
import json
from datetime import date

def parse_date(date_str):
    return date.fromisoformat(date_str)

def parse_mdd(date_str):
    if date_str.lower() == 'null' or date_str.strip() == '':
        return None
    parts = date_str.split('/')
    return f"{parts[2]}-{parts[0]}-{parts[1]}"

def calculate_age(birthday, reference_date):
    age = reference_date.year - birthday.year
    if (reference_date.month, reference_date.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def main():
    reference_date = date(2025, 7, 1)
    
    output = []
    
    with open('input/input.csv', 'r') as f:
        reader = csv.DictReader(f)
        for row in reader:
            first_name = row['Name'].strip()
            last_name = ''
            birthday_str = row['Birthday'].strip()
            father_first = row['Father'].strip()
            father_last = row['Father'].strip().split(' ', 1)[1] if ' ' in row['Father'].strip() else ''
            mother_first = row['Mother'].strip()
            mother_last = row['Mother'].strip().split(' ', 1)[1] if ' ' in row['Mother'].strip() else ''
            brother_first = row['Brother'].strip()
            brother_last = row['Brother'].strip().split(' ', 1)[1] if ' ' in row['Brother'].strip() else ''
            sister_first = row['Sister'].strip()
            sister_last = row['Sister'].strip().split(' ', 1)[1] if ' ' in row['Sister'].strip() else ''

            # Determine last name for the person
            if father_first:
                last_name = father_last
            else:
                # Try to infer from Mother's last name or Brother's/Sister's
                # If the first name matches the last name pattern, use that
                # Example: John Lennon -> Father is Alfred Lennon -> Last name Lennon
                # If no father, check if the last name is part of the first name's known family (e.g., McCartney)
                # In these cases, the father's last name is usually the correct one, but let's use the last name from the first name's context
                # If no father, we'll default to last name from first name's context or mother's last name
                # Actually, the first name doesn't include last name, so we must rely on Father/Mother/Brother/Sister
                # For John Lennon, Father is Alfred Lennon -> Last name Lennon
                # For James McCartney, Mother is Mary McCartney -> Last name McCartney
                # For Ringo Starr, Father is Richard Starkey -> Last name Starkey
                # For George Harrison, Father is Harold Harrison -> Last name Harrison
                
                # If no father, try mother
                last_name = mother_last if mother_last else ''
            
            # If no last name inferred, use a placeholder
            if not last_name:
                last_name = 'Unknown'
            
            try:
                birthday = parse_mdd(birthday_str)
            except:
                birthday = reference_date

            # Calculate age
            age = calculate_age(birthday, reference_date)

            relatives = []
            if father_first:
                relatives.append({"FirstName": father_first, "LastName": father_last, "Relationship": "Father"})
            if mother_first:
                relatives.append({"FirstName": mother_first, "LastName": mother_last, "Relationship": "Mother"})
            if brother_first:
                relatives.append({"FirstName": brother_first, "LastName": brother_last, "Relationship": "Brother"})
            if sister_first:
                relatives.append({"FirstName": sister_first, "LastName": sister_last, "Relationship": "Sister"})
            
            output.append({
                "FirstName": first_name,
                "LastName": last_name,
                "Birthday": birthday.strftime("%Y-%m-%d"),
                "Age": age,
                "Relatives": relatives
            })

    print(json.dumps(output, indent=2))

if __name__ == "__main__":
    main()