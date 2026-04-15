using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the input file
string jsonString = File.ReadAllText("input/data.json");

// Parse the JSON into a JsonArray
JsonNode rootNode = JsonNode.Parse(jsonString);
JsonArray dataArray = rootNode.AsArray();

// Filter: active == true AND age >= 30
// Sort: by name ascending
var filteredRecords = dataArray
    .Where(node => 
        node["active"]?.GetValue<bool>() == true && 
        node["age"]?.GetValue<int>() >= 30)
    .OrderBy(node => node["name"]?.GetValue<string>())
    .ToList();

// Serialize the filtered and sorted list back to JSON
var options = new JsonSerializerOptions { WriteIndented = true };
string outputJson = JsonSerializer.Serialize(filteredRecords, options);

// Output to stdout
Console.WriteLine(outputJson);