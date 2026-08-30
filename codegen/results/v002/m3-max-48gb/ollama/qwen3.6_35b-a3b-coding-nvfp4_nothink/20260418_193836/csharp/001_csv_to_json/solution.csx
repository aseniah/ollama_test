using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
var jsonArray = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i].Trim();
    if (string.IsNullOrEmpty(line))
        continue;

    var parts = line.Split(',');
    if (parts.Length < 4)
        continue;

    var obj = new JsonObject
    {
        ["Name"] = parts[0].Trim(),
        ["Age"] = int.Parse(parts[1].Trim()),
        ["Email"] = parts[2].Trim(),
        ["Score"] = float.Parse(parts[3].Trim())
    };

    jsonArray.Add(obj);
}

var json = JsonSerializer.Serialize(jsonArray, new JsonSerializerOptions { WriteIndented = false });
Console.Write(json);