import json

# Read the input JSON
with open("input/data.json", "r") as f:
    data = json.load(f)

# Filter and sort
result = [record for record in data if record["active"] and record["age"] >= 30]
result.sort(key=lambda x: x["name"])

# Output JSON array to stdout
print(json.dumps(result))