using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Globalization;

if (!File.Exists("input/data.csv"))
{
    Environment.Exit(1);
}

string[] lines = File.ReadAllLines("input/data.csv");
if (lines.Length <= 1) 
{
    Console.WriteLine("[]");
    return;
}

var results = new List<object>();

// Skip header row
for (int i = 1; i < lines.Length; i++)
{
    if (string.IsNullOrWhiteSpace(lines[i])) continue;

    string[] parts = lines[i].Split(',');
    if (parts.Length < 4) continue;

    results.Add(new
    {
        Name = parts[0].Trim(),
        Age = int.Parse(parts[1].Trim()),
        Email = parts[2].Trim(),
        Score = float.Parse(parts[3].Trim(), CultureInfo.InvariantCulture)
    });
}

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(results, options));