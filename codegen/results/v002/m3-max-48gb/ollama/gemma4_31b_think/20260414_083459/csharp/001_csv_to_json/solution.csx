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

// Skip the header row
var data = new List<object>();

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    string[] columns = line.Split(',');
    if (columns.Length >= 4)
    {
        data.Add(new
        {
            Name = columns[0],
            Age = int.Parse(columns[1]),
            Email = columns[2],
            Score = float.Parse(columns[3], CultureInfo.InvariantCulture)
        });
    }
}

string jsonOutput = JsonSerializer.Serialize(data, new JsonSerializerOptions 
{ 
    WriteIndented = true 
});

Console.WriteLine(jsonOutput);