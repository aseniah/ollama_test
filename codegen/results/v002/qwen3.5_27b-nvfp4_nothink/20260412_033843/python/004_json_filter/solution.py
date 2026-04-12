import json

def main():
    # Read and parse the JSON file
    with open('input/data.json', 'r') as f:
        data = json.load(f)

    # Filter records where active is true and age >= 30
    filtered_data = [
        record for record in data 
        if record.get("active") is True and record.get("age", 0) >= 30
    ]

    # Sort the filtered records by name ascending
    sorted_data = sorted(filtered_data, key=lambda x: x["name"])

    # Output the result as a JSON array to stdout
    print(json.dumps(sorted_data))

if __name__ == "__main__":
    main()