using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;

string filePath = "input/data.csv";

if (!File.Exists(filePath))
{
    return;
}

string[] lines = File.ReadAllLines(filePath);

// Skip header and process rows
var records = lines.Skip(1)
    .Where(line => !string.IsNullOrWhiteSpace(line))
    .Select(line =>
    {
        string[] columns = line.Split(',');
        return new
        {
            Name = columns[0].Trim(),
            Age = int.Parse(columns[1].Trim()),
            Email = columns[2].Trim(),
            Score = float.Parse(columns[3].Trim())
        };
    })
    .ToList();

// Serialize to JSON and write to stdout
string jsonOutput = JsonSerializer.Serialize(records, new JsonSerializerOptions 
{ 
    WriteIndented = true 
});

Console.WriteLine(jsonOutput);