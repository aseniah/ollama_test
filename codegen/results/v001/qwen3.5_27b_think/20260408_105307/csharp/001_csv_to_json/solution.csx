using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var filePath = "input/data.csv";
var lines = File.ReadAllLines(filePath);
var result = new List<JsonObject>();

// Skip header (first line)
for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;
    
    var parts = line.Split(',');
    
    if (parts.Length >= 4)
    {
        var obj = new JsonObject
        {
            ["Name"] = parts[0],
            ["Age"] = int.Parse(parts[1]),
            ["Email"] = parts[2],
            ["Score"] = float.Parse(parts[3])
        };
        result.Add(obj);
    }
}

var options = new JsonSerializerOptions
{
    WriteIndented = false
};

Console.WriteLine(JsonSerializer.Serialize(result, options));