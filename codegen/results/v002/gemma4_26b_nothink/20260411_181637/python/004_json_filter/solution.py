import json
import sys

def main():
    input_file = 'input/data.json'
    
    try:
        with open(input_file, 'r') as f:
            data = json.load(f)
            
        # Filter: active is true AND age is 30 or older
        filtered_data = [
            record for record in data 
            if record.get('active') is True and record.get('age', 0) >= 30
        ]
        
        # Sort by name ascending
        sorted_data = sorted(filtered_data, key=lambda x: x['name'])
        
        # Output to stdout
        print(json.dumps(sorted_data, indent=2))
        
    except FileNotFoundError:
        # If file doesn't exist, exit silently or handle as needed
        pass
    except json.JSONDecodeError:
        # If JSON is invalid, exit silently
        pass

if __name__ == "__main__":
    main()