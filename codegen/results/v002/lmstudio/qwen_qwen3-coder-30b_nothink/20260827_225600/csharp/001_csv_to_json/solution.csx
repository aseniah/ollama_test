using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
var jsonArray = new List<JsonObject>();

for (int i = 1; i < lines.Length; i++)
{
    var values = lines[i].Split(',');
    var obj = new JsonObject
    {
        ["Name"] = values[0],
        ["Age"] = int.Parse(values[1]),
        ["Email"] = values[2],
        ["Score"] = float.Parse(values[3])
    };
    jsonArray.Add(obj);
}

var options = new JsonSerializerOptions
{
    WriteIndented = true
};

Console.WriteLine(JsonSerializer.Serialize(jsonArray, options));