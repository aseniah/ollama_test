using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");

if (lines.Length == 0)
{
    Console.Write("[]");
    return;
}

var array = new JsonArray();

// Skip header (index 0)
for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;
    
    var parts = line.Split(',');
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
    
    array.Add(obj);
}

Console.Write(array.ToJsonString(new JsonSerializerOptions { WriteIndented = false }));