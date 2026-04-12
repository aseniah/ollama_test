import csv
import json
from datetime import datetime, timedelta

# Read the CSV file and calculate ages as of July 1, 2025
def calculate_ages(filename):
    with open(filename, newline='', encoding='utf-8') as csvfile:
        reader = csv.DictReader(csvfile)
        data = []
        
        for row in reader:
            birthdate = datetime.strptime(row['Birthday'], '%m/%d/%Y')
            birthdate += timedelta(days=7)  # Adjust for July 1, 2025
            
            age = (datetime.strptime('2025-07-01', '%Y-%m-%d') - birthdate).days // 365
            
            # Create a new entry with the calculated age
            new_entry = {
                'FirstName': row['Name'],
                'LastName': row['LastName'],
                'Birthday': row['Birthday'],
                'Age': age,
                'Relatives': row['Relatives']
            }
            data.append(new_entry)
    
    return data

# Main function
def main():
    # Calculate ages from the CSV file
    ages = calculate_ages('input/input.csv')
    
    # Output the JSON array to stdout
    print(json.dumps(ages, indent=2))

# Run the program
if __name__ == "__main__":
    main()