import json
import sys

def main():
    try:
        # Open and read the JSON file
        with open('input/data.json', 'r') as f:
            data = json.load(f)
        
        # Filter: active is true AND age >= 30
        filtered_data = [
            record for record in data 
            if record.get('active') is True and record.get('age', 0) >= 30
        ]
        
        # Sort: name ascending
        sorted_data = sorted(filtered_data, key=lambda x: x['name'])
        
        # Output the resulting array as JSON to stdout
        print(json.dumps(sorted_data))
        
    except (FileNotFoundError, json.JSONDecodeError):
        # In a production environment, handle errors appropriately.
        # For this task, we ensure we only output the final JSON result.
        pass

if __name__ == "__main__":
    main()