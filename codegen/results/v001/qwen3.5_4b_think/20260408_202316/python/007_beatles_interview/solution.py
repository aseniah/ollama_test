import json
from datetime import datetime

# Read expected format
with open('input/expected_format.json', 'r') as f:
    expected = json.load(f)

# Read input CSV
with open('input/input.csv', 'r') as f:
    reader = csv.DictReader(f)
    rows = list(reader)

# Process each row
output = []
birth_dates = []

for row in rows:
    birth_date = datetime.strptime(row['birth_date'], '%Y-%m-%d')
    birthday = datetime(2025, 7, 1)
    age = (birthday - birth_date).days / 365.25
    
    # Format: round to 1 decimal place
    age_str = f"{age:.1f}"
    
    # Create output object
    output_entry = {
        'name': row['name'],
        'age': age_str
    }
    output.append(output_entry)

# Write output
print(json.dumps(output))