import json

# Read the JSON data from file
with open('input/data.json', 'r') as file:
    data = json.load(file)

# Filter records where active is true and age is 30 or older
filtered_data = [record for record in data if record['active'] and record['age'] >= 30]

# Sort the filtered data by name ascending
sorted_data = sorted(filtered_data, key=lambda x: x['name'])

# Output the sorted JSON array to stdout
print(json.dumps(sorted_data, indent=2))