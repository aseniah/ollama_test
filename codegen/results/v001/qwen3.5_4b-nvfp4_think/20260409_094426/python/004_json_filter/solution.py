import json

with open('input/data.json', 'r') as f:
    data = json.load(f)

filtered_data = [
    record for record in data 
    if record.get('active', False) and record.get('age', 0) >= 30
]

filtered_data.sort(key=lambda x: x['name'])

print(json.dumps(filtered_data))