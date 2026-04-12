import json

with open('input/data.json', 'r') as f:
    data = json.load(f)

filtered_records = [
    record for record in data 
    if record['active'] and record['age'] >= 30
]

sorted_records = sorted(filtered_records, key=lambda x: x['name'])

print(json.dumps(sorted_records))