import json

with open('input/data.json', 'r') as f:
    data = json.load(f)

result = [
    record for record in data
    if record['active'] and record['age'] >= 30
]

result.sort(key=lambda x: x['name'])

print(json.dumps(result))