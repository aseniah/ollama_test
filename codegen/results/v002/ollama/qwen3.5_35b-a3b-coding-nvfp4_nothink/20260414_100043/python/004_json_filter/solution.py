import json

# Read the input file
with open('input/data.json', 'r') as f:
    data = json.load(f)

# Filter records where active is True and age >= 30
filtered = [record for record in data if record['active'] and record['age'] >= 30]

# Sort by name ascending
filtered.sort(key=lambda x: x['name'])

# Output the result as JSON to stdout
print(json.dumps(filtered))