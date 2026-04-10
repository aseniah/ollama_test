import csv
import json

def csv_to_json(file_path):
    data = []
    with open(file_path, mode='r') as csv_file:
        csv_reader = csv.DictReader(csv_file)
        for row in csv_reader:
            name = row['Name']
            age = int(row['Age'])
            email = row['Email']
            score = float(row['Score'])
            data.append({
                'Name': name,
                'Age': age,
                'Email': email,
                'Score': score
            })
    return data

def main():
    csv_file_path = 'input/data.csv'
    json_output = csv_to_json(csv_file_path)
    print(json.dumps(json_output, indent=4))

if __name__ == "__main__":
    main()