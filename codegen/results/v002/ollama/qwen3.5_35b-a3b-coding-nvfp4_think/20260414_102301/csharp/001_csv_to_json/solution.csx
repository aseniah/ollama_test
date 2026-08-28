using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

string[] lines = File.ReadAllLines("input/data.csv");

var result = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;
    
    string[] parts = line.Split(',');
    
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

Console.WriteLine(JsonSerializer.Serialize(result));