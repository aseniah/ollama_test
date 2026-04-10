import json

# Read the input file
with open('input/data.json', 'r') as file:
    data = json.load(file)

# Filter and sort the data
filtered_data = [item for item in data if item['active'] and item['age'] >= 30]
sorted_data = sorted(filtered_data, key=lambda x: x['name'])

# Output the result to stdout
print(json.dumps(sorted_data))