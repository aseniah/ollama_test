import json
import sys

try:
    with open("input/data.json", "r") as file:
        data = json.load(file)

    # Filter records where active is True and age is 30 or older
    filtered_data = [record for record in data if record["active"] is True and record["age"] >= 30]

    # Sort the filtered records by name in ascending order
    sorted_data = sorted(filtered_data, key=lambda x: x["name"])

    # Output the result as a JSON array to stdout
    print(json.dumps(sorted_data))

except FileNotFoundError:
    sys.stderr.write("Error: input/data.json not found.\n")
    sys.exit(1)
except json.JSONDecodeError:
    sys.stderr.write("Error: input/data.json is not valid JSON.\n")
    sys.exit(1)
except (KeyError, TypeError) as e:
    sys.stderr.write(f"Error: Unexpected data format. {e}\n")
    sys.exit(1)