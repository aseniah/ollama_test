using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the input file
string jsonContent = File.ReadAllText("input/data.json");

// Parse the JSON array
JsonArray jsonArray = JsonNode.Parse(jsonContent)!.AsArray();

// Filter and sort
var filtered = new List<object>();

foreach (JsonNode node in jsonArray)
{
    JsonObject obj = node!.AsObject();
    
    // Get active field
    bool active = obj["active"]!.GetValue<bool>();
    
    // Get age field
    int age = obj["age"]!.GetValue<int>();
    
    // Filter: active is true AND age >= 30
    if (active && age >= 30)
    {
        filtered.Add(obj);
    }
}

// Sort by name ascending
filtered.Sort((a, b) =>
{
    string nameA = a.AsObject()["name"]!.GetValue<string>();
    string nameB = b.AsObject()["name"]!.GetValue<string>();
    return string.Compare(nameA, nameB, StringComparison.Ordinal);
});

// Build output JSON array
JsonArray outputArray = new JsonArray();
foreach (var item in filtered)
{
    outputArray.Add(item.Clone());
}

// Serialize and output
string outputJson = JsonSerializer.Serialize(outputArray, new JsonSerializerOptions { WriteIndented = true });
Console.WriteLine(outputJson);