import json

def solve():
    try:
        # Read the input file
        with open('input/data.json', 'r') as f:
            data = json.load(f)
        
        # Filter records: active is true and age is 30 or older
        filtered_data = [
            record for record in data 
            if record.get('active') is True and record.get('age', 0) >= 30
        ]
        
        # Sort records by name ascending
        sorted_data = sorted(filtered_data, key=lambda x: x['name'])
        
        # Output the result as a JSON array to stdout
        print(json.dumps(sorted_data))
        
    except (FileNotFoundError, json.JSONDecodeError):
        # Since the prompt requires no other output, 
        # we handle errors silently or as per standard Python behavior.
        pass

if __name__ == "__main__":
    solve()