using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Text;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
var results = new List<JsonObject>();

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;
    
    var parts = line.Split(',');
    if (parts.Length < 4) continue;
    
    var name = parts[0].Trim();
    var age = int.Parse(parts[1].Trim());
    var email = parts[2].Trim();
    var score = float.Parse(parts[3].Trim(), CultureInfo.InvariantCulture);
    
    var obj = new JsonObject
    {
        ["Name"] = name,
        ["Age"] = age,
        ["Email"] = email,
        ["Score"] = score
    };
    
    results.Add(obj);
}

var json = results.ToString();
Console.Write(json);