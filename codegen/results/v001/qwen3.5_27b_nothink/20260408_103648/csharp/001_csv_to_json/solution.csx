#r "System.Text.Json"
using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

string filePath = "input/data.csv";
if (!File.Exists(filePath))
{
    Console.WriteLine("[]");
    return;
}

string[] lines = File.ReadAllLines(filePath);
List<JsonNode> jsonArray = new List<JsonNode>();

// Skip header (index 0)
for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    string[] parts = line.Split(',');
    
    // Ensure we have enough columns
    if (parts.Length < 4) continue;

    try 
    {
        string name = parts[0].Trim();
        int age = int.Parse(parts[1].Trim());
        string email = parts[2].Trim();
        float score = float.Parse(parts[3].Trim(), System.Globalization.NumberStyles.Float, System.Globalization.CultureInfo.InvariantCulture);

        JsonNode node = new JsonObject();
        node["Name"] = name;
        node["Age"] = age;
        node["Email"] = email;
        node["Score"] = score;

        jsonArray.Add(node);
    }
    catch
    {
        // Skip malformed rows
        continue;
    }
}

JsonArray result = new JsonArray();
result.AddRange(jsonArray);

var options = new JsonSerializerOptions
{
    WriteIndented = true,
    PropertyNamingPolicy = JsonNamingPolicy.CamelCase,
    Encoder = System.Text.Encodings.Web.JavaScriptEncoder.UnsafeRelaxedJsonEscaping
};

// Force PascalCase for specific fields as requested, overriding default CamelCase
// We construct a custom output or just serialize and let the default work, 
// but the prompt asked for "Name", "Age", "Email", "Score" (PascalCase).
// System.Text.Json default is CamelCase. We need to map manually to ensure exact keys.

var finalArray = new JsonArray();
foreach (JsonNode node in jsonArray)
{
    var obj = new JsonObject();
    obj["Name"] = node.AsObject()["Name"];
    obj["Age"] = node.AsObject()["Age"];
    obj["Email"] = node.AsObject()["Email"];
    obj["Score"] = node.AsObject()["Score"];
    finalArray.Add(obj);
}

// Custom serialization to ensure PascalCase keys without a global policy change
var optionsPascal = new JsonSerializerOptions
{
    WriteIndented = true,
    PropertyNamingPolicy = null // No policy, keeps existing names
};

Console.Write(finalArray.ToJsonString(optionsPascal));