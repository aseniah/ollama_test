#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the JSON file
string filePath = "input/data.json";
string jsonContent = File.ReadAllText(filePath);

// Parse the JSON content
JsonNode? root = JsonNode.Parse(jsonContent);
if (root is JsonArray jsonArray)
{
    // Filter and sort
    var filtered = jsonArray
        .Where(node => 
        {
            if (node is JsonObject obj)
            {
                bool active = obj["active"]?.GetValue<bool>() ?? false;
                int age = obj["age"]?.GetValue<int>() ?? 0;
                return active && age >= 30;
            }
            return false;
        })
        .Select(node => node as JsonObject)
        .Where(o => o != null)
        .OrderBy(obj => obj!["name"]?.GetValue<string>() ?? string.Empty);

    // Convert back to array and output
    var result = new List<JsonNode>();
    foreach (var obj in filtered)
    {
        result.Add(obj);
    }

    JsonArray outputArray = new JsonArray(result);
    Console.WriteLine(outputArray.ToJsonString());
}