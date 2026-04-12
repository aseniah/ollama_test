using System;
using System.IO;
using System.Collections.Generic;
using System.Globalization;
using System.Text.Json;

// Check if the input file exists
string filePath = "input/data.csv";

if (!File.Exists(filePath))
{
    // If the file doesn't exist, we output an empty array to maintain valid JSON output requirements
    Console.Write("[]");
    return;
}

// Read all lines from the CSV
string[] lines = File.ReadAllLines(filePath);

if (lines.Length <= 1)
{
    // Only header or empty file
    Console.Write("[]");
    return;
}

var resultList = new List<object>();

// Skip the header row (index 0)
for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    string[] columns = line.Split(',');

    // Ensure we have the required number of columns
    if (columns.Length >= 4)
    {
        try
        {
            // Parsing Name, Age, Email, Score
            string name = columns[0].Trim();
            int age = int.Parse(columns[1].Trim());
            string email = columns[2].Trim();
            float score = float.Parse(columns[3].Trim(), CultureInfo.InvariantCulture);

            // Create an anonymous object for the JSON structure
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
            // If a row is malformed, skip it to ensure valid JSON output for the rest
            continue;
        }
    }
}

// Serialize the list to a JSON array string
var options = new JsonSerializerOptions
{
    WriteIndented = true
};

string jsonOutput = JsonSerializer.Serialize(resultList, options);

// Output the final JSON to stdout
Console.Write(jsonOutput);