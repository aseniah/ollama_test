import json

with open("input/data.json", "r") as f:
    data = json.load(f)

filtered = [rec for rec in data if rec["active"] is True and rec["age"] >= 30]
filtered.sort(key=lambda x: x["name"])

print(json.dumps(filtered))