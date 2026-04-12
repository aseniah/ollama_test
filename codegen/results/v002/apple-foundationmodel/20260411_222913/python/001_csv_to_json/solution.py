import csv
import json

def csv_to_json():
    # Open the CSV file
    with open('input/data.csv', mode='r', newline='', encoding='utf-8') as csvfile:
        # Create a CSV reader object
        csv_reader = csv.DictReader(csvfile)
        
        # Prepare a list to store JSON objects
        json_objects = []
        
        # Iterate over each row in the CSV file
        for row in csv_reader:
            # Create a JSON object for the current row
            json_object = {
                'Name': row['Name'],
                'Age': int(row['Age']),
                'Email': row['Email'],
                'Score': float(row['Score'])
            }
            # Append the JSON object to the list
            json_objects.append(json_object)
    
    # Convert the list of JSON objects to a JSON string
    json_string = json.dumps(json_objects, indent=4)
    
    # Print the JSON string to stdout
    print(json_string)

# Run the function
csv_to_json()