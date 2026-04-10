import json
import sys

def process_json():
    try:
        # Read the input file
        with open('input/data.json', 'r') as f:
            data = json.load(f)
        
        # Filter: active is true and age is 30 or older
        filtered_data = [
            record for record in data 
            if record.get('active') is True and record.get('age', 0) >= 30
        ]
        
        # Sort by name ascending
        sorted_data = sorted(filtered_data, key=lambda x: x.get('name', ''))
        
        # Output to stdout
        sys.stdout.write(json.dumps(sorted_data, indent=2) + '\n')
        
    except FileNotFoundError:
        # If file doesn't exist, we output nothing or handle as error.
        # Given the constraint "Do not output anything else", 
        # we assume the file exists per instructions.
        pass
    except json.JSONDecodeError:
        pass

if __name__ == "__main__":
    process_json()