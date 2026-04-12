import json

with open('input/data.json', 'r') as f:
    data = json.load(f)

filtered_records = [record for record in data if record['active'] and record['age'] >= 30]
filtered_records.sort(key=lambda x: x['name'])

print(json.dumps(filtered_records))