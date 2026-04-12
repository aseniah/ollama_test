import json, sys

with open("input/data.json") as f:
    data = json.load(f)

result = sorted(
    [r for r in data if r["active"] and r["age"] >= 30],
    key=lambda r: r["name"]
)

print(json.dumps(result))
