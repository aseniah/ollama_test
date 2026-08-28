using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Globalization;

// Define the structure for the output objects
// Using an anonymous type or a dictionary for simplicity in top-level script
// Since we need to output a JSON array of objects.

string filePath = "input/data.csv";

if (!File.Exists(filePath))
{
    // If file doesn't exist, we should output an empty array to remain valid JSON
    Console.WriteLine("[]");
    return;
}

string[] lines = File.ReadAllLines(filePath);

if (lines.Length <= 1)
{
    // Only header or empty file
    Console.WriteLine("[]");
    return;
}

var resultList = new List<object>();

// Skip the header row (index 0)
for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    string[] columns = line.Split(',');

    if (columns.Length >= 4)
    {
        try
        {
            // Parse values according to specified types
            string name = columns[0].Trim();
            int age = int.Parse(columns[1].Trim());
            string email = columns[2].Trim();
            // Use InvariantCulture to ensure '.' is treated as decimal separator regardless of local settings
            float score = float.Parse(columns[3].Trim(), CultureInfo.InvariantCulture);

            resultList.Add(new
            {
                Name = name,
                Age = age,
                Email = email,
                Score = score
            });
        }
        catch (Exception)
        {
            // In a production environment, we might log errors. 
            // For this task, we skip malformed rows to ensure valid JSON output.
            continue;
        }
    }
}

// Serialize the list to a JSON string
var options = new JsonSerializerOptions
{
    WriteIndented = true // Pretty print for readability, though not strictly required
};

string jsonOutput = JsonSerializer.Serialize(resultList, options);

// Write the final JSON array to stdout
Console.Write(jsonOutput);