using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
var jsonArray = new List<JsonObject>();

for (int i = 1; i < lines.Length; i++)
{
    var parts = lines[i].Split(',');
    var obj = new JsonObject
    {
        ["Name"] = parts[0],
        ["Age"] = int.Parse(parts[1]),
        ["Email"] = parts[2],
        ["Score"] = float.Parse(parts[3])
    };
    jsonArray.Add(obj);
}

var options = new JsonSerializerOptions
{
    WriteIndented = true
};

Console.WriteLine(JsonSerializer.Serialize(jsonArray, options));