import csv
import json
from datetime import date

def calculate_age(birthday, as_of):
    """
    Calculate age as of a given date.
    If the person has died before the as_of date, calculate age at death.
    Otherwise, calculate age as of the as_of date.
    """
    birthday_date = datetime.datetime.strptime(birthday, "%m/%d/%Y").date()
    if died:
        death_date = datetime.datetime.strptime(died, "%m/%d/%Y").date()
        if death_date < as_of:
            return (death_date.year - birthday_date.year) - ((death_date.month, death_date.day) < (birthday_date.month, birthday_date.day))
    # Calculate age as of as_of date
    if (as_of.month, as_of.day) < (birthday_date.month, birthday_date.day):
        return as_of.year - birthday_date.year - 1
    else:
        return as_of.year - birthday_date.year

import datetime

as_of = date(2025, 7, 1)
results = []

with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    for row in reader:
        # Parse name
        parts = row['Name'].split()
        first_name = parts[0]
        last_name = parts[-1]
        
        birthday = row['Birthday']
        died = row['Died']
        
        # Calculate age
        birthday_date = datetime.datetime.strptime(birthday, "%m/%d/%Y").date()
        if died and died != 'null':
            death_date = datetime.datetime.strptime(died, "%m/%d/%Y").date()
            if death_date < as_of:
                # Age at death
                age = (death_date.year - birthday_date.year) - ((death_date.month, death_date.day) < (birthday_date.month, birthday_date.day))
            else:
                # Age as of as_of date (if died after as_of or on as_of, treat as alive on as_of)
                age = (as_of.year - birthday_date.year) - ((as_of.month, as_of.day) < (birthday_date.month, birthday_date.day))
        else:
            # Alive on as_of date
            age = (as_of.year - birthday_date.year) - ((as_of.month, as_of.day) < (birthday_date.month, birthday_date.day))
        
        # Build relatives list
        relatives = []
        
        father = row['Father']
        if father and father != 'null':
            father_parts = father.split()
            relatives.append({
                "FirstName": father_parts[0],
                "LastName": father_parts[-1],
                "Relationship": "Father"
            })
        
        mother = row['Mother']
        if mother and mother != 'null':
            mother_parts = mother.split()
            relatives.append({
                "FirstName": mother_parts[0],
                "LastName": mother_parts[-1],
                "Relationship": "Mother"
            })
        
        brother = row['Brother']
        if brother and brother != 'null':
            brother_parts = brother.split()
            relatives.append({
                "FirstName": brother_parts[0],
                "LastName": brother_parts[-1],
                "Relationship": "Brother"
            })
        
        sister = row['Sister']
        if sister and sister != 'null':
            sister_parts = sister.split()
            relatives.append({
                "FirstName": sister_parts[0],
                "LastName": sister_parts[-1],
                "Relationship": "Sister"
            })
        
        # Format birthday as YYYY-MM-DD
        birthday_formatted = datetime.datetime.strptime(birthday, "%m/%d/%Y").strftime("%Y-%m-%d")
        
        results.append({
            "FirstName": first_name,
            "LastName": last_name,
            "Birthday": birthday_formatted,
            "Age": age,
            "Relatives": relatives
        })

print(json.dumps(results, indent=2))