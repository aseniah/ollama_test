import json
import sys

def main():
    input_file = "input/data.json"
    
    try:
        with open(input_file, 'r', encoding='utf-8') as f:
            data = json.load(f)
    except FileNotFoundError:
        sys.exit(1)
    except json.JSONDecodeError:
        sys.exit(1)

    # Filter records where active is True and age >= 30
    filtered = [
        record for record in data
        if record.get("active") is True and record.get("age", 0) >= 30
    ]

    # Sort by name ascending
    filtered.sort(key=lambda x: x.get("name", ""))

    # Output the result as a JSON array to stdout
    print(json.dumps(filtered))

if __name__ == "__main__":
    main()