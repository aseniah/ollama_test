import json

with open('input/data.json', 'r') as f:
    data = json.load(f)

filtered_data = [record for record in data if record['active'] and record['age'] >= 30]
sorted_data = sorted(filtered_data, key=lambda x: x['name'])

print(json.dumps(sorted_data))