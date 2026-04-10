import json

# Read the JSON data from the file
with open("input/data.json", "r") as f:
    data = json.load(f)

# Filter records where active is true and age is 30 or older
filtered_data = [
    record for record in data
    if record["active"] is True and record["age"] >= 30
]

# Sort the filtered records by name in ascending order
filtered_data.sort(key=lambda x: x["name"])

# Output the result as a JSON array to stdout
print(json.dumps(filtered_data))