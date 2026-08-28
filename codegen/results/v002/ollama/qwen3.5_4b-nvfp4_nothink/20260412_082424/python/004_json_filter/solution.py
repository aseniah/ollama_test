import json

# Read the input file
with open('input/data.json', 'r') as f:
    data = json.load(f)

# Filter records where active is True and age is 30 or older
filtered_data = [
    record for record in data 
    if record.get('active') == True and record.get('age', -1) >= 30
]

# Sort by name ascending (case-sensitive, which matches the requirement)
filtered_data.sort(key=lambda x: x['name'])

# Output to stdout as JSON
print(json.dumps(filtered_data))