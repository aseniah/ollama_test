using System;
using System.IO;
using System.Text.Json.Nodes;
using System.Globalization;

var lines = File.ReadAllLines("input/data.csv");
var array = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var parts = lines[i].Split(',');
    var obj = new JsonObject
    {
        ["Name"] = parts[0],
        ["Age"] = int.Parse(parts[1]),
        ["Email"] = parts[2],
        ["Score"] = float.Parse(parts[3], CultureInfo.InvariantCulture)
    };
    array.Add(obj);
}

Console.Write(array.ToJsonString(new System.Text.Json.JsonWriterOptions { Indented = true }));