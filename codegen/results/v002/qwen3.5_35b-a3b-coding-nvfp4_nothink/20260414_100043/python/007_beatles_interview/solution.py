import csv
import json
from datetime import date

def calculate_age(birthdate_str, reference_date):
    birth_date = date.fromisoformat(birthdate_str)
    age = reference_date.year - birth_date.year
    if (reference_date.month, reference_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def parse_name(name_str):
    if not name_str or name_str.strip() == '':
        return None, None
    parts = name_str.strip().split()
    if not parts:
        return None, None
    first_name = parts[0]
    last_name = ' '.join(parts[1:]) if len(parts) > 1 else ''
    return first_name, last_name

def process_person(row):
    name_str = row['Name']
    first_name, last_name = parse_name(name_str)
    
    birthday_str = row['Birthday']
    # Convert MM/DD/YYYY to YYYY-MM-DD
    parts = birthday_str.split('/')
    if len(parts) == 3:
        month, day, year = parts
        birthday_formatted = f"{year}-{month.zfill(2)}-{day.zfill(2)}"
    else:
        birthday_formatted = birthday_str  # Fallback if format is unexpected
    
    ref_date = date(2025, 7, 1)
    age = calculate_age(birthday_formatted, ref_date)
    
    relatives = []
    
    father_str = row['Father']
    if father_str and father_str != 'null' and father_str.strip():
        f_first, f_last = parse_name(father_str)
        if f_first:
            relatives.append({"FirstName": f_first, "LastName": f_last, "Relationship": "Father"})
    
    mother_str = row['Mother']
    if mother_str and mother_str != 'null' and mother_str.strip():
        m_first, m_last = parse_name(mother_str)
        if m_first:
            relatives.append({"FirstName": m_first, "LastName": m_last, "Relationship": "Mother"})
    
    brother_str = row['Brother']
    if brother_str and brother_str != 'null' and brother_str.strip():
        b_first, b_last = parse_name(brother_str)
        if b_first:
            relatives.append({"FirstName": b_first, "LastName": b_last, "Relationship": "Brother"})
    
    sister_str = row['Sister']
    if sister_str and sister_str != 'null' and sister_str.strip():
        s_first, s_last = parse_name(sister_str)
        if s_first:
            relatives.append({"FirstName": s_first, "LastName": s_last, "Relationship": "Sister"})
    
    return {
        "FirstName": first_name,
        "LastName": last_name,
        "Birthday": birthday_formatted,
        "Age": age,
        "Relatives": relatives
    }

def main():
    input_file = "input/input.csv"
    
    result = []
    with open(input_file, mode='r', encoding='utf-8') as csvfile:
        reader = csv.DictReader(csvfile)
        for row in reader:
            person = process_person(row)
            result.append(person)
    
    print(json.dumps(result, indent=None))

if __name__ == "__main__":
    main()