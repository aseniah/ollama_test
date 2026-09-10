using System;
using System.IO;
using System.Collections.Generic;
using System.Text.Json;
using System.Globalization;

string filePath = "input/data.csv";

if (!File.Exists(filePath))
{
    Environment.Exit(1);
}

string[] lines = File.ReadAllLines(filePath);
if (lines.Length <= 1)
{
    Console.WriteLine("[]");
    return;
}

var resultList = new List<object>();

// Skip header (index 0)
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
        resultList.Add(record);
    }
}

var options = new JsonSerializerOptions { WriteIndented = true };
string jsonOutput = JsonSerializer.Serialize(resultList, options);
Console.Write(jsonOutput);