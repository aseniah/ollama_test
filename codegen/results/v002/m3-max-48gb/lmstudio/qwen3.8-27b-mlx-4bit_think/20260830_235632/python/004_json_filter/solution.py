import json

with open("input/data.json") as f:
    data = json.load(f)

filtered = [record for record in data if record["active"] and record["age"] >= 30]
filtered.sort(key=lambda r: r["name"])

print(json.dumps(filtered, indent=2))