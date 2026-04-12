using System;
using System.IO;
using System.Collections.Generic;
using System.Text.Json;
using System.Globalization;

// Path to the input file
string filePath = "input/data.csv";

// Check if the file exists before attempting to read
if (File.Exists(filePath))
{
    // Read all lines from the CSV file
    string[] lines = File.ReadAllLines(filePath);
    
    // This list will hold our objects as dictionaries to easily represent JSON objects
    var result = new List<Dictionary<string, object>>();

    // Iterate through the lines, skipping the header row (index 0)
    for (int i = 1; i < lines.Length; i++)
    {
        string line = lines[i];
        
        // Skip empty lines
        if (string.IsNullOrWhiteSpace(line))
        {
            continue;
        }

        // Split the CSV line by comma
        string[] parts = line.Split(',');

        // Ensure we have enough columns to parse the required fields
        if (parts.Length >= 4)
        {
            try
            {
                // Create a dictionary for the current row
                // We use Trim() to handle any potential whitespace around the values
                var entry = new Dictionary<string, object>
                {
                    { "Name", parts[0].Trim() },
                    { "Age", int.Parse(parts[1].Trim()) },
                    { "Email", parts[2].Trim() },
                    { "Score", float.Parse(parts[3].Trim(), CultureInfo.InvariantCulture) }
                };
                
                result.Add(entry);
            }
            catch (Exception)
            {
                // In case of a parsing error (e.g. bad number format), we skip this row
                // to ensure the final JSON output remains valid.
                continue;
            }
        }
    }

    // Configure JSON serialization options for pretty printing
    var options = new JsonSerializerOptions
    {
        WriteIndented = true
    };

    // Serialize the list of dictionaries into a JSON array string
    string jsonOutput = JsonSerializer.Serialize(result, options);

    // Output the JSON array to stdout
    Console.Write(jsonOutput);
}