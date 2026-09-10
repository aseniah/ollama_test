using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Collections.Generic;

string filePath = "input/data.json";

if (!File.Exists(filePath))
{
    return;
}

string jsonString = File.ReadAllText(filePath);

// Deserialize into a list of dynamic objects (JsonElement) to handle the filtering
var data = JsonSerializer.Deserialize<List<JsonElement>>(jsonString);

if (data == null)
{
    Console.WriteLine("[]");
    return;
}

// Filter: active == true AND age >= 30
// Sort: name ascending
var filteredData = data
    .Where(item => 
        item.GetProperty("active").GetBoolean() == true && 
        item.GetProperty("age").GetInt32() >= 30)
    .OrderBy(item => item.GetProperty("name").GetString())
    .Select(item => 
    {
        // Create a dictionary to represent the object for clean serialization
        var dict = new Dictionary<string, object>();
        foreach (var prop in item.EnumerateObject())
        {
            // Extract values based on type
            object value = prop.Value.ValueKind switch
            {
                JsonValueKind.String => prop.Value.GetString(),
                JsonValueKind.Number => prop.Value.GetDouble(),
                JsonValueKind.True => true,
                JsonValueKind.False => false,
                _ => prop.Value.GetRawText()
            };
            dict[prop.Name] = value;
        }
        return dict;
    })
    .ToList();

// Serialize the result back to JSON
var options = new JsonSerializerOptions { WriteIndented = false };
string resultJson = JsonSerializer.Serialize(filteredData, options);

Console.WriteLine(resultJson);