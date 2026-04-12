import json

def filter_and_sort_records():
    # Load data from the JSON file
    with open('input/data.json', 'r') as file:
        data = json.load(file)
    
    # Filter records based on the criteria
    filtered_records = [record for record in data if record['active'] and record['age'] >= 30]
    
    # Sort by name in ascending order
    filtered_records.sort(key=lambda x: x['name'])
    
    # Print the filtered and sorted records as JSON
    print(json.dumps(filtered_records, indent=2))

if __name__ == "__main__":
    filter_and_sort_records()