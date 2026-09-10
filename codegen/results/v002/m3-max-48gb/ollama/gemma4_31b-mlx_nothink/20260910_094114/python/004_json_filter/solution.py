import json
import sys

def main():
    try:
        # Read the file input/data.json
        with open('input/data.json', 'r') as f:
            data = json.load(f)
        
        # Filter records: active is True AND age >= 30
        filtered_data = [
            item for item in data 
            if item.get('active') is True and item.get('age', 0) >= 30
        ]
        
        # Sort by name ascending
        sorted_data = sorted(filtered_data, key=lambda x: x['name'])
        
        # Output to stdout as a JSON array
        print(json.dumps(sorted_data))
        
    except FileNotFoundError:
        # Depending on requirements, you could print an error or just an empty array
        pass
    except json.JSONDecodeError:
        pass

if __name__ == "__main__":
    main()