import json

# Read the input file
with open('input/data.json', 'r') as f:
    data = json.load(f)

# Filter records where active is True and age >= 30, then sort by name
filtered_data = [
    record for record in data 
    if record['active'] is True and record['age'] >= 30
]
filtered_data.sort(key=lambda x: x['name'])

# Output the result as JSON
print(json.dumps(filtered_data))