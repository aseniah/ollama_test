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

// Skip the header row
var dataList = new List<object>();

for (int i = 1; i < lines.Length; i++)
{
    if (string.IsNullOrWhiteSpace(lines[i])) continue;

    string[] columns = lines[i].Split(',');
    if (columns.Length < 4) continue;

    var person = new
    {
        Name = columns[0].Trim(),
        Age = int.Parse(columns[1].Trim()),
        Email = columns[2].Trim(),
        Score = float.Parse(columns[3].Trim(), CultureInfo.InvariantCulture)
    };

    dataList.Add(person);
}

var options = new JsonSerializerOptions { WriteIndented = true };
string jsonOutput = JsonSerializer.Serialize(dataList, options);
Console.WriteLine(jsonOutput);