using System;
using System.IO;
using System.Linq;
using System.Text.Json;

// Read the JSON file
string jsonContent = File.ReadAllText("input/data.json");

// Parse the JSON array
var jsonArray = JsonNode.Parse(jsonContent)!.AsArray();

// Filter records where active is true and age >= 30, then sort by name
var filtered = jsonArray
    .Where(node =>
    {
        var record = node.AsObject();
        bool active = record["active"]!.GetValue<bool>();
        int age = record["age"]!.GetValue<int>();
        return active && age >= 30;
    })
    .OrderBy(node => node.AsObject()["name"]!.GetValue<string>())
    .ToArray();

// Create a new JSON array with filtered records
var resultArray = new JsonArray();
foreach (var item in filtered)
{
    resultArray.Add(item);
}

// Output the result as JSON
Console.WriteLine(JsonSerializer.Serialize(resultArray));