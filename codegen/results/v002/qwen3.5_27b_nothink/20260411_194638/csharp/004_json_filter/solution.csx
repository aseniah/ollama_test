#r "System.Text.Json"
#r "System.Linq"

using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Collections.Generic;
using System.Linq;

var inputFile = "input/data.json";
var jsonContent = File.ReadAllText(inputFile);
var jsonArray = JsonArray.Parse(jsonContent);

var results = new List<object>();

foreach (JsonElement elem in jsonArray)
{
    var obj = (JsonObject)elem;
    
    // Extract properties
    var name = obj["name"]?.ToString() ?? string.Empty;
    var age = obj["age"]?.GetValue<int>() ?? 0;
    var active = obj["active"]?.GetValue<bool>() ?? false;
    // score is not needed for filtering or sorting, but we keep it in the output object

    if (active && age >= 30)
    {
        results.Add((JsonObject)obj);
    }
}

// Sort by name ascending
var sortedResults = results.OrderBy(o => ((JsonObject)o)["name"]?.ToString() ?? string.Empty).ToList();

// Convert back to JSON array and output
var outputArray = new JsonArray();
foreach (var item in sortedResults)
{
    outputArray.Add((JsonObject)item);
}

Console.WriteLine(outputArray.ToJsonString());