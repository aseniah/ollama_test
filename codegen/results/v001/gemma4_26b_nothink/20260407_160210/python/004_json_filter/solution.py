import json
import sys

def process_json():
    try:
        # Load the input file
        with open('input/data.json', 'r') as f:
            data = json.load(f)
        
        # Filter: active is true AND age >= 30
        # Sort: name ascending
        filtered_data = [
            record for record in data 
            if record.get('active') is True and record.get('age', 0) >= 30
        ]
        
        sorted_data = sorted(filtered_data, key=lambda x: x.get('name', ''))
        
        # Output the result to stdout
        sys.stdout.write(json.dumps(sorted_data, indent=2))
        
    except FileNotFoundError:
        pass
    except json.JSONDecodeError:
        pass
    except Exception:
        pass

if __name__ == "__main__":
    process_json()