import json
import sys

def main():
    try:
        with open('input/data.json', 'r') as file:
            data = json.load(file)
        
        # Filter records where active is true and age is 30 or older
        filtered_records = [record for record in data 
                            if record.get('active') is True and record.get('age', 0) >= 30]
        
        # Sort by name ascending
        sorted_records = sorted(filtered_records, key=lambda x: x.get('name', ''))
        
        # Output to stdout as JSON array
        print(json.dumps(sorted_records))
    except FileNotFoundError:
        print("[]")
    except json.JSONDecodeError:
        print("[]")

if __name__ == "__main__":
    main()