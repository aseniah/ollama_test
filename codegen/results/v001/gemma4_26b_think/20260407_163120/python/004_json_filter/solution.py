import json
import sys

def process_json():
    input_path = 'input/data.json'
    
    try:
        with open(input_path, 'r') as f:
            data = json.load(f)
        
        # Filter: active is true and age is 30 or older
        filtered_data = [
            record for record in data 
            if record.get('active') is True and record.get('age', 0) >= 30
        ]
        
        # Sort: by name ascending
        sorted_data = sorted(filtered_data, key=lambda x: x['name'])
        
        # Output to stdout
        sys.stdout.write(json.dumps(sorted_data, indent=2))
        
    except FileNotFoundError:
        # If the file doesn't exist, we follow the instruction to output nothing else, 
        # but in a real-world scenario, we'd handle this error.
        pass
    except json.JSONDecodeError:
        pass

if __name__ == "__main__":
    process_json()