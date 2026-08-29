import json
import sys

with open("input/data.json", "r", encoding="utf-8") as f:
    records = json.load(f)

selected = [
    record
    for record in records
    if record.get("active") is True and record.get("age", 0) >= 30
]

selected.sort(key=lambda record: record.get("name", ""))
sys.stdout.write(json.dumps(selected))