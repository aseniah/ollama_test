using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the JSON file content
string jsonContent = File.ReadAllText("input/data.json");
JsonNode? root = JsonNode.Parse(jsonContent);

// Ensure the parsed root is a JSON array
if (root is JsonArray array) {
    var matches = new List<JsonObject>();

    // Iterate through each element in the array
    foreach (JsonNode? node in array) {
        if (node == null) continue;
        JsonObject? obj = node.AsObject();
        if (obj == null) continue;

        // Check if 'active' is true
        bool active = node["active"]?.GetValue<bool>() == true;
        
        // Check if 'age' is 30 or older
        int? age = node["age"]?.GetValue<int>();

        // Add to matches if both conditions are met
        if (active && age >= 30) {
            matches.Add(obj);
        }
    }

    // Sort the matches by 'name' ascending
    var sortedMatches = matches.OrderBy(x => x["name"]?.GetValue<string>() ?? "").ToList();

    // Serialize the result to a JSON array string
    var options = new JsonSerializerOptions { WriteIndented = true };
    string result = JsonSerializer.Serialize(sortedMatches, options);

    // Output to stdout
    Console.Write(result);
}