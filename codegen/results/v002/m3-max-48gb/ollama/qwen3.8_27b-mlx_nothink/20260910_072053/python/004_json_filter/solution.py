import json
import sys

# Read the input file
with open('input/data.json', 'r') as f:
    data = json.load(f)

# Filter: active is true and age >= 30
filtered = [record for record in data if record.get('active') is True and record.get('age', 0) >= 30]

# Sort by name ascending
filtered.sort(key=lambda x: x['name'])

# Output as JSON array
print(json.dumps(filtered, indent=None))