import csv
import json

# Define the header fields for reference
header = ['Name', 'Age', 'Email', 'Score']

# Initialize an empty list to store the data objects
data_list = []

# Open the CSV file and read its contents
with open('input/data.csv', newline='') as csvfile:
    reader = csv.DictReader(csvfile)
    
    # Iterate over each row in the CSV file (excluding the header)
    for row in reader:
        # Convert each value to the appropriate type
        name = row['Name']
        age = int(row['Age'])
        email = row['Email']
        score = float(row['Score'])
        
        # Create a dictionary for each row and add it to the list
        data_list.append({
            'Name': name,
            'Age': age,
            'Email': email,
            'Score': score
        })

# Convert the list of dictionaries to a JSON array and print it to stdout
print(json.dumps(data_list, indent=2))