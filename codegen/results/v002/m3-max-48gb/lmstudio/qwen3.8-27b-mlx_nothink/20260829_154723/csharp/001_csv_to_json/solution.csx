using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
var result = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    if (string.IsNullOrWhiteSpace(lines[i])) continue;
    var parts = lines[i].Split(',');
    if (parts.Length < 4) continue;

    var obj = new JsonObject
    {
        ["Name"] = parts[0],
        ["Age"] = int.Parse(parts[1], CultureInfo.InvariantCulture),
        ["Email"] = parts[2],
        ["Score"] = float.Parse(parts[3], CultureInfo.InvariantCulture)
    };
    result.Add(obj);
}

Console.Write(JsonSerializer.Serialize(result));