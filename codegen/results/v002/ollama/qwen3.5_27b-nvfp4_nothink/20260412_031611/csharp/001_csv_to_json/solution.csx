#r "System.Text.Json"

using System;
using System.IO;
using System.Text.Json;
using System.Collections.Generic;

string filePath = "input/data.csv";
string[] lines = File.ReadAllLines(filePath);

// Skip the header (first line)
List<object> records = new List<object>();

for (int i = 1; i < lines.Length; i++)
{
    if (string.IsNullOrWhiteSpace(lines[i])) continue;

    string[] parts = lines[i].Split(',');

    // Create a dictionary to ensure property order and handle serialization explicitly
    var record = new Dictionary<string, object>
    {
        ["Name"] = parts[0],
        ["Age"] = int.Parse(parts[1]),
        ["Email"] = parts[2],
        ["Score"] = float.Parse(parts[3])
    };

    records.Add(record);
}

// Serialize to JSON with indentation for readability (optional, but standard for arrays)
var jsonOptions = new JsonSerializerOptions
{
    WriteIndented = false, // Compact JSON as requested for stdout
};

string jsonOutput = JsonSerializer.Serialize(records, jsonOptions);

Console.WriteLine(jsonOutput);