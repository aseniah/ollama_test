import json

def filter_and_sort_records(filename):
    with open(filename, 'r') as file:
        data = json.load(file)
    
    # Filter records where 'active' is True and 'age' is 30 or older
    filtered_records = [record for record in data if record['active'] and record['age'] >= 30]
    
    # Sort the filtered records by 'name' in ascending order
    sorted_records = sorted(filtered_records, key=lambda x: x['name'])
    
    # Convert the sorted list back to a JSON array and print it
    print(json.dumps(sorted_records, indent=2))

# Call the function with the input file path
filter_and_sort_records('input/data.json')