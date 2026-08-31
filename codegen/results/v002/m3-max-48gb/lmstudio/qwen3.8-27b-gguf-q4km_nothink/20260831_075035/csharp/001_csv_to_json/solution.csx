using System;
using System.Collections.Generic;
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
    
    result.Add(obj);
}

var json = result.ToJsonString(new JsonSerializerOptions
{
    WriteIndented = false
});

Console.Write(json);