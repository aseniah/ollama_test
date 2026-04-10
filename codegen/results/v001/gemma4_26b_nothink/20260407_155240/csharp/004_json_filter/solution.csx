using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Define the file path
string filePath = "input/data.json";

if (!File.Exists(filePath))
{
    // If file doesn't exist, we exit silently as per "Do not output anything else"
    // though in a real scenario one might log an error.
    return;
}

try
{
    // Read the JSON content
    string jsonString = File.ReadAllText(filePath);

    // Parse the JSON array
    JsonArray rootArray = JsonNode.Parse(jsonString).AsArray();

    // Filter and Sort the array
    // We iterate through the nodes, applying the logic: active == true AND age >= 30
    var filteredNodes = rootArray
        .Where(node => 
        {
            bool isActive = node["active"]?.GetValue<bool>() ?? false;
            int age = node["age"]?.GetValue<int>() ?? 0;
            return isActive && age >= 30;
        })
        .OrderBy(node => 
        {
            // Sort by name ascending
            return node["name"]?.ToString() ?? string.Empty;
        })
        .ToList();

    // Create a new JsonArray for the output
    JsonArray result = new JsonArray();
    foreach (var node in filteredNodes)
    {
        result.Add(node.DeepClone());
    }

    // Output the resulting JSON to stdout
    // Using WriteIndented = true for readable output, but the prompt implies a standard JSON array output
    var options = new JsonSerializerOptions { WriteIndented = true };
    Console.Write(JsonSerializer.Serialize(result, options));
}
catch (Exception)
{
    // If there is a parsing error or file error, we exit without outputting anything
    // to adhere strictly to "Do not output anything else"
}