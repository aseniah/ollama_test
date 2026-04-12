import json

# Read the JSON data from the file
with open('input/data.json', 'r') as file:
    data = json.load(file)

# Filter and sort the records
filtered_data = sorted(
    [record for record in data if record['active'] and record['age'] >= 30],
    key=lambda x: x['name']
)

# Output the filtered and sorted data to stdout
json.dump(filtered_data, stdout, indent=2)