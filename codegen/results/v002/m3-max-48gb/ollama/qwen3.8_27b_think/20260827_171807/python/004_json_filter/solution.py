import json

with open("input/data.json", "r") as f:
    data = json.load(f)

result = [
    record for record in data
    if record["active"] is True and record["age"] >= 30
]
result.sort(key=lambda r: r["name"])

print(json.dumps(result))