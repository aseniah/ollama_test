import csv
import json

# Initialize a list to store the data as dictionaries
data = []

# Open the CSV file and read its contents
with open('input/data.csv', mode='r') as file:
    reader = csv.DictReader(file)
    for row in reader:
        # Convert the Age field to an integer
        row['Age'] = int(row['Age'])
        # Convert the Score field to a float
        row['Score'] = float(row['Score'])
        data.append(row)

# Output the JSON array to stdout
json.dump(data, stdout, indent=2)