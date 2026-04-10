import pandas as pd
import json
from datetime import datetime

# Load the CSV file
file_path = 'input/input.csv'
df = pd.read_csv(file_path)

# Function to calculate age as of July 1, 2025
def calculate_age(birth_date):
    today = datetime(2025, 7, 1)
    birth_date = datetime.strptime(birth_date, "%Y-%m-%d")
    return today.year - birth_date.year - ((today.month, today.day) < (birth_date.month, birth_date.day))

# Calculate ages for each person
df['age'] = df['birthdate'].apply(calculate_age)

# Convert DataFrame to a list of dictionaries
result = df.to_dict(orient='records')

# Output the JSON array
print(json.dumps(result, indent=4))