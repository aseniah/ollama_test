import json

def main():
    # Read the input JSON file
    with open('input/data.json', 'r') as f:
        data = json.load(f)

    # Filter records where active is True and age is 30 or older
    filtered_data = [
        record for record in data
        if record.get('active') is True and record.get('age', 0) >= 30
    ]

    # Sort the filtered records by name ascending
    sorted_data = sorted(filtered_data, key=lambda x: x.get('name', ''))

    # Output the result as a JSON array
    print(json.dumps(sorted_data))

if __name__ == "__main__":
    main()