using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Globalization;

var lines = File.ReadAllLines("input/data.csv");
var jsonArray = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;
    
    var values = line.Split(',');
    var obj = new JsonObject
    {
        ["Name"] = values[0].Trim(),
        ["Age"] = int.Parse(values[1].Trim()),
        ["Email"] = values[2].Trim(),
        ["Score"] = float.Parse(values[3].Trim(), CultureInfo.InvariantCulture)
    };
    jsonArray.Add(obj);
}

Console.WriteLine(JsonSerializer.Serialize(jsonArray));