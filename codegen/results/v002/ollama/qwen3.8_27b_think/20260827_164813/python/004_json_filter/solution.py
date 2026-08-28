import json

with open("input/data.json", "r") as f:
    data = json.load(f)

filtered = [r for r in data if r["active"] is True and r["age"] >= 30]
filtered.sort(key=lambda r: r["name"])

print(json.dumps(filtered, indent=2))