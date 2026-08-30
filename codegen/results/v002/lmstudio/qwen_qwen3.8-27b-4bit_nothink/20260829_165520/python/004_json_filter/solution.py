import json

# Read the input file
with open("input/data.json", "r") as f:
    data = json.load(f)

# Filter: active is true and age is 30 or older
filtered = [record for record in data if record["active"] and record["age"] >= 30]

# Sort by name ascending
filtered.sort(key=lambda x: x["name"])

# Output to stdout
print(json.dumps(filtered))