python
import json

with open('input/data.json', 'r') as f:
    records = json.load(f)

# Filter and sort
filtered = [
    record for record in records
    if record.get('active', False) is True and record.get('age', 0) >= 30
]
filtered.sort(key=lambda x: x.get('name', ''))

print(json.dumps(filtered, ensure_ascii=False))
```