using System;
using System.Collections.Generic;
using System.Text.Json;

var lines = File.ReadAllLines("input/data.csv");
var jsonOutput = new List<JsonNode>();

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i].Trim();
    if (string.IsNullOrEmpty(line))
        continue;

    var parts = line.Split(',');
    if (parts.Length < 4)
        continue;

    var name = parts[0].Trim();
    var age = int.Parse(parts[1].Trim());
    var email = parts[2].Trim();
    var score = float.Parse(parts[3].Trim());

    var obj = new JsonObject
    {
        ["Name"] = name,
        ["Age"] = age,
        ["Email"] = email,
        ["Score"] = score
    };

    jsonOutput.Add(obj);
}

var array = new JsonArray(jsonOutput.ToArray());
var json = array.ToJsonString();
Console.Write(json);