import json
import sys

def solve():
    try:
        # Read the input file
        with open('input/data.json', 'r') as f:
            data = json.load(f)
        
        # Filter: active is True AND age >= 30
        filtered_data = [
            record for record in data 
            if record.get('active') is True and record.get('age', 0) >= 30
        ]
        
        # Sort by name ascending
        sorted_data = sorted(filtered_data, key=lambda x: x.get('name', ''))
        
        # Output to stdout as JSON array
        print(json.dumps(sorted_data))
        
    except FileNotFoundError:
        pass
    except json.JSONDecodeError:
        pass

if __name__ == "__main__":
    solve()