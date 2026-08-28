using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the input file
string inputPath = "input/data.json";
string jsonInput = File.ReadAllText(inputPath);

// Parse the JSON array
JsonArray jsonArray = JsonNode.Parse(jsonInput)!.GetValue<JsonArray>();

// Filter and sort
var filtered = new List<JsonObject>();
foreach (var item in jsonArray)
{
    if (item is JsonObject obj)
    {
        bool active = obj["active"]!.GetValue<bool>();
        int age = obj["age"]!.GetValue<int>();
        
        if (active && age >= 30)
        {
            filtered.Add(obj);
        }
    }
}

// Sort by name ascending
var sorted = filtered.OrderBy(o => o["name"]!.GetValue<string>()).ToList();

// Serialize back to JSON
var options = new JsonSerializerOptions
{
    WriteIndented = true
};

// Create a new array of JSON objects with only the required fields
var outputArray = new JsonArray();
foreach (var item in sorted)
{
    var newObj = new JsonObject
    {
        ["name"] = item["name"]!,
        ["age"] = item["age"]!,
        ["active"] = item["active"]!,
        ["score"] = item["score"]!
    };
    outputArray.Add(newObj);
}

// Output the result
Console.WriteLine(outputArray.ToJsonString(options));