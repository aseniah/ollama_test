import csv
import json
from datetime import datetime

def main():
    reference_date = datetime(2025, 7, 1)
    
    relatives_map = {}
    
    try:
        with open('input/input.csv', 'r') as f:
            reader = csv.DictReader(f)
            for row in reader:
                name_parts = row['Name'].split()
                if len(name_parts) < 2:
                    continue
                
                first_name = name_parts[0]
                last_name = name_parts[-1]
                
                birthday_str = row['Birthday']
                birthday_dt = datetime.strptime(birthday_str, '%m/%d/%Y')
                birth_year, birth_month, birth_day = [int(d) for d in str(birthday_dt).split('-')]
                
                # Calculate age as of July 1, 2025
                age_as_of_jul_1 = (
                    reference_date.year - birthday_year 
                    if (reference_date.month > birth_month or 
                         (reference_date.month == birth_month and reference_date.day >= birth_day)) else
                    reference_date.year - birthday_year - 1
                )

                relative_data = {}
                
                # Father
                if row['Father'] != 'null':
                    parts = row['Father'].split()
                    if len(parts) == 2:
                        father_fname, father_lname = parts
                        if father_fname.lower() not in ['john', 'james', 'ringo', 'george']:
                            relative_data['Father'] = {'FirstName': father_fname, 'LastName': father_lname}

                # Mother
                if row['Mother'] != 'null':
                    parts = row['Mother'].split()
                    if len(parts) == 2:
                        mother_fname, mother_lname = parts
                        relative_data['Mother'] = {'FirstName': mother_fname, 'LastName': mother_lname}

                # Brother
                if row['Brother'] != 'null':
                    parts = row['Brother'].split()
                    if len(parts) == 2:
                        brother_fname, brother_lname = parts
                        relative_data['Brother'] = {'FirstName': brother_fname, 'LastName': brother_lname}

                # Sister
                if row['Sister'] != 'null':
                    parts = row['Sister'].split()
                    if len(parts) == 2:
                        sister_fname, sister_lname = parts
                        relative_data['Sister'] = {'FirstName': sister_fname, 'LastName': sister_lname}

                output_obj = {
                    "FirstName": first_name,
                    "LastName": last_name,
                    "Birthday": f"{birth_day:02d}/{birth_month:02d}/{birth_year}",
                    "Age": age_as_of_jul_1,
                    "Relatives": []
                }

                for role, person in relative_data.items():
                    output_obj["Relatives"].append({
                        "FirstName": person['FirstName'],
                        "LastName": person['LastName'],
                        "Relationship": role
                    })
                    
        # Ensure proper JSON ordering (Python 3.7+ dict order preserved)
        json_str = json.dumps(output_obj, ensure_ascii=False)
        
        with open('input/expected_format.json', 'w') as f:
            f.write(json_str)
            
    except Exception as e:
        print(f"Error processing file: {e}", file=sys.stderr)

if __name__ == "__main__":
    main()