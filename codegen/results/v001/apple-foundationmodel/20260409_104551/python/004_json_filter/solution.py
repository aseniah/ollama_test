import json

def filter_and_sort_records(file_path):
    try:
        # Read the JSON data from the file
        with open(file_path, 'r') as file:
            data = json.load(file)
        
        # Filter records where active is True and age is 30 or older
        filtered_data = [record for record in data if record['active'] and record['age'] >= 30]
        
        # Sort the filtered data by name in ascending order
        sorted_data = sorted(filtered_data, key=lambda x: x['name'])
        
        # Convert the sorted list back to a JSON array for output
        output_data = json.dumps(sorted_data, indent=2)
        
        # Print the output
        print(output_data)
    
    except FileNotFoundError:
        print(f"File not found: {file_path}")
    except json.JSONDecodeError:
        print("Error decoding JSON from file.")

# Call the function with the file path
filter_and_sort_records('input/data.json')