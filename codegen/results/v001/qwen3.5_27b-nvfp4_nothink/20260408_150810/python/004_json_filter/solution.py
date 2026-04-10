import json

# Read the input data
try:
    with open('input/data.json', 'r') as f:
        data = json.load(f)
except FileNotFoundError:
    print("[]")
    exit()
except json.JSONDecodeError:
    print("[]")
    exit()

# Filter and sort the records
filtered_data = [
    record for record in data
    if isinstance(record, dict) 
    and record.get('active') is True 
    and isinstance(record.get('age'), int) 
    and record['age'] >= 30
]

sorted_data = sorted(filtered_data, key=lambda x: x.get('name', ''))

# Output the result to stdout
print(json.dumps(sorted_data))