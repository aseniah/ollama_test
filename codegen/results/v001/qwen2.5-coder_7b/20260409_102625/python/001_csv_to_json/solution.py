import csv
import json

# Define the CSV file path
file_path = 'input/data.csv'

# Initialize an empty list to store the data
data_list = []

# Open the CSV file and read the data
with open(file_path, mode='r', newline='') as file:
    reader = csv.DictReader(file)
    
    # Iterate over each row in the CSV file
    for row in reader:
        # Convert the Age field to integer
        row['Age'] = int(row['Age'])
        
        # Convert the Score field to float
        row['Score'] = float(row['Score'])
        
        # Append the row to the data list
        data_list.append(row)

# Convert the data list to a JSON string and print it
print(json.dumps(data_list, indent=4))