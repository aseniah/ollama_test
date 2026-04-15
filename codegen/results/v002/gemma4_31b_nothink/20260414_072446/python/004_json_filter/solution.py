import json
import sys

def solve():
    try:
        # Read the input file
        with open('input/data.json', 'r') as f:
            data = json.load(f)
        
        # Filter records: active is true AND age >= 30
        filtered_data = [
            record for record in data 
            if record.get('active') is True and record.get('age', 0) >= 30
        ]
        
        # Sort by name ascending
        sorted_data = sorted(filtered_data, key=lambda x: x['name'])
        
        # Output as JSON array to stdout
        json.dump(sorted_data, sys.stdout)
        
    except (FileNotFoundError, json.JSONDecodeError):
        # Ensuring nothing else is outputted unless it's the result
        pass

if __name__ == "__main__":
    solve()