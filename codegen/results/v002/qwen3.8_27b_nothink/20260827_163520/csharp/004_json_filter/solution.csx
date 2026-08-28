using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Linq;

var json = File.ReadAllText("input/data.json");
var array = JsonNode.Parse(json)!.AsArray();

var filtered = array
    .Select(node => node!.AsObject())
    .Where(obj => obj["active"]!.GetValue<bool>() && obj["age"]!.GetValue<int>() >= 30)
    .OrderBy(obj => obj["name"]!.GetValue<string>(), StringComparer.Ordinal)
    .Select(obj => (JsonObject)obj)
    .ToList();

var result = new JsonArray();
foreach (var obj in filtered)
{
    result.Add(JsonNode.Parse(obj.ToJsonString()));
}

Console.WriteLine(result.ToJsonString(new JsonSerializerOptions { WriteIndented = true }));