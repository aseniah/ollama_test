using System;
using System.IO;
using System.Collections.Generic;
using System.Globalization;
using System.Text.Json;

// Define the path to the input file
string filePath = "input/data.csv";

// Ensure the script only outputs valid JSON. 
// If the file doesn't exist, we output an empty array.
if (!File.Exists(filePath))
{
    Console.WriteLine("[]");
    return;
}

try
{
    // Read all lines from the CSV file
    string[] lines = File.ReadAllLines(filePath);
    
    // Prepare a list to hold our objects
    var jsonData = new List<object>();

    // The first line is the header, so we start iterating from index 1
    for (int i = 1; i < lines.Length; i++)
    {
        string line = lines[i];
        
        // Skip empty lines
        if (string.IsNullOrWhiteSpace(line))
        {
            continue;
        }

        // Split the CSV line by comma
        string[] columns = line.Split(',');

        // Ensure we have enough columns to populate the required fields
        if (columns.Length >= 4)
        {
            // Parse fields with InvariantCulture to ensure consistent numeric parsing
            string name = columns[0].Trim();
            int age = int.Parse(columns[1].Trim(), CultureInfo.InvariantCulture);
            string email = columns[2].Trim();
            float score = float.Parse(columns[3].Trim(), CultureInfo.InvariantCulture);

            // Add an anonymous object to the list
            jsonData.Add(new
            {
                Name = name,
                Age = age,
                Email = email,
                Score = score
            });
        }
    }

    // Configure JSON serialization options for pretty printing
    var options = new JsonSerializerOptions
    {
        WriteIndented = true
    };

    // Output the serialized JSON array to stdout
    Console.WriteLine(JsonSerializer.Serialize(jsonData, options));
}
catch (Exception)
{
    // In case of parsing errors or IO errors, output an empty array to maintain valid JSON output
    Console.WriteLine("[]");
}