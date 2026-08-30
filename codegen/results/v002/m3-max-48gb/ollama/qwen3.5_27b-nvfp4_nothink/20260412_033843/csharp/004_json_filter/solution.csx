#r "System.Text.Json.dll"

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the JSON file
string jsonContent = File.ReadAllText("input/data.json");

// Parse the JSON array
var rootArray = JsonNode.Parse(jsonContent) as JsonArray;

if (rootArray == null)
{
    throw new InvalidDataException("The input is not a valid JSON array.");
}

// Filter and sort records
var filteredRecords = new List<JsonElement>();

foreach (var node in rootArray.OfType<JsonObject>())
{
    string name = node["name"]?.GetValue<string>() ?? "";
    int? age = node["age"]?.GetValue<int?>();
    bool? active = node["active"]?.GetValue<bool?>();

    if (active == true && age.HasValue && age.Value >= 30)
    {
        filteredRecords.Add(JsonSerializer.SerializeToElement(node));
    }
}

// Sort by name ascending
var sortedRecords = filteredRecords
    .OrderBy(element => element.GetProperty("name").GetString())
    .ToList();

// Create the output array
var resultArray = new JsonArray(sortedRecords.ToArray());

// Serialize to JSON string
string outputJson = JsonSerializer.Serialize(resultArray, new JsonSerializerOptions 
{ 
    WriteIndented = true 
});

// Output to stdout
Console.WriteLine(outputJson);