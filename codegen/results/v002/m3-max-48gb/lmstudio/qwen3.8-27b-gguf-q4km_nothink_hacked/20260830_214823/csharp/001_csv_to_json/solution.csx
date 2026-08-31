using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
var array = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var parts = lines[i].Split(',');
    if (parts.Length < 4) continue;

    var obj = new JsonObject
    {
        ["Name"] = parts[0],
        ["Age"] = int.Parse(parts[1]),
        ["Email"] = parts[2],
        ["Score"] = float.Parse(parts[3])
    };
    array.Add(obj);
}

var json = array.ToJsonString(new JsonSerializerOptions { WriteIndented = false });
Console.Write(json);