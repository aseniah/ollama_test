import json

# Read the input JSON file
with open('input/data.json', 'r') as f:
    data = json.load(f)

# Filter records where active is true and age is 30 or older
filtered_records = [
    record for record in data
    if record['active'] is True and record['age'] >= 30
]

# Sort by name ascending
filtered_records.sort(key=lambda x: x['name'])

# Output the result as a JSON array to stdout
print(json.dumps(filtered_records))