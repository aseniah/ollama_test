using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Collections.Generic;

// Define the path to the input file
string inputPath = "input/data.json";

// Check if the file exists before attempting to read
if (!File.Exists(inputPath))
{
    return;
}

// Read the entire content of the JSON file
string jsonContent = File.ReadAllText(inputPath);

// Parse the JSON string into a JsonArray
// JsonNode.Parse returns a JsonNode, so we convert it to a JsonArray
var rootNode = JsonNode.Parse(jsonContent);

if (rootNode is JsonArray array)
{
    // Filter the array based on the requirements:
    // 1. active must be true
    // 2. age must be 30 or older
    // 3. Sort the resulting list by name in ascending order
    var filteredList = array
        .Where(node => 
            node != null &&
            node["active"] != null && node["active"].GetValue<bool>() == true &&
            node["age"] != null && node["age"].GetValue<int>() >= 30)
        .OrderBy(node => (string)node["name"])
        .ToList();

    // Serialize the filtered List<JsonNode> back to a JSON string and output to stdout
    // No extra text or formatting is added to satisfy the "Do not output anything else" requirement
    Console.WriteLine(JsonSerializer.Serialize(filteredList));
}