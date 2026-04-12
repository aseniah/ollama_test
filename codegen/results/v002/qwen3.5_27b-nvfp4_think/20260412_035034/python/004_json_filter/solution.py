import json

# Read the JSON file
with open('input/data.json', 'r') as f:
    data = json.load(f)

# Filter records where active is true and age is 30 or older
filtered = [record for record in data if record['active'] and record['age'] >= 30]

# Sort by name ascending
sorted_records = sorted(filtered, key=lambda x: x['name'])

# Output to stdout as JSON
print(json.dumps(sorted_records))