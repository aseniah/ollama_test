using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Globalization;

string filePath = "input/data.csv";

if (!File.Exists(filePath))
{
    GC.Collect(); // Ensure no unexpected output
    return;
}

string[] lines = File.ReadAllLines(filePath);

// Skip header row
var dataObjects = new List<object>();

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    string[] columns = line.Split(',');

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