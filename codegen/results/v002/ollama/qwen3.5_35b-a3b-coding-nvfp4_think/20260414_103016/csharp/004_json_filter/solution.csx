using System;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Linq;
using System.Collections.Generic;

// Read the JSON file
var jsonContent = File.ReadAllText("input/data.json");
var jsonNodes = JsonNode.Parse(jsonContent);

// Filter records where active is true and age >= 30
var filtered = ((JsonArray)jsonNodes!)!
    .Cast<JsonObject>()
    .Where(obj =>
        obj["active"]?.GetValue<bool>() == true &&
        obj["age"]?.GetValue<int>() >= 30)
    .OrderBy(obj => obj["name"]?.GetValue<string>() ?? "")
    .ToList();

// Create output array
var output = new JsonArray();
foreach (var obj in filtered)
{
    output.Add(obj);
}

// Write to stdout
Console.WriteLine(JsonSerializer.Serialize(output));