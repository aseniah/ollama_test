using System;
using System.Collections.Generic;
using System.IO;
using System.Globalization;
using System.Text.Json;

string filePath = "input/data.csv";

if (!File.Exists(filePath))
{
    return;
}

string[] lines = File.ReadAllLines(filePath);
if (lines.Length <= 1)
{
    Console.WriteLine("[]");
    return;
}

var records = new List<object>();

// Skip the header row
for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    string[] parts = line.Split(',');
    if (parts.Length >= 4)
    {
        records.Add(new
        {
            Name = parts[0],
            Age = int.Parse(parts[1]),
            Email = parts[2],
            Score = float.Parse(parts[3], CultureInfo.InvariantCulture)
        });
    }
}

string jsonOutput = JsonSerializer.Serialize(records, new JsonSerializerOptions 
{ 
    WriteIndented = false 
});

Console.WriteLine(jsonOutput);