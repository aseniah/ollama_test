#r "System.Text.Json"
using System;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Collections.Generic;
using System.Linq;

// Read the JSON file
string jsonContent = System.IO.File.ReadAllText("input/data.json");

// Parse the JSON array
JsonArray dataArray = JsonNode.Parse(jsonContent).AsArray();

// Filter records where active is true and age is 30 or older
var filtered = dataArray.AsEnumerable()
    .Select(x => x.AsObject())
    .Where(obj => obj["active"].GetValue<bool>() && obj["age"].GetValue<int>() >= 30)
    .OrderBy(obj => obj["name"].GetValue<string>())
    .ToList();

// Create output array
JsonArray result = new JsonArray();
foreach (var item in filtered)
{
    result.Add(item);
}

// Output the JSON result
Console.WriteLine(JsonSerializer.Serialize(result));