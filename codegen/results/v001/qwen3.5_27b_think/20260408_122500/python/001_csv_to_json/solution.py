import csv
import json

def main():
    output = []
    try:
        with open('input/data.csv', 'r', newline='') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                entry = {
                    "Name": str(row.get('Name', '')),
                    "Age": int(row.get('Age', 0)),
                    "Email": str(row.get('Email', '')),
                    "Score": float(row.get('Score', 0.0))
                }
                output.append(entry)
        print(json.dumps(output))
    except FileNotFoundError:
        print(json.dumps([]))
    except Exception as e:
        print(json.dumps({"error": str(e)}))

if __name__ == "__main__":
    main()