import csv
import json
from datetime import datetime

def calculate_age(birth_date, reference_date):
    birth = datetime.strptime(birth_date, "%m/%d/%Y")
    age = reference_date.year - birth.year
    if reference_date.month < birth.month or (reference_date.month == birth.month and reference_date.day < birth.day):
        age -= 1
    return age

def main():
    input_file = "input/input.csv"
    output_file = "output.json"
    
    # Read CSV file
    with open(input_file, "r") as f:
        reader = csv.DictReader(f)
        data = []
        
        for row in reader:
            # Process each row
            person = {
                "FirstName": row["Name"].split()[0],
                "LastName": row["Name"].split()[-1],
                "Birthday": row["Birthday"],
                "Age": calculate_age(row["Birthday"], datetime(2025, 7, 1)),
                "Relatives": []
            }
            
            # Process relatives
            relatives = []
            if row["Father"] != "null":
                relatives.append({
                    "FirstName": row["Father"].split()[0],
                    "LastName": row["Father"].split()[-1],
                    "Relationship": "Father"
                })
            if row["Mother"] != "null":
                relatives.append({
                    "FirstName": row["Mother"].split()[0],
                    "LastName": row["Mother"].split()[-1],
                    "Relationship": "Mother"
                })
            if row["Brother"] != "null":
                relatives.append({
                    "FirstName": row["Brother"].split()[0],
                    "LastName": row["Brother"].split()[-1],
                    "Relationship": "Brother"
                })
            if row["Sister"] != "null":
                relatives.append({
                    "FirstName": row["Sister"].split()[0],
                    "LastName": row["Sister"].split()[-1],
                    "Relationship": "Sister"
                })
            
            person["Relatives"] = relatives
            data.append(person)
        
        print(json.dumps(data, indent=2))

if __name__ == "__main__":
    main()