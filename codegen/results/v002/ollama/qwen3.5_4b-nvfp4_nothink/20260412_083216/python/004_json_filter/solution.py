import json

# Read the input data
with open("input/data.json", "r") as f:
    data = json.load(f)

# Filter records where active is True and age >= 30, then sort by name
filtered_records = [
    record for record in data 
    if record["active"] and record["age"] >= 30
]

filtered_records.sort(key=lambda x: x["name"])

# Output the filtered records as JSON
print(json.dumps(filtered_records))