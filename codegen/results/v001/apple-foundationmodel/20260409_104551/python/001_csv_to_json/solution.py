import csv
import json

def process_csv_to_json():
    # Open the CSV file for reading
    with open('input/data.csv', mode='r') as csv_file:
        # Create a list to store the JSON objects
        json_data = []
        
        # Read the CSV file
        csv_reader = csv.DictReader(csv_file)
        
        # Iterate through each row (excluding the header)
        for row in csv_reader:
            # Create a dictionary object with the specified fields
            obj = {
                'Name': row['Name'],
                'Age': int(row['Age']),
                'Email': row['Email'],
                'Score': float(row['Score'])
            }
            # Append the dictionary to the JSON data list
            json_data.append(obj)
    
    # Convert the list of dictionaries to a JSON array
    json_array = json.dumps(json_data, indent=2)
    
    # Output the JSON array to standard output
    print(json_array)

# Call the function to process the CSV file
process_csv_to_json()