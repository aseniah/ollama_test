import json

# Read input file
with open("input/data.json", "r") as f:
    data = json.load(f)

# Filter and sort
filtered = [
    record for record in data
    if record["active"] and record["age"] >= 30
]

filtered.sort(key=lambda x: x["name"])

# Output as JSON array to stdout
print(json.dumps(filtered))