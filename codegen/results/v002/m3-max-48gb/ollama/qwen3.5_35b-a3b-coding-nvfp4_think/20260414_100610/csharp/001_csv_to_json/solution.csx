using System;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
var header = lines[0];
var dataLines = lines.Skip(1);

var jsonArray = new JsonArray();

foreach (var line in dataLines)
{
    var parts = line.Split(',');
    var obj = new JsonObject
    {
        ["Name"] = parts[0],
        ["Age"] = int.Parse(parts[1]),
        ["Email"] = parts[2],
        ["Score"] = float.Parse(parts[3])
    };
    
    jsonArray.Add(obj);
}

Console.WriteLine(JsonSerializer.Serialize(jsonArray));