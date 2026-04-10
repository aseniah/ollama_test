#r "System.Text.Json"
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var filePath = "input/data.csv";

if (!File.Exists(filePath))
{
    Console.WriteLine("[]");
    return;
}

var lines = File.ReadAllLines(filePath).Skip(1).ToList(); // Skip header
var rows = new List<JsonNode>();

foreach (var line in lines)
{
    if (string.IsNullOrWhiteSpace(line)) continue;

    var parts = line.Split(',');
    
    // Basic validation to ensure enough columns exist
    if (parts.Length < 4) continue;

    try 
    {
        var name = parts[0].Trim().Replace("\"", "");
        var age = int.Parse(parts[1].Trim().Replace("\"", ""));
        var email = parts[2].Trim().Replace("\"", "");
        var score = float.Parse(parts[3].Trim().Replace("\"", ""), System.Globalization.CultureInfo.InvariantCulture);

        var obj = new JsonObject
        {
            ["Name"] = JsonNode.Parse($"\"{name}\""),
            ["Age"] = JsonNode.Create(age),
            ["Email"] = JsonNode.Parse($"\"{email}\""),
            ["Score"] = JsonNode.Create(score)
        };

        rows.Add(obj);
    }
    catch (FormatException) 
    {
        // Skip malformed rows to ensure valid JSON output structure
        continue;
    }
}

var jsonOpt = new JsonSerializerOptions { WriteIndented = true, UseUtcDateTimeKindHandling = false };
Console.WriteLine(JsonSerializer.Serialize(rows, jsonOpt));