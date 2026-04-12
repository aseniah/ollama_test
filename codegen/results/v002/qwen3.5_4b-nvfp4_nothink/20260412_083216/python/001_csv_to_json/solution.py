import csv
import json

with open("input/data.csv", newline="", encoding="utf-8") as file:
    reader = csv.DictReader(file)
    data = []
    for row in reader:
        name = row["Name"]
        age = int(row["Age"])
        email = row["Email"]
        score = float(row["Score"])
        data.append({"name": name, "age": age, "email": email, "score": score})

print(json.dumps(data))