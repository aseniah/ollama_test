import json

# Read input data
with open('input/data.json', 'r') as f:
    data = json.load(f)

# Filter and process records
filtered_data = [
    {'name': r['name'], 'age': r['age'], 'active': r['active'], 'score': r['score']}
    for r in data
    if r['active'] and r['age'] >= 30
]

# Sort by name ascending
filtered_data.sort(key=lambda x: x['name'])

# Output JSON to stdout
print(json.dumps(filtered_data))