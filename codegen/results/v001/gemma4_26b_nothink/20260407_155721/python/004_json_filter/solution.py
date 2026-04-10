import json
import sys

def process_json():
    try:
        # Load the data from the specified path
        with open('input/data.json', 'r') as f:
            data = json.load(f)
        
        # Filter: active is true AND age >= 30
        filtered_data = [
            record for record in data 
            if record.get('active') is True and record.get('age', 0) >= 30
        ]
        
        # Sort by name ascending
        sorted_data = sorted(filtered_data, key=lambda x: x.get('name', ''))
        
        # Output to stdout as a JSON array
        json.dump(sorted_data, sys.stdout)
        
    except FileNotFoundError:
        # If file is missing, output an empty list as per standard filtering logic
        # or handle as error. Requirement implies reading exists.
        pass
    except json.JSONDecodeError:
        pass

if __name__ == "__main__":
    process_json()