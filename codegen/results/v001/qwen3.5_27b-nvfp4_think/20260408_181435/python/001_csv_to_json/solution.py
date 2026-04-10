{
  "script": "import csv\nimport json\n\nwith open('input/data.csv', newline='') as f:\n    reader = csv.DictReader(f)\n    data = []\n    for row in reader:\n        record = {\n            'Name': row['Name'],\n            'Age': int(row['Age']),\n            'Email': row['Email'],\n            'Score': float(row['Score'])\n        }\n        data.append(record)\n\nprint(json.dumps(data))"
}