using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Define the input file path
string inputPath = "input/data.json";

if (File.Exists(inputPath))
{
    // Read the JSON file content
    string jsonString = File.ReadAllText(inputPath);

    // Parse the JSON string into a JsonArray
    var data = JsonNode.Parse(jsonString).AsArray();

    // Filter the records: active must be true and age must be 30 or older
    // Then sort the filtered records by name in ascending order
    var filteredAndSorted = data
        .Where(node => 
            node["active"]?.GetValue<bool>() == true && 
            node["age"]?.GetValue<int>() >= 30)
        .OrderBy(node => node["name"]?.GetValue<string>())
        .ToList();

    // Serialize the resulting list back to a JSON string and print to stdout
    // We use JsonSerializer to ensure the output is a valid JSON array
    Console.WriteLine(JsonSerializer.Serialize(filteredAndSorted));
}