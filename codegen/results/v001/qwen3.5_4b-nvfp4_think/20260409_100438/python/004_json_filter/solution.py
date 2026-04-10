import json

with open('input/data.json', 'r') as f:
    data = json.load(f)

filtered = [
    record for record in data 
    if record['active'] is True and record['age'] >= 30
]

filtered_sorted = sorted(filtered, key=lambda x: x['name'])

print(json.dumps(filtered_sorted))