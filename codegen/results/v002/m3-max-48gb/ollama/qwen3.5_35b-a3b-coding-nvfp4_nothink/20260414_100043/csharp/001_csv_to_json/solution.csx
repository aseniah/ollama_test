using System;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Collections.Generic;

var lines = File.ReadAllLines("input/data.csv");
var jsonNodes = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    if (string.IsNullOrWhiteSpace(lines[i])) continue;
    
    var parts = lines[i].Split(',');
    if (parts.Length != 4) continue;
    
    var obj = new JsonObject
    {
        ["Name"] = parts[0],
        ["Age"] = int.Parse(parts[1]),
        ["Email"] = parts[2],
        ["Score"] = float.Parse(parts[3])
    };
    
    jsonNodes.Add(obj);
}

Console.WriteLine(jsonNodes.ToJsonString());