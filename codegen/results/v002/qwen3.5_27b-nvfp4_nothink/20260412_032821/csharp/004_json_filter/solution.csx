#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the input file
string filePath = "input/data.json";
string jsonContent = File.ReadAllText(filePath);

// Parse JSON array
JsonNode? node = JsonNode.Parse(jsonContent);
if (node is not JsonArray people)
{
    Console.Write("[]");
    return;
}

var result = new List<object>();

foreach (JsonElement el in ((System.Text.Json.Utf8JsonReader)null!).GetRawText().AsSpan()) 
{
    // We need to iterate properly. Let's re-parse or cast correctly.
}

// Correct iteration approach for JsonArray in System.Text.Json.Nodes
var filtered = people
    .Where(person =>
    {
        var obj = (JsonObject)person;
        bool isActive = false;
        int age = 0;

        if (obj.TryGetValue("active", out JsonNode? activeNode))
            _ = activeNode.Deserialize<bool>() is true ? isActive = true : false;
        
        if (obj.TryGetValue("age", out JsonNode? ageNode))
            _ = ageNode.Deserialize<int>() is int a ? age = a : 0;

        return isActive && age >= 30;
    })
    .OrderBy(person => ((JsonObject)person).TryGetValue("name", out var n) ? n.ToString() : "")
    .Select(person => (object)person)
    .ToList();

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(filtered, options));