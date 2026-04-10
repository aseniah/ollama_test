using System;
using System.Collections.Generic;
using System.IO;
using System.Globalization;
using System.Text.Json;

// Define the path to the input file
string filePath = "input/data.csv";

// Check if file exists to avoid runtime errors
if (!File.Exists(filePath))
{
    // Output empty array if file is missing to maintain valid JSON output requirement
    Console.Write("[]");
    return;
}

var results = new List<Dictionary<string, object>>();

try
{
    // Read all lines from the CSV
    string[] lines = File.ReadAllLines(filePath);

    if (lines.Length > 1)
    {
        // The first line is the header. We skip it to process data rows.
        // Expected format: Name,Age,Email,Score
        for (int i = 1; i < lines.Length; i++)
        {
            string line = lines[i];
            if (string.IsNullOrWhiteSpace(line)) continue;

            string[] columns = line.Split(',');

            // Ensure we have enough columns to prevent IndexOutOfRangeException
            if (columns.Length >= 4)
            {
                var rowObject = new Dictionary<string, object>
                {
                    { "Name", columns[0].Trim() },
                    { "Age", int.Parse(columns[1].Trim()) },
                    { "Email", columns[2].Trim() },
                    { "Score", float.Parse(columns[3].Trim(), CultureInfo.InvariantCulture) }
                };
                results.Add(rowObject);
            }
        }
    }
}
catch (Exception)
{
    // In case of parsing errors, we output an empty array to ensure 
    // we don't break the "Output must be valid JSON" constraint.
    Console.Write("[]");
    return;
}

// Serialize the list of objects to a JSON array string
string jsonOutput = JsonSerializer.Serialize(results, new JsonSerializerOptions 
{ 
    WriteIndented = false 
});

// Output the final JSON to stdout
Console.Write(jsonOutput);