import json

def solve():
    try:
        # Read the input file
        with open('input/data.json', 'r') as f:
            data = json.load(f)
        
        # Filter: active must be True and age must be 30 or older
        filtered_data = [
            record for record in data 
            if record.get('active') is True and record.get('age', 0) >= 30
        ]
        
        # Sort: by name ascending
        sorted_data = sorted(filtered_data, key=lambda x: x['name'])
        
        # Output to stdout as JSON array
        print(json.dumps(sorted_data))
        
    except (FileNotFoundError, json.JSONDecodeError):
        # In a production environment, handle errors accordingly. 
        # Here we maintain the constraint of outputting only the result.
        pass

if __name__ == "__main__":
    solve()