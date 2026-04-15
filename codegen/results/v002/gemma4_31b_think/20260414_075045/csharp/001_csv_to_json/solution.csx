using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;
using System.Globalization;
using System.Text.Json;

if (!File.Exists("input/data.csv"))
{
    Environment.Exit(1);
}

var lines = File.ReadAllLines("input/data.csv");

// Skip the header row
var rows = lines.Skip(1);

var result = new List<object>();

foreach (var row in rows)
{
    if (string.IsNullOrWhiteSpace(row)) continue;

    var columns = row.Split(',');
    if (columns.Length < 4) continue;

    var entry = new
    {
        Name = columns[0],
        Age = int.Parse(columns[1]),
        Email = columns[2],
        Score = float.Parse(columns[3], CultureInfo.InvariantCulture)
    };

    result.Add(entry);
}

var jsonOptions = new JsonSerializerOptions 
{ 
    WriteIndented = true 
};

Console.WriteLine(JsonSerializer.Serialize(result, jsonOptions));