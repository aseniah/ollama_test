import json
import sys

with open("input/data.json", encoding="utf-8") as f:
    data = json.load(f)

result = [
    record
    for record in data
    if record.get("active") is True and record.get("age", 0) >= 30
]

result.sort(key=lambda record: record["name"])

sys.stdout.write(json.dumps(result))