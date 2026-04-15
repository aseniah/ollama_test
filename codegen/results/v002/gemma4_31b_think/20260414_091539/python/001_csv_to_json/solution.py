import csv
import json
import sys

def solve():
    file_path = 'input/data.csv'
    data_list = []

    try:
        with open(file_path, mode='r', encoding='utf-8') as csvfile:
            # DictReader uses the first row as keys for each row dictionary
            reader = csv.DictReader(csvfile)
            
            for row in reader:
                # Construct the object with the specific types requested
                record = {
                    "Name": str(row["Name"]),
                    "Age": int(row["Age"]),
                    "Email": str(row["Email"]),
                    "Score": float(row["Score"])
                }
                data_list.append(record)
                
        # Output the resulting list as a valid JSON array to stdout
        print(json.dumps(data_list))
        
    except FileNotFoundError:
        # In a real scenario, you might handle this differently, 
        # but for this task we assume the file exists.
        sys.exit(1)
    except Exception as e:
        sys.exit(1)

if __name__ == "__main__":
    solve()