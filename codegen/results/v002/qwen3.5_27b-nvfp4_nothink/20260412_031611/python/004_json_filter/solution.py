import json

# Read the input JSON file
with open("input/data.json", "r") as file:
    data = json.load(file)

# Filter records where active is True and age is 30 or older
filtered_data = [
    record for record in data
    if record["active"] is True and record["age"] >= 30
]

# Sort the filtered records by name ascending
sorted_data = sorted(filtered_data, key=lambda x: x["name"])

# Output the result as a JSON array to stdout
print(json.dumps(sorted_data))