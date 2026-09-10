import csv
import json
from datetime import date

# Read the CSV file
data = []
with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    for row in reader:
        data.append(row)

# Reference date: July 1, 2025
ref_date = date(2025, 7, 1)

result = []

for row in data:
    # Parse name
    parts = row['Name'].split()
    # Assume first name is first word, last name is last word
    # For names like "John Winston Lennon", first name = "John", last name = "Lennon"
    first_name = parts[0]
    last_name = parts[-1]
    
    # Parse birthday
    birthday_str = row['Birthday']  # format: M/D/YYYY
    bmonth, bday, byear = map(int, birthday_str.split('/'))
    birthday_date = date(byear, bmonth, bday)
    
    # Calculate age as of July 1, 2025
    age = ref_date.year - birthday_date.year
    if (ref_date.month, ref_date.day) < (birthday_date.month, birthday_date.day):
        age -= 1
    
    # Format birthday as YYYY-MM-DD
    birthday_formatted = f"{byear:04d}-{bmonth:02d}-{bday:02d}"
    
    # Parse relatives
    relatives = []
    
    father = row['Father']
    if father.lower() != 'null':
        parts = father.split()
        relatives.append({
            "FirstName": parts[0],
            "LastName": parts[-1],
            "Relationship": "Father"
        })
    
    mother = row['Mother']
    if mother.lower() != 'null':
        parts = mother.split()
        relatives.append({
            "FirstName": parts[0],
            "LastName": parts[-1],
            "Relationship": "Mother"
        })
    
    brother = row['Brother']
    if brother.lower() != 'null':
        parts = brother.split()
        relatives.append({
            "FirstName": parts[0],
            "LastName": parts[-1],
            "Relationship": "Brother"
        })
    
    sister = row['Sister']
    if sister.lower() != 'null':
        parts = sister.split()
        relatives.append({
            "FirstName": parts[0],
            "LastName": parts[-1],
            "Relationship": "Sister"
        })
    
    person = {
        "FirstName": first_name,
        "LastName": last_name,
        "Birthday": birthday_formatted,
        "Age": age,
        "Relatives": relatives
    }
    
    result.append(person)

# Output JSON array to stdout
print(json.dumps(result, indent=3))