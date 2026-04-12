import json

try:
    with open('input/data.json', 'r') as f:
        data = json.load(f)
except FileNotFoundError:
    import sys
    sys.exit(1)

filtered_data = [record for record in data if record.get('active') is True and record.get('age', 0) >= 30]
sorted_data = sorted(filtered_data, key=lambda x: x['name'])

print(json.dumps(sorted_data))