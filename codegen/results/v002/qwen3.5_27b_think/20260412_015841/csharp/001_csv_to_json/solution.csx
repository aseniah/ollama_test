using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
var jsonArray = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;
    
    var parts = line.Split(',');
    if (parts.Length >= 4)
    {
        var obj = new JsonObject
        {
            ["Name"] = parts[0],
            ["Age"] = int.Parse(parts[1]),
            ["Email"] = parts[2],
            ["Score"] = float.Parse(parts[3])
        };
        jsonArray.Add(obj);
    }
}

Console.WriteLine(JsonSerializer.Serialize(jsonArray));