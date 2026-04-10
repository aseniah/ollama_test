import json
import sys

def process_json():
    try:
        # Read the input file
        with open('input/data.json', 'r') as f:
            data = json.load(f)
        
        # Filter: active is True AND age >= 30
        # Sort: by name ascending
        filtered_data = [
            record for record in data 
            if record.get('active') is True and record.get('age', 0) >= 30
        ]
        
        filtered_data.sort(key=lambda x: x.get('name', ''))

        # Output the resulting JSON array to stdout
        json.dump(filtered_data, sys.stdout)
        
    except FileNotFoundError:
        pass
    except json.JSONDecodeError:
        pass
    except Exception:
        pass

if __name__ == "__main__":
    process_json()