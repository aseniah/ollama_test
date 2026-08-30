using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Text.Json;

// Define the path to the input file
string filePath = "input/data.csv";

// Check if file exists to prevent runtime errors
if (!File.Exists(filePath))
{
    return;
}

// Read all lines from the CSV file
string[] lines = File.ReadAllLines(filePath);

// List to hold our parsed objects
var dataList = new List<object>();

// Iterate through the lines, skipping the header (index 0)
for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    // Split the CSV row by comma
    string[] columns = line.Split(',');

    // Ensure we have the correct number of columns before parsing
    if (columns.Length >= 4)
    {
        try
        {
            // Parse fields according to the required schema
            string name = columns[0].Trim();
            int age = int.Parse(columns[1].Trim());
            string email = columns[2].Trim();
            // Use InvariantCulture to ensure '.' is parsed correctly as a decimal separator
            float score = float.Parse(columns[3].Trim(), CultureInfo.InvariantCulture);

            // Create an anonymous object for the current row
            dataList.Add(new
            {
                Name = name,
                Age = age,
                Email = email,
                Score = score
            });
        }
        catch (Exception)
        {
            // In a real-world scenario, you might log parsing errors.
            // For this task, we skip malformed rows.
            continue;
        }
    }
}

// Serialize the list to a JSON array string
// Using WriteIndented = true for readable output, though not strictly required by the prompt
var options = new JsonSerializerOptions
{
    WriteIndented = true
};

string jsonOutput = JsonSerializer.Serialize(dataList, options);

// Output the JSON to stdout
Console.WriteLine(jsonOutput);