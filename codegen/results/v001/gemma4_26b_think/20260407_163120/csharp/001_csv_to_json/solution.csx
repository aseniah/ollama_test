using System;
using System.Collections.Generic;
using System.IO;
using System.Globalization;
using System.Text.Json;

string filePath = "input/data.csv";

if (!File.Exists(filePath))
{
    // If file is missing, output an empty array to maintain valid JSON requirement
    Console.WriteLine("[]");
}
else
{
    string[] lines = File.ReadAllLines(filePath);
    var jsonData = new List<Dictionary<string, object>>();

    // Skip the header row (index 0) and iterate through the data rows
    for (int i = 1; i < lines.Length; i++)
    {
        string line = lines[i];
        if (string.IsNullOrWhiteSpace(line)) continue;

        string[] parts = line.Split(',');

        // Ensure the row has enough columns as defined (Name, Age, Email, Score)
        if (parts.Length >= 4)
        {
            try
            {
                var record = new Dictionary<string, object>
                {
                    { "Name", parts[0].Trim() },
                    { "Age", int.Parse(parts[1].Trim(), CultureInfo.InvariantCulture) },
                    { "Email", parts[2].Trim() },
                    { "Score", float.Parse(parts[3].Trim(), CultureInfo.InvariantCulture) }
                };
                jsonData.Add(record);
            }
            catch (Exception)
            {
                // If a specific row is malformed, we skip it to ensure the 
                // final output remains a valid JSON array of the successfully parsed objects.
                continue;
            }
        }
    }

    // Configure JSON serialization to be pretty-printed
    var options = new JsonSerializerOptions 
    { 
        WriteIndented = true 
    };

    string jsonOutput = JsonSerializer.Serialize(jsonData, options);
    Console.WriteLine(jsonOutput);
}