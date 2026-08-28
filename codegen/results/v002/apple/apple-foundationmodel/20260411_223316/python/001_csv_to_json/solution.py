import csv
import json

def read_csv_to_json(csv_file_path):
    data = []
    with open(csv_file_path, mode='r') as file:
        csv_reader = csv.DictReader(file)
        for row in csv_reader:
            # Create a dictionary with the fields mapped to their values
            row_dict = {
                "Name": row["Name"],
                "Age": int(row["Age"]),
                "Email": row["Email"],
                "Score": float(row["Score"])
            }
            # Append the dictionary to the data list
            data.append(row_dict)
    return data

def main():
    # Specify the path to the CSV file
    csv_file_path = 'input/data.csv'
    
    # Read CSV data into a list of dictionaries
    csv_data = read_csv_to_json(csv_file_path)
    
    # Convert the list of dictionaries to a JSON array
    json_data = json.dumps(csv_data, indent=4)
    
    # Print the JSON array to stdout
    print(json_data)

if __name__ == "__main__":
    main()