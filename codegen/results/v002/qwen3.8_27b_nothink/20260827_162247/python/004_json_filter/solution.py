import json

# Read the input data
with open('input/data.json', 'r') as f:
    data = json.load(f)

# Filter records where active is True and age is 30 or older
filtered = [record for record in data if record['active'] and record['age'] >= 30]

# Sort by name ascending
filtered.sort(key=lambda x: x['name'])

# Output as JSON array
print(json.dumps(filtered, indent=2))