import csv
import json
from datetime import datetime

def calculate_age(birthday_str, reference_date):
    """Calculates age based on the given birthday and reference date."""
    # Input format from CSV is M/D/YYYY
    birth_date = datetime.strptime(birthday_str, "%m/%d/%Y")
    
    # If the person died, age is calculated as of death date (implied by expected_format.json logic)
    # However, the prompt asks for age as of July 1, 2025.
    # Looking at expected_format.json:
    # John Lennon (b 1940, d 1980) -> Age 40. (1980 - 1940 = 40)
    # George Harrison (b 1943, d 2001) -> Age 58. (2001 - 1943 = 58)
    # James McCartney (b 1942, alive) -> Age 83. (2025 - 1942 = 83)
    # Ringo Starr (b 1940, alive) -> Age 84. (2025 - 1940 = 84)
    # This means if 'Died' is present, use death date. Otherwise, use July 1, 2025.
    return birth_date, None

def get_actual_age(birth_date, death_str, ref_date):
    target_date = ref_date
    if death_str and death_str != 'null':
        target_date = datetime.strptime(death_str, "%m/%d/%Y")
    
    age = target_date.year - birth_date.year - ((target_date.month, target_date.day) < (birth_date.month, birth_date.day))
    return age

def parse_name(full_name):
    if not full_name or full_name == 'null':
        return None
    parts = full_name.split()
    # Based on expected_format.json:
    # "John Winston Lennon" -> FirstName: John, LastName: Lennon
    return {"FirstName": parts[0], "LastName": parts[-1]}

def main():
    ref_date = datetime(2025, 7, 1)
    results = []

    try:
        with open('input/input.csv', mode='r', encoding='utf-8') as f:
            reader = csv.DictReader(f)
            for row in reader:
                # Parse main person
                full_name_parts = row['Name'].split()
                first_name = full_name_parts[0]
                last_name = full_name_parts[-1]
                
                # Birthday formatting: 10/9/1940 -> 1940-10-09
                bday_dt = datetime.strptime(row['Birthday'], "%m/%d/%Y")
                bday_str = bday_dt.strftime("%Y-%m-%d")
                
                # Age calculation
                age = get_actual_age(bday_dt, row['Died'], ref_date)
                
                # Relatives mapping
                relatives = []
                rel_mapping = [
                    ('Father', row['Father']),
                    ('Mother', row['Mother']),
                    ('Brother', row['Brother']),
                    ('Sister', row['Sister']),
                ]
                
                for rel_type, rel_name in rel_mapping:
                    if rel_name and rel_name != 'null':
                        name_data = parse_name(rel_name)
                        if name_data:
                            relatives.append({
                                "FirstName": name_data["FirstName"],
                                "LastName": name_data["LastName"],
                                "Relationship": rel_type
                            })
                
                results.append({
                    "FirstName": first_name,
                    "LastName": last_name,
                    "Birthday": bday_str,
                    "Age": age,
                    "Relatives": relatives
                })
                
        print(json.dumps(results, indent=2))
    except FileNotFoundError:
        pass

if __name__ == "__main__":
    main()