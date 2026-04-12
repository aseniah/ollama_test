import json

def filter_and_sort_records(json_data):
    # Parse the JSON data
    data = json.loads(json_data)
    
    # Filter records where 'active' is True and 'age' is 30 or older
    filtered_records = [record for record in data if record['active'] and record['age'] >= 30]
    
    # Sort the filtered records by 'name' in ascending order
    sorted_records = sorted(filtered_records, key=lambda x: x['name'])
    
    # Prepare the output in JSON format
    output_json = json.dumps(sorted_records, indent=2)
    
    # Print the output to stdout
    print(output_json)

# Read the JSON data from the file
with open('input/data.json', 'r') as file:
    json_data = file.read()

filter_and_sort_records(json_data)