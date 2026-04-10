import json

with open("input/data.json", "r") as f:
    data = json.load(f)

filtered = [record for record in data if record.get("active") and record.get("age", 0) >= 30]
filtered.sort(key=lambda x: x["name"])

print(json.dumps(filtered, ensure_ascii=False))