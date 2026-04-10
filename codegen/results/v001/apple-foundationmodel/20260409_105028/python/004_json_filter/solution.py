import json

def filter_and_sort_records(input_file='input/data.json', output_file='output.json'):
    # Read data from the input file
    with open(input_file, 'r') as file:
        data = json.load(file)

    # Filter records based on the criteria: active=True and age>=30
    filtered_records = [record for record in data if record['active'] and record['age'] >= 30]

    # Sort the filtered records by name in ascending order
    filtered_records.sort(key=lambda x: x['name'])

    # Convert the filtered and sorted list back to JSON format
    output_data = json.dumps(filtered_records, indent=4)

    # Output the result to stdout
    print(output_data)

# Call the function
filter_and_sort_records()