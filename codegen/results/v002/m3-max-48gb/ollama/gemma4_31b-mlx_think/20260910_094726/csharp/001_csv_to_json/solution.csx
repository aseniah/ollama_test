using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using System.Globalization;

string filePath = "input/data.csv";

if (!File.Exists(filePath))
{
    return;
}

string[] lines = File.ReadAllLines(filePath);
var dataList = new List<object>();

// Skip the header line
for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    string[] parts = line.Split(',');
    if (parts.Length < 4) continue;

    dataList.Add(new
    {
        Name = parts[0].Trim(),
        Age = int.Parse(parts[1].Trim()),
        Email = parts[2].Trim(),
        Score = float.Parse(parts[3].Trim(), CultureInfo.InvariantCulture)
    });
}

var options = new JsonSerializerOptions 
{ 
    WriteIndented = true 
};

string jsonOutput = JsonSerializer.Serialize(dataList, options);
Console.WriteLine(jsonOutput);