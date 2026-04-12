import json

def filter_and_sort_records():
    # Load the JSON data from the file
    with open('input/data.json', 'r') as file:
        data = json.load(file)

    # Filter records where 'active' is true and 'age' is 30 or older
    filtered_records = [record for record in data if record['active'] and record['age'] >= 30]

    # Sort the filtered records by 'name' in ascending order
    filtered_records.sort(key=lambda x: x['name'])

    # Print the filtered and sorted records
    for record in filtered_records:
        print(json.dumps(record, indent=2))

# Execute the function
filter_and_sort_records()