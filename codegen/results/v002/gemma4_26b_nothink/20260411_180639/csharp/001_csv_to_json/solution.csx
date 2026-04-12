using System;
using System.IO;
using System.Collections.Generic;
using System.Globalization;
using System.Text.Json;

// Define the path to the input file
string filePath = "input/data.csv";

// Check if the file exists to prevent runtime errors
if (!File.Exists(filePath))
{
    return;
}

// List to hold the parsed objects
var results = new List<object>();

// Read all lines from the CSV
string[] lines = File.ReadAllLines(filePath);

// Ensure there is at least a header and one data row
if (lines.Length > 1)
{
    // Skip the first line (header) and iterate through the rest
    for (int i = 1; i < lines.Length; i++)
    {
        string line = lines[i];
        if (string.IsNullOrWhiteSpace(line)) continue;

        string[] columns = line.Split(',');

        // Ensure the row has the expected number of columns
        if (columns.Length >= 4)
        {
            try
            {
                // Parse the fields with appropriate types
                string name = columns[0].Trim();
                int age = int.Parse(columns[1].Trim());
                string email = columns[2].Trim();
                float score = float.Parse(columns[3].Trim(), CultureInfo.InvariantCulture);

                // Add an anonymous object to the list
                results.Add(new
                {
                    Name = name,
                    Age = age,
                    Email = email,
                    Score = score
                });
            }
            catch (Exception)
            {
                // In a real-world scenario, we might log parsing errors.
                // For this task, we skip malformed rows to maintain valid JSON output.
                continue;
            }
        }
    }
}

// Serialize the list to a JSON array string
var options = new JsonSerializerOptions
{
    WriteIndented = true
};

string jsonOutput = JsonSerializer.Serialize(results, options);

// Output the final JSON to stdout
Console.Write(jsonOutput);