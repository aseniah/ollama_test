#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;

var lines = File.ReadAllLines("input/data.csv");
var result = new List<object>();

for (int i = 1; i < lines.Length; i++) // Skip header
{
    var parts = lines[i].Split(',');
    if (parts.Length >= 4)
    {
        result.Add(new
        {
            Name = parts[0].Trim(),
            Age = int.Parse(parts[1].Trim()),
            Email = parts[2].Trim(),
            Score = float.Parse(parts[3].Trim())
        });
    }
}

var json = JsonSerializer.Serialize(result);
Console.Write(json);