using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the JSON file from the specified path
string filePath = "input/data.json";
if (!File.Exists(filePath))
{
    return;
}

string jsonString = File.ReadAllText(filePath);

// Parse the JSON content into a JsonArray
var root = JsonNode.Parse(jsonString)?.AsArray();

if (root == null)
{
    Console.WriteLine("[]");
    return;
}

// Filter the records: active is true and age is 30 or older
// Sort the filtered records by name ascending
var filteredRecords = root
    .Where(node => 
        node["active"]?.GetValue<bool>() == true && 
        node["age"]?.GetValue<int>() >= 30)
    .OrderBy(node => node["name"]?.GetValue<string>())
    .ToList();

// Create a new JsonArray to hold the filtered and sorted results
var resultArray = new JsonArray();
foreach (var record in filteredRecords)
{
    // Since we are dealing with JsonNodes, we can add them directly to the new array
    resultArray.Add(record);
}

// Output the final JSON array to stdout
Console.WriteLine(resultArray.ToJsonString());