import json

with open('input/data.json', 'r') as f:
    data = json.load(f)

filtered = [record for record in data if record['active'] and record['age'] >= 30]
filtered.sort(key=lambda x: x['name'])
print(json.dumps(filtered))