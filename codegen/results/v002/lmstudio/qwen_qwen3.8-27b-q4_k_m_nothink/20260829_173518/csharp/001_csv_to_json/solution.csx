using System;
using System.IO;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
var result = new JsonArray();

foreach (var line in lines.Skip(1))
{
    if (string.IsNullOrWhiteSpace(line)) continue;
    var parts = line.Split(',');
    if (parts.Length < 4) continue;
    
    var obj = new JsonObject
    {
        ["Name"] = parts[0].Trim(),
        ["Age"] = int.Parse(parts[1].Trim()),
        ["Email"] = parts[2].Trim(),
        ["Score"] = float.Parse(parts[3].Trim())
    };
    result.Add(obj);
}

Console.WriteLine(System.Text.Json.JsonSerializer.Serialize(result));