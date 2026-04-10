import json

# Read the input file
with open('input/data.json', 'r') as f:
    data = json.load(f)

# Filter and sort
result = sorted([record for record in data if record['active'] and record['age'] >= 30], key=lambda x: x['name'])

# Output as JSON array
print(json.dumps(result))