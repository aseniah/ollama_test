#r "System.Text.Json"

using System;
using System.IO;
using System.Text.Json;
using System.Collections.Generic;
using System.Linq;

var lines = File.ReadAllLines("input/data.csv");
var data = new List<object>();

for (int i = 1; i < lines.Length; i++)
{
    var parts = lines[i].Split(',');
    data.Add(new {
        Name = parts[0],
        Age = int.Parse(parts[1]),
        Email = parts[2],
        Score = float.Parse(parts[3])
    });
}

Console.WriteLine(JsonSerializer.Serialize(data));