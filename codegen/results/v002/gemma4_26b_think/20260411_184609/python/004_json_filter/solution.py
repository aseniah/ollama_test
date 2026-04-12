import json
import os

def process_json_data(file_path):
    # Check if the file exists to avoid FileNotFoundError
    if not os.path.exists(file_path):
        return

    try:
        # Read the JSON file
        with open(file_path, 'r') as f:
            data = json.load(f)

        # Filter records: active must be true and age must be 30 or older
        filtered_data = [
            record for record in data 
            if record.get('active') is True and record.get('age', 0) >= 30
        ]

        # Sort the filtered records by name in ascending order
        sorted_data = sorted(filtered_data, key=lambda x: x['name'])

        # Output the result as a JSON array to stdout
        print(json.dumps(sorted_data, indent=2))

    except (json.JSONDecodeError, KeyError, TypeError) as e:
        # In a production environment, you might want to log errors.
        # For this task, we follow the instruction to output only the JSON array.
        pass

if __name__ == "__main__":
    # Path to the input file as specified in the prompt
    input_file_path = 'input/data.json'
    process_json_data(input_file_path)