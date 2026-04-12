using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Globalization;

// Define the path to the input file
string filePath = "input/data.csv";

// Check if the file exists before attempting to read
if (File.Exists(filePath))
{
    // Read all lines from the CSV file
    string[] lines = File.ReadAllLines(filePath);
    
    // Initialize a JsonArray to hold our objects
    JsonArray jsonArray = new JsonArray();

    // Skip the first line (header) and iterate through the remaining lines
    for (int i = 1; i < lines.Length; i++)
    {
        string line = lines[i];
        
        // Skip empty lines
        if (string.IsNullOrWhiteSpace(line))
        {
            continue;
        }

        // Split the line by comma
        string[] parts = line.Split(',');

        // Ensure we have enough columns to process the row
        if (parts.Length >= 4)
        {
            try
            {
                // Create a JsonObject for the current row
                // We use CultureInfo.InvariantCulture to ensure decimal points are parsed correctly 
                // regardless of the system's local culture settings.
                var obj = new JsonObject
                {
                    ["Name"] = parts[0].Trim(),
                    ["Age"] = int.Parse(parts[1].Trim(), CultureInfo.InvariantCulture),
                    ["Email"] = parts[2].Trim(),
                    ["Score"] = float.Parse(parts[3].Trim(), CultureInfo.InvariantCulture)
                };

                // Add the object to our array
                jsonArray.Add(obj);
            }
            catch (Exception)
            {
                // In a real-world scenario, we might log parsing errors.
                // For this task, we skip rows that don't match the expected format.
                continue;
            }
        }
    }

    // Serialize the JsonArray to a JSON string and output to stdout
    // We use the default settings which produces a compact JSON string.
    Console.WriteLine(jsonArray.ToJsonString());
}