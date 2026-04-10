import json

# Read the input JSON file
with open('input/data.json', 'r') as f:
    data = json.load(f)

# Filter and sort the records
filtered = [record for record in data 
            if record.get('active', False) and record.get('age', 0) >= 30]
sorted_records = sorted(filtered, key=lambda x: x['name'])

# Output JSON array to stdout
print(json.dumps(sorted_records))