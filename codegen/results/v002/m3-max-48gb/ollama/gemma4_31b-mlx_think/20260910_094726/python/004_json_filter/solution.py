import json
import sys

def main():
    try:
        # Read the JSON file
        with open('input/data.json', 'r') as f:
            data = json.load(f)
        
        # Filter: active is true AND age is 30 or older
        filtered_data = [
            record for record in data 
            if record.get('active') is True and record.get('age', 0) >= 30
        ]
        
        # Sort by name ascending
        sorted_data = sorted(filtered_data, key=lambda x: x['name'])
        
        # Output to stdout as a JSON array
        json.dump(sorted_data, sys.stdout)
        print() # Add newline for clean output
        
    except FileNotFoundError:
        pass
    except json.JSONDecodeError:
        pass

if __name__ == "__main__":
    main()