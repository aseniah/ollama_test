import json
import sys

try:
    with open('input/data.json', 'r') as f:
        data = json.load(f)
except FileNotFoundError:
    # If file doesn't exist, output an empty list as per request logic (no records to process)
    data = []
except json.JSONDecodeError:
    sys.exit(1)

# Filter and sort
filtered_data = [record for record in data if record.get("active") is True and record.get("age", 0) >= 30]
sorted_data = sorted(filtered_data, key=lambda x: x.get("name", ""))

# Output to stdout
print(json.dumps(sorted_data))