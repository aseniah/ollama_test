#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
if (lines.Length < 2)
{
    Console.WriteLine("[]");
    return;
}

var results = new List<JsonNode>();

// Skip header, start from index 1
for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    var fields = line.Split(',');
    
    // Expected format: Name,Age,Email,Score
    if (fields.Length < 4) continue; 

    string name = fields[0].Trim().Replace("\"", "");
    int age = int.TryParse(fields[1].Trim(), out int parsedAge) ? parsedAge : 0;
    string email = fields[2].Trim().Replace("\"", "");
    
    // Use invariant culture to ensure decimal points are handled as dots
    float score = float.TryParse(fields[3].Trim(), NumberStyles.Float, CultureInfo.InvariantCulture, out float parsedScore) ? parsedScore : 0.0f;

    var obj = new JsonObject();
    obj["Name"] = JsonValue.Create(name);
    obj["Age"] = JsonValue.Create(age);
    obj["Email"] = JsonValue.Create(email);
    obj["Score"] = JsonValue.Create(score);

    results.Add(obj);
}

var options = new JsonSerializerOptions { WriteIndented = true };
string jsonOutput = JsonSerializer.Serialize(results, options);

Console.Write(jsonOutput);