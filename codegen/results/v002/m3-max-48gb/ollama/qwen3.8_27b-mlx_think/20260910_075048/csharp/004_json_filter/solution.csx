using System.Text.Json;
using System.Text.Json.Nodes;

var input = File.ReadAllText("input/data.json");
var arr = JsonNode.Parse(input) as JsonArray;

var filtered = arr
    .Select(node => node as JsonObject)
    .Where(obj => (bool)obj["active"] && (int)obj["age"] >= 30)
    .OrderBy(obj => (string)obj["name"])
    .Select(obj => (JsonNode)obj)
    .ToList();

var result = new JsonArray(filtered.ToArray());
Console.Write(result.ToJsonString(new JsonSerializerOptions { WriteIndented = true }));