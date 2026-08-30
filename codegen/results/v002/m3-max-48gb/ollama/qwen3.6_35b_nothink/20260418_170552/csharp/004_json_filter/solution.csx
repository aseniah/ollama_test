using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the input file
string inputPath = "input/data.json";
string jsonContent = File.ReadAllText(inputPath);

// Parse the JSON array
JsonArray dataArray = JsonNode.Parse(jsonContent)!.AsArray();

// Filter: active is true and age is 30 or older
var filtered = new List<JsonObject>();
foreach (var item in dataArray)
{
    if (item == null) continue;
    
    var obj = item.AsObject();
    
    // Check active is true
    bool active = obj["active"]?.GetValue<bool>() ?? false;
    
    // Check age is 30 or older
    int age = obj["age"]?.GetValue<int>() ?? 0;
    
    if (active && age >= 30)
    {
        filtered.Add(obj);
    }
}

// Sort by name ascending
var sorted = filtered.OrderBy(obj => obj["name"]!.GetValue<string>()).ToList();

// Create output array
var outputArray = new JsonArray();
foreach (var obj in sorted)
{
    outputArray.Add(obj.Clone());
}

// Write to stdout
JsonSerializerOptions options = new JsonSerializerOptions
{
    WriteIndented = true
};
string outputJson = outputArray.ToJsonString(options);
Console.Write(outputJson);