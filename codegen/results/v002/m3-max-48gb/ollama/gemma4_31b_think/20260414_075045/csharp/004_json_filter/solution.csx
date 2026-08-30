using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the JSON file
string jsonString = File.ReadAllText("input/data.json");

// Parse the JSON input into a JsonArray
JsonArray data = JsonNode.Parse(jsonString).AsArray();

// Filter records: active is true AND age >= 30
// Then sort by name ascending
var filteredRecords = data
    .Where(node => 
        node["active"]?.GetValue<bool>() == true && 
        node["age"]?.GetValue<int>() >= 30)
    .OrderBy(node => node["name"]?.GetValue<string>())
    .ToList();

// Serialize the result back to JSON and print to stdout
// We use an array to hold the filtered nodes for serialization
string outputJson = JsonSerializer.Serialize(filteredRecords, new JsonSerializerOptions 
{ 
    WriteIndented = true 
});

Console.WriteLine(outputJson);