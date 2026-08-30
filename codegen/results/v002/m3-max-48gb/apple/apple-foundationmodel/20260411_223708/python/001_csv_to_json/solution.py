import csv
import json

def read_csv_to_json():
    # Open the CSV file
    with open('input/data.csv', mode='r', newline='') as csvfile:
        # Create a CSV reader object
        csvreader = csv.DictReader(csvfile)
        
        # Initialize an empty list to store the JSON objects
        json_data = []
        
        # Iterate over each row in the CSV file
        for row in csvreader:
            # Create a dictionary from the row data
            data_object = {
                'Name': row['Name'],
                'Age': int(row['Age']),
                'Email': row['Email'],
                'Score': float(row['Score'])
            }
            
            # Append the dictionary to the list
            json_data.append(data_object)
    
    # Convert the list of dictionaries to a JSON string
    json_output = json.dumps(json_data, indent=4)
    
    # Print the JSON output to stdout
    print(json_output)

# Call the function to execute
read_csv_to_json()