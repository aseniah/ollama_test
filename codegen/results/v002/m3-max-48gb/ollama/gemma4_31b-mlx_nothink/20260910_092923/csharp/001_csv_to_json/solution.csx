using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Globalization;

string filePath = "input/data.csv";

if (!File.Exists(filePath))
{
    GC.Collect();
    return;
}

string[] lines = File.ReadAllLines(filePath);
if (lines.Length <= 1) 
{
    Console.WriteLine("[]");
    return;
}

var results = new List<object>();

// Skip header row
for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    string[] columns = line.Split(',');
    if (columns.Length < 4) continue;

    var entry = new
    {
        Name = columns[0].Trim(),
        Age = int.Parse(columns[1].Trim()),
        Email = columns[2].Trim(),
        Score = float.Parse(columns[3].Trim(), CultureInfo.InvariantCulture)
    };

    results.Add(entry);
}

var options = new JsonSerializerOptions 
{ 
    WriteIndented = true 
};

Console.WriteLine(JsonSerializer.Serialize(results, options));