import csv
import json
from datetime import datetime

def calculate_age(birthday_str, target_date):
    # Expected input format from CSV: M/D/YYYY
    birth_date = datetime.strptime(birthday_str, '%m/%d/%Y')
    
    # Calculate age: Target Year - Birth Year
    # Subtract 1 if target date is before the birthday in the target year
    age = target_date.year - birth_date.year
    if (target_date.month, target_date.day) < (birth_date.month, birth_date.day):
        age -= 1
    return age

def split_name(full_name):
    if not full_name or full_name.lower() == 'null':
        return None
    parts = full_name.split()
    if len(parts) > 2:
        # Handle cases like "John Winston Lennon" -> FirstName: John, LastName: Lennon
        return {"FirstName": parts[0], "LastName": parts[-1]}
    elif len(parts) == 2:
        return {"FirstName": parts[0], "LastName": parts[1]}
    else:
        return {"FirstName": parts[0], "LastName": ""}

def main():
    target_date = datetime(2025, 7, 1)
    results = []

    try:
        with open('input/input.csv', mode='r', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                # Process main person name
                name_data = split_name(row['Name'])
                
                # Format Birthday to YYYY-MM-DD
                bday_obj = datetime.strptime(row['Birthday'], '%m/%d/%Y')
                birthday_formatted = bday_obj.strftime('%Y-%m-%d')
                
                # Age calculation
                age = calculate_age(row['Birthday'], target_date)
                
                # Process relatives
                relatives = []
                rel_cols = [('Father', 'Father'), ('Mother', 'Mother'), 
                            ('Brother', 'Brother'), ('Sister', 'Sister')]
                
                for col_name, relationship in rel_cols:
                    val = row.get(col_name)
                    if val and val.lower() != 'null':
                        rel_name = split_name(val)
                        if rel_name:
                            rel_name['Relationship'] = relationship
                            relatives.append(rel_name)
                
                person = {
                    "FirstName": name_data["FirstName"],
                    "LastName": name_data["LastName"],
                    "Birthday": birthday_formatted,
                    "Age": age,
                    "Relatives": relatives
                }
                results.append(person)

        print(json.dumps(results, indent=2))
    except FileNotFoundError:
        pass

if __name__ == "__main__":
    main()