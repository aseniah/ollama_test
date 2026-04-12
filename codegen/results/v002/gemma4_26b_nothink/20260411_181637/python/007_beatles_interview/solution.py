import csv
import json
from datetime import datetime

def calculate_age(birthday_str, reference_date):
    """Calculates age as of a specific reference date."""
    # The input format for birthday is M/D/YYYY based on input.csv
    birthday = datetime.strptime(birthday_str, '%m/%d/%Y')
    age = reference_date.year - birthday.year
    # Adjust if birthday hasn't occurred yet in the reference year
    if (reference_date.month, reference_date.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def main():
    reference_date = datetime(2025, 7, 1)
    input_file = 'input/input.csv'
    output_data = []

    try:
        with open(input_file, mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # Parse Name into FirstName and LastName
                # The CSV contains 'Full Name' in the 'Name' column.
                # Looking at the expected JSON, we split by space.
                # Note: John Winston Lennon -> FirstName: John, LastName: Lennon? 
                # Actually, the JSON shows John Lennon. 
                # Let's look closer at the input/expected mapping.
                # Input: John Winston Lennon -> Output: FirstName: John, LastName: Lennon
                # Input: James Paul McCartney -> Output: FirstName: James, LastName: McCartney
                # It seems we take the first and the last part of the name string.
                name_parts = row['Name'].split()
                first_name = name_parts[0]
                last_name = name_parts[-1]

                birthday_str = row['Birthday']
                age = calculate_age(birthday_str, reference_date)
                
                # Reformat Birthday to YYYY-MM-DD
                birthday_dt = datetime.strptime(birthday_str, '%m/%d/%Y')
                birthday_formatted = birthday_dt.strftime('%Y-%m-%d')

                relatives = []

                # Helper to parse relative names
                # The CSV contains full names for parents like 'Alfred Lennon'
                # The JSON expects FirstName and LastName separately.
                def parse_relative(full_name, relationship):
                    if not full_name or full_name.lower() == 'null':
                        return None
                    parts = full_name.split()
                    return {
                        "FirstName": parts[0],
                        "LastName": parts[-1],
                        "Relationship": relationship
                    }

                # Process Father
                father = parse_relative(row.get('Father'), 'Father')
                if father:
                    relatives.append(father)

                # Process Mother
                mother = parse_relative(row.get('Mother'), 'Mother')
                if mother:
                    relatives.append(mother)

                # Process Brother
                brother = parse_relative(row.get('Brother'), 'Brother')
                if brother:
                    relatives.append(brother)

                # Process Sister
                sister = parse_relative(row.get('Sister'), 'Sister')
                if sister:
                    relatives.append(sister)

                output_data.append({
                    "FirstName": first_name,
                    "LastName": last_name,
                    "Birthday": birthday_formatted,
                    "Age": age,
                    "Relatives": relatives
                })

        print(json.dumps(output_data, indent=2))

    except FileNotFoundError:
        pass

if __name__ == "__main__":
    main()