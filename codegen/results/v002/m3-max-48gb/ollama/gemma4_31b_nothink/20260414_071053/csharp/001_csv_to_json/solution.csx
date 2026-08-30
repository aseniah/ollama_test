using System;
using System.IO;
using System.Collections.Generic;
using System.Text.Json;
using System.Globalization;

if (!File.Exists("input/data.csv"))
{
    Environment.Exit(1);
}

string[] lines = File.ReadAllLines("input/data.csv");

// Skip header row
var dataObjects = new List<object>();

for (int i = 1; i < lines.Length; i++)
{
    if (string.IsNullOrWhiteSpace(lines[i])) continue;

    string[] columns = lines[i].Split(',');

    if (columns.Length >= 4)
    {
        var record = new
        {
            Name = columns[0].Trim(),
            Age = int.Parse(columns[1].Trim()),
            Email = columns[2].Trim(),
            Score = float.Parse(columns[3].Trim(), CultureInfo.InvariantCulture)
        };
        dataObjects.Add(record);
    }
}

var options = new JsonSerializerOptions { WriteIndented = true };
string jsonOutput = JsonSerializer.Serialize(dataObjects, options);
Console.WriteLine(jsonOutput);