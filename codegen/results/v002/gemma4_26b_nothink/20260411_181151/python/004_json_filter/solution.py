import json
import sys

def main():
    input_file = 'input/data.json'
    
    try:
        with open(input_file, 'r') as f:
            data = json.load(f)
        
        # Filter: active is true AND age >= 30
        filtered_data = [
            record for record in data 
            if record.get('active') is True and record.get('age', 0) >= 30
        ]
        
        # Sort: by name ascending
        sorted_data = sorted(filtered_data, key=lambda x: x['name'])
        
        # Output to stdout as JSON array
        print(json.dumps(sorted_data, indent=2))
        
    except FileNotFoundError:
        pass
    except json.JSONDecodeError:
        pass
    except Exception:
        pass

if __name__ == "__main__":
    main()