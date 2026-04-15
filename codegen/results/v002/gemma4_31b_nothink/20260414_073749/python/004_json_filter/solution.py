import json
import sys

def solve():
    try:
        # Open and read the input file
        with open('input/data.json', 'r') as f:
            data = json.load(f)
        
        # Filter records: active must be True and age must be 30 or older
        filtered_data = [
            item for item in data 
            if item.get('active') is True and item.get('age', 0) >= 30
        ]
        
        # Sort the filtered list by the 'name' field in ascending order
        sorted_data = sorted(filtered_data, key=lambda x: x['name'])
        
        # Output the result to stdout as a JSON array
        print(json.dumps(sorted_data))
        
    except FileNotFoundError:
        # Handle case where file does not exist silently or as needed
        pass
    except json.JSONDecodeError:
        # Handle invalid JSON
        pass

if __name__ == "__main__":
    solve()