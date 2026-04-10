import json
import sys

# Read input data
with open('input/data.json', 'r') as f:
    data = json.load(f)

# Filter and sort
result = [
    record for record in data 
    if record.get('active') is True and record.get('age', 0) >= 30
]
result.sort(key=lambda x: x.get('name', ''))

# Output
print(json.dumps(result))