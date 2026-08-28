using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string json = File.ReadAllText("input/data.json");
var array = JsonNode.Parse(json)!.AsArray();

var filtered = array
    .OfType<JsonObject>()
    .Where(obj => obj["active"]!.GetValue<bool>() && obj["age"]!.GetValue<int>() >= 30)
    .OrderBy(obj => obj["name"]!.GetValue<string>())
    .ToArray();

var output = new JsonArray();
foreach (var item in filtered)
{
    output.Add(item.DeepClone());
}

Console.WriteLine(JsonSerializer.Serialize(output, new JsonSerializerOptions { WriteIndented = true }));