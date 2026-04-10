import json

# Read the input file
with open('input/data.json', 'r') as file:
    data = json.load(file)

# Filter and sort the data
filtered_data = [record for record in data if record['active'] and record['age'] >= 30]
sorted_data = sorted(filtered_data, key=lambda x: x['name'])

# Output the result
print(json.dumps(sorted_data))