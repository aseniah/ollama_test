import json
import os

def main():
    # Define the path to the input file
    input_path = 'input/data.json'
    
    # Check if the file exists to avoid runtime errors
    if not os.path.exists(input_path):
        return

    try:
        # Read the JSON data from the file
        with open(input_path, 'r') as f:
            data = json.load(f)
        
        # Filter the records:
        # 1. active must be true
        # 2. age must be 30 or older
        filtered_records = [
            record for record in data 
            if record.get('active') is True and record.get('age', 0) >= 30
        ]
        
        # Sort the filtered records by 'name' in ascending order
        sorted_records = sorted(filtered_records, key=lambda x: x['name'])
        
        # Output the resulting JSON array to stdout
        # Using indent=2 for readability, though standard json.dumps() is also acceptable
        print(json.dumps(sorted_records, indent=2))
        
    except (json.JSONDecodeError, KeyError, TypeError):
        # If there is an issue parsing the JSON or accessing keys, 
        # we exit silently to comply with "Do not output anything else"
        pass

if __name__ == "__main__":
    main()