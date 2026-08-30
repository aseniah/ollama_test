import json

def main():
    # Read the input file
    try:
        with open('input/data.json', 'r') as f:
            data = json.load(f)
    except FileNotFoundError:
        return

    # Filter records: active is true AND age is 30 or older
    filtered_data = [
        record for record in data 
        if record.get('active') is True and record.get('age', 0) >= 30
    ]

    # Sort records by name ascending
    sorted_data = sorted(filtered_data, key=lambda x: x['name'])

    # Output result as a JSON array to stdout
    print(json.dumps(sorted_data))

if __name__ == "__main__":
    main()