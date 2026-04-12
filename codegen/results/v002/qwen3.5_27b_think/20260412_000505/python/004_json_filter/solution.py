import json

def main():
    # Read the JSON data from the specified file
    with open('input/data.json', 'r') as file:
        data = json.load(file)

    # Filter records where 'active' is True and 'age' is 30 or older
    filtered_records = [
        record for record in data 
        if record['active'] is True and record['age'] >= 30
    ]

    # Sort the filtered records by 'name' in ascending order
    sorted_records = sorted(filtered_records, key=lambda x: x['name'])

    # Output the result as a JSON array to stdout
    print(json.dumps(sorted_records))

if __name__ == "__main__":
    main()