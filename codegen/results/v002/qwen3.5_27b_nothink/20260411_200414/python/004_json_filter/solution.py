import json

def main():
    try:
        with open('input/data.json', 'r') as f:
            data = json.load(f)
    except (FileNotFoundError, json.JSONDecodeError) as e:
        # Exit silently on file issues as per "do not output anything else"
        return

    # Filter: active is true AND age >= 30
    filtered = [
        record for record in data
        if record.get('active') is True and record.get('age', 0) >= 30
    ]

    # Sort by name ascending
    sorted_data = sorted(filtered, key=lambda x: x['name'])

    # Output to stdout
    print(json.dumps(sorted_data))

if __name__ == "__main__":
    main()