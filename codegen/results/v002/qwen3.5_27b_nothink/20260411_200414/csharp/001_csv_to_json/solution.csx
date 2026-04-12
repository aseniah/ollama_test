#r "nuget: System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var filePath = "input/data.csv";
if (!File.Exists(filePath))
{
    Console.Error.WriteLine($"File not found: {filePath}");
    Environment.Exit(1);
}

var lines = File.ReadAllLines(filePath);
var records = new List<JsonNode>();

// Skip header (index 0)
for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    // Split by comma. Note: This simple split assumes no commas inside quoted fields.
    var parts = line.Split(',');
    if (parts.Length != 4) continue;

    var name = parts[0];
    if (!int.TryParse(parts[1], out int age)) continue;
    var email = parts[2];
    if (!float.TryParse(parts[3], out float score)) continue;

    var record = new JsonObject
    {
        ["Name"] = name,
        ["Age"] = age,
        ["Email"] = email,
        ["Score"] = score
    };

    records.Add(record);
}

var options = new JsonWriterOptions { Indented = true };
Console.WriteLine(JsonNode.Parse(JsonSerializer.Serialize(new JsonArray(records), options).ToString()));