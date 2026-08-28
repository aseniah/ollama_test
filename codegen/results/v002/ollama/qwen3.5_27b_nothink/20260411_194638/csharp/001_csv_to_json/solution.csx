#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Serialization;

string filePath = "input/data.csv";
string[] lines = File.ReadAllLines(filePath);

var results = new List<Dictionary<string, object>>();

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    // Split only on commas, limiting to 4 parts to handle potential edge cases, 
    // though the provided data is simple.
    string[] parts = line.Split(',');

    if (parts.Length >= 4)
    {
        var item = new Dictionary<string, object>
        {
            ["Name"] = parts[0],
            ["Age"] = int.Parse(parts[1]),
            ["Email"] = parts[2],
            ["Score"] = float.Parse(parts[3])
        };
        results.Add(item);
    }
}

// Configure options to output the array directly
var options = new JsonSerializerOptions
{
    WriteIndented = true,
    PropertyNamingPolicy = JsonNamingPolicy.CamelCase // Optional, usually expected in JSON APIs, but TitleCase used in prompt description.
                                                        // The prompt asks for specific keys: Name, Age, Email, Score.
                                                        // By default, C# properties are TitleCase in the dict, so we keep default naming policy.
    PropertyNamingPolicy = null 
};

string jsonOutput = JsonSerializer.Serialize(results, options);
Console.WriteLine(jsonOutput);