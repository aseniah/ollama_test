using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Collections.Generic;

var lines = File.ReadAllLines("input/data.csv");
var array = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    if (string.IsNullOrWhiteSpace(lines[i])) continue;
    var parts = lines[i].Split(',');
    var obj = new JsonObject
    {
        ["Name"] = parts[0].Trim(),
        ["Age"] = int.Parse(parts[1].Trim()),
        ["Email"] = parts[2].Trim(),
        ["Score"] = float.Parse(parts[3].Trim())
    };
    array.Add(obj);
}

Console.Write(array.ToJsonString(new JsonSerializerOptions { WriteIndented = true }));