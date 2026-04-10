using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;
using System.Text.Json;

// Define the structure of our data objects
// Since we cannot use classes with specific properties easily in top-level scripts 
// without defining them, we'll use JsonNode for flexibility.

string filePath = "input/data.json";

if (!File.Exists(filePath))
{
    // If file doesn't exist, we output an empty array to remain valid JSON
    Console.WriteLine("[]");
    return;
}

try
{
    string jsonString = File.ReadAllText(filePath);
    using JsonDocument doc = JsonDocument.Parse(jsonString);
    
    // Convert to a list of JsonNodes to allow for easy manipulation and sorting
    var root = JsonNode.Parse(jsonString)!.AsArray();

    var filteredList = root
        .Where(node =>
        {
            var name = node["name"];
            var age = node["age"];
            var active = node["active"];

            // Check existence and types to prevent NullReferenceException
            bool isActive = active != null && active.GetValue<bool>();
            bool isOldEnough = age != null && age.GetValue<int>() >= 30;

            return isActive && isOldEnough;
        })
        .OrderBy(node => node["name"]?.GetValue<string>() ?? string.Empty)
        .ToList();

    // Serialize the filtered list back to a JSON string
    var options = new JsonSerializerOptions { WriteIndented = true };
    string result = JsonSerializer.Serialize(filteredList, options);

    Console.Write(result);
}
catch (Exception)
{
    // In case of malformed JSON or unexpected structure, output empty array
    Console.Write("[]");
}