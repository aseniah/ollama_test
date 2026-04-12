import csv
import json
import sys
from datetime import datetime

def calculate_age(birthday_str, reference_date):
    """Calculates age as of the reference date."""
    # Input format in CSV is M/D/YYYY based on example
    birthday = datetime.strptime(birthday_str, "%m/%d/%Y")
    age = reference_date.year - birthday.year
    # Check if birthday has occurred yet in the reference year
    if (reference_date.month, reference_date.day) < (birthday.month, birthday.day):
        age -= 1
    return age

def solve():
    input_file = 'input/input.csv'
    reference_date = datetime(2025, 7, 1)
    output_data = []

    try:
        with open(input_file, mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # Parse Name (Split full name into First and Last)
                # The CSV 'Name' column contains 'FirstName MiddleName LastName' or 'FirstName LastName'
                # Based on expected format, we need to separate First and Last.
                # Looking at 'John Winston Lennon' -> 'John' and 'Lennon'
                # The expected format shows "FirstName": "John", "LastName": "Lennon"
                full_name_parts = row['Name'].split()
                first_name = full_name_parts[0]
                last_name = full_name_parts[-1]

                # Parse Birthday and Age
                birthday_str = row['Birthday']
                age = calculate_age(birthday_str, reference_date)

                # Parse Relatives
                relatives = []
                
                # Helper to process relative fields
                def add_relative(field_name, relationship):
                    val = row.get(field_name)
                    if val and val.lower() != 'null':
                        # The CSV format for relatives is "FirstName LastName"
                        rel_parts = val.split()
                        rel_first = rel_parts[0]
                        rel_last = rel_parts[-1] if len(rel_parts) > 1 else ""
                        relatives.append({
                            "FirstName": rel_first,
                            "LastName": rel_last,
                            "Relationship": relationship
                        })

                add_relative('Father', 'Father')
                add_relative('Mother', 'Mother')
                add_relative('Brother', 'Brother')
                add_relative('Sister', 'Sister')

                # Construct object
                person_obj = {
                    "FirstName": first_name,
                    "LastName": last_name,
                    "Birthday": datetime.strptime(birthday_str, "%m/%d/%Y").strftime("%Y-%m-%d"),
                    "Age": age,
                    "Relatives": relatives
                }
                output_data.append(person_obj)

        # Output the JSON array to stdout
        print(json.dumps(output_data, indent=2))

    except FileNotFoundError:
        pass

if __name__ == "__main__":
    solve()