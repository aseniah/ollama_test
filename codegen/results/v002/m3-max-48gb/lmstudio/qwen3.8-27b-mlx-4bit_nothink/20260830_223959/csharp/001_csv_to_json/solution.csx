using System;
using System.IO;
using System.Text;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Collections.Generic;

string[] lines = File.ReadAllLines("input/data.csv");

var result = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;

    string[] parts = line.Split(',');

    var obj = new JsonObject
    {
        ["Name"] = parts[0].Trim(),
        ["Age"] = int.Parse(parts[1].Trim()),
        ["Email"] = parts[2].Trim(),
        ["Score"] = float.Parse(parts[3].Trim())
    };

    result.Add(obj);
}

Console.WriteLine(result.ToJsonString(new JsonSerializerOptions { WriteIndented = true }));