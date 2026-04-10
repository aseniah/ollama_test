using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;
using System.Text.Json;

string filePath = "input/data.json";

if (!File.Exists(filePath))
{
    return;
}

try
{
    string jsonString = File.ReadAllText(filePath);
    
    // Parse the JSON array into a list of JsonNodes for flexible manipulation
    var options = new JsonSerializerOptions { AllowTrailingCommas = true };
    var data = JsonSerializer.Deserialize<List<JsonNode>>(jsonString, options);

    if (data == null)
    {
        Console.WriteLine("[]");
        return;
    }

    // Filter: active == true AND age >= 30
    // Sort: name ascending
    var filteredData = data
        .Where(node => 
            node["active"]?.GetValue<bool>() == true && 
            node["age"]?.GetValue<int>() >= 30)
        .OrderBy(node => node["name"]?.GetValue<string>(), StringComparer.Ordinal)
        .ToList();

    // Write the resulting JSON array to stdout
    string result = JsonSerializer.Serialize(filteredData, new JsonSerializerOptions { WriteIndented = true });
    Console.Write(result);
}
catch (Exception)
{
    // In case of malformed JSON or IO errors, output an empty array as a fallback
    // to ensure we "Do not output anything else" besides a JSON array.
    Console.Write("[]");
}