import csv
import json
from datetime import datetime

def calculate_age(birthday_str):
    # Parse date from MM/D/YYYY or M/D/YYYY format
    try:
        birth_date = datetime.strptime(birthday_str, "%m/%d/%Y")
    except ValueError:
        # Fallback if parsing fails (though input should be valid)
        return 0
    
    target_date = datetime(2025, 7, 1)
    age = target_date.year - birth_date.year
    
    # Adjust if birthday hasn't occurred yet in the target year
    if (target_date.month, target_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    
    return age

def parse_full_name(name):
    # Split name into parts. Assuming last name is the last word for this specific dataset context
    # based on the expected output where "John Winston Lennon" -> "John", "Lennon"
    # However, for "James Paul McCartney" -> "James", "McCartney"
    # And "Ringo Starr" -> "Ringo", "Starr"
    parts = name.split()
    first_name = parts[0]
    last_name = parts[-1]
    return first_name, last_name

def format_relative_name(name):
    # If the name in the CSV is "Mike McGear", output should be "Mike", "McGear"
    # If it's "null", we handle it separately
    if not name or name.lower() == 'null':
        return None
    
    parts = name.split()
    return {
        "FirstName": parts[0],
        "LastName": parts[-1]
    }

def main():
    input_file = 'input/input.csv'
    output = []

    try:
        with open(input_file, 'r', newline='', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            
            for row in reader:
                # Extract and parse main person
                full_name = row['Name']
                first_name, last_name = parse_full_name(full_name)
                birthday = row['Birthday']
                age = calculate_age(birthday)
                
                # Extract relatives
                relatives = []
                
                # Map relationships to CSV columns
                rel_map = [
                    ('Father', row['Father']),
                    ('Mother', row['Mother']),
                    ('Brother', row['Brother']),
                    ('Sister', row['Sister'])
                ]
                
                for relationship, name_str in rel_map:
                    if name_str and name_str.lower() != 'null':
                        relative_data = format_relative_name(name_str)
                        if relative_data:
                            relative_data['Relationship'] = relationship
                            relatives.append(relative_data)
                
                # Construct the person object
                person = {
                    "FirstName": first_name,
                    "LastName": last_name,
                    "Birthday": birthday, # Keep original string format as per expected output "1940-10-09" wait, input is "10/9/1940"
                    # Correction: The expected output format for Birthday is "YYYY-MM-DD".
                    # Input is "MM/D/YYYY". We need to convert.
                }
                
                # Re-parse birthday to format as YYYY-MM-DD
                birth_date_obj = datetime.strptime(birthday, "%m/%d/%Y")
                person["Birthday"] = birth_date_obj.strftime("%Y-%m-%d")
                person["Age"] = age
                person["Relatives"] = relatives
                
                output.append(person)
                
        print(json.dumps(output, indent=2))
        
    except FileNotFoundError:
        print("Error: input/input.csv not found.")
    except Exception as e:
        print(f"Error: {e}")

if __name__ == "__main__":
    main()