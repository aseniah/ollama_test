import json

# Read the input file
with open('input/data.json', 'r') as file:
    data = json.load(file)

# Filter records where active is true and age is 30 or older
filtered_data = [record for record in data if record['active'] and record['age'] >= 30]

# Sort by name ascending
filtered_data.sort(key=lambda x: x['name'])

# Output the result
print(json.dumps(filtered_data))