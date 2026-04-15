import json

def solve():
    # Read the input file
    try:
        with open('input/data.json', 'r') as f:
            data = json.load(f)
    except FileNotFoundError:
        return

    # Filter records: active is true AND age is 30 or older
    filtered_records = [
        record for record in data 
        if record.get('active') is True and record.get('age', 0) >= 30
    ]

    # Sort records by name ascending
    sorted_records = sorted(filtered_records, key=lambda x: x['name'])

    # Output the resulting JSON array to stdout
    print(json.dumps(sorted_records))

if __name__ == "__main__":
    solve()