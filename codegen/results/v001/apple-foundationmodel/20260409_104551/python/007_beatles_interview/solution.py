import pandas as pd
import json

def calculate_ages(birthdates):
    """
    Calculate the age of each entry as of July 1, 2025.

    Parameters:
    birthdates (list): List of birthdates.

    Returns:
    list: List of ages.
    """
    return [2025 - pd.to_datetime(birthdate).year]

def read_csv_and_calculate_ages(csv_file):
    """
    Read a CSV file, calculate ages, and return as JSON array.

    Parameters:
    csv_file (str): Path to the CSV file.

    Returns:
    list: JSON array of ages.
    """
    # Read the CSV file
    df = pd.read_csv(csv_file)
    
    # Assuming the CSV has columns 'birthdate' and 'age'
    if 'birthdate' in df.columns and 'age' in df.columns:
        # Calculate ages
        df['age'] = calculate_ages(df['birthdate'])
    else:
        raise ValueError("CSV file must contain 'birthdate' and 'age' columns.")

    # Convert to JSON array
    return df.to_json(orient='records')

if __name__ == "__main__":
    csv_file = "input/input.csv"
    json_output = read_csv_and_calculate_ages(csv_file)
    print(json_output)