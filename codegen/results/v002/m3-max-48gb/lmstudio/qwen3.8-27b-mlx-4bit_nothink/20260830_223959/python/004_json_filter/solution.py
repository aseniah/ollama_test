import json

with open('input/data.json', 'r') as f:
    data = json.load(f)

# Filter: active is true and age >= 30
filtered = [record for record in data if record['active'] and record['age'] >= 30]

# Sort by name ascending
filtered.sort(key=lambda x: x['name'])

# Output JSON array
print(json.dumps(filtered, indent=2))