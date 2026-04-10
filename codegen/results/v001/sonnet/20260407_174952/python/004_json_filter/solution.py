import json
import sys

with open("input/data.json") as f:
    data = json.load(f)

filtered = [r for r in data if r.get("active") is True and r.get("age", 0) >= 30]
filtered.sort(key=lambda r: r.get("name", ""))

print(json.dumps(filtered))
