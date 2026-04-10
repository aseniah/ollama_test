import json
import sys

def process_json():
    file_path = 'input/data.json'
    
    try:
        with open(file_path, 'r') as f:
            data = json.load(f)
        
        # Filter: active is True and age is 30 or older
        filtered_data = [
            record for record in data 
            if record.get('active') is True and record.get('age', 0) >= 30
        ]
        
        # Sort: by name ascending
        sorted_data = sorted(filtered_data, key=lambda x: x.get('name', ''))
        
        # Output to stdout
        sys.stdout.write(json.dumps(sorted_data, indent=4) + '\n')
        
    except FileNotFoundError:
        # If the file doesn't exist, we output an empty array or handle as per requirements
        # Since the prompt asks for a specific output, we assume the file exists.
        pass
    except json.JSONDecodeError:
        pass

if __name__ == "__main__":
    process_json()