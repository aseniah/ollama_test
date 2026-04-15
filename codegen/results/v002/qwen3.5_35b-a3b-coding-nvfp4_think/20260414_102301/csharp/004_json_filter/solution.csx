using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Linq;

// Read the JSON file
var jsonString = File.ReadAllText("input/data.json");

// Parse the JSON array
var jsonNodes = JsonNode.Parse(jsonString);
var jsonArray = jsonNodes.AsArray();

// Filter: active is true AND age >= 30
var filtered = jsonArray
    .Where(node =>
    {
        if (node["active"] is JsonValue activeValue && activeValue.GetValueKind() == JsonValueKind.True)
        {
            if (node["age"] is JsonValue ageValue && ageValue.TryGetValue<int>(out var age))
            {
                return age >= 30;
            }
        }
        return false;
    })
    .OrderBy(node => node["name"]?.GetValue<string>())
    .ToArray();

// Serialize to JSON and output
var output = JsonSerializer.Serialize(filtered);
Console.WriteLine(output);