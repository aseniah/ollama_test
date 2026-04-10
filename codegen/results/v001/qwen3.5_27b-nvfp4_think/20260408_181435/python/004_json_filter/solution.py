import json

    with open('input/data.json') as f:
        data = json.load(f)

    filtered = [item for item in data if item.get('active') is True and item.get('age', 0) >= 30]
    sorted_data = sorted(filtered, key=lambda x: x['name'])

    print(json.dumps(sorted_data))