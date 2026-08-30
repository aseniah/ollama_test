using System;
using System.IO;
using System.Globalization;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
var jsonArray = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;
    
    var parts = line.Split(',');
    var obj = new JsonObject
    {
        ["Name"] = parts[0],
        ["Age"] = int.Parse(parts[1]),
        ["Email"] = parts[2],
        ["Score"] = float.Parse(parts[3], CultureInfo.InvariantCulture)
    };
    
    jsonArray.Add(obj);
}

Console.Write(JsonSerializer.Serialize(jsonArray));