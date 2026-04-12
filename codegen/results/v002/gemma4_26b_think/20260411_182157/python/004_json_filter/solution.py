import json
import sys

def solve():
    """
    Reads input/data.json, filters records where active is True and age >= 30,
    sorts them by name ascending, and prints the resulting JSON array to stdout.
    """
    try:
        # Define the path to the input file
        file_path = 'input/data.json'
        
        # Read the JSON data from the file
        with open(file_path, 'r') as f:
            data = json.load(f)
            
        # Filter the records: active must be true and age must be 30 or older
        filtered_records = [
            record for record in data 
            if record.get('active') is True and record.get('age', 0) >= 30
        ]
        
        # Sort the filtered records by the 'name' field in ascending order
        sorted_records = sorted(filtered_records, key=lambda x: x['name'])
        
        # Output the resulting JSON array to stdout
        # indent=2 is used for readability, though not strictly required by the prompt
        print(json.dumps(sorted_records, indent=2))
        
    except FileNotFoundError:
        # If the file is not found, we exit silently to avoid extra output to stdout
        pass
    except json.JSONDecodeError:
        # If JSON is malformed, we exit silently
        pass
    except Exception:
        # For any other unexpected error, we exit silently
        pass

if __name__ == "__main__":
    solve()