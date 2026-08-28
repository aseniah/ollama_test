using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;

var lines = File.ReadAllLines("input/data.csv");
var result = new List<object>();

for (int i = 1; i < lines.Length; i++)
{
    var parts = lines[i].Split(',');
    if (parts.Length >= 4)
    {
        result.Add(new
        {
            Name = parts[0],
            Age = int.Parse(parts[1]),
            Email = parts[2],
            Score = float.Parse(parts[3])
        });
    }
}

Console.WriteLine(JsonSerializer.Serialize(result));