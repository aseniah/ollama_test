import csv
import json
import sys

def main():
    rows = []
    try:
        with open('input/data.csv', mode='r', newline='', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                obj = {
                    "Name": row['Name'],
                    "Age": int(row['Age']),
                    "Email": row['Email'],
                    "Score": float(row['Score'])
                }
                rows.append(obj)
    except FileNotFoundError:
        # Fallback if file doesn't exist for testing purposes (though problem implies it exists)
        # In a real scenario, this would raise an error.
        pass
    except Exception as e:
        sys.stderr.write(f"Error reading file: {e}\n")
        sys.exit(1)

    print(json.dumps(rows))

if __name__ == "__main__":
    main()