#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Text.Json.Serialization;

var filePath = "input/data.csv";

if (!File.Exists(filePath))
{
    Console.Write("[]");
    return;
}

var lines = File.ReadAllLines(filePath);
if (lines.Length < 2)
{
    Console.Write("[]");
    return;
}

// Skip header (index 0)
var objects = new List<object>();

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;

    // Split by comma, handling potential whitespace
    var parts = line.Split(',');
    if (parts.Length < 4) continue;

    try
    {
        var name = parts[0].Trim();
        var age = int.Parse(parts[1].Trim());
        var email = parts[2].Trim();
        var score = float.Parse(parts[3].Trim());

        objects.Add(new
        {
            Name = name,
            Age = age,
            Email = email,
            Score = score
        });
    }
    catch (Exception)
    {
        // Skip malformed rows
        continue;
    }
}

var options = new JsonSerializerOptions
{
    WriteIndented = true,
    PropertyNamingPolicy = JsonNamingPolicy.CamelCase
};

// Output to stdout
Console.WriteLine(JsonSerializer.Serialize(objects, options));