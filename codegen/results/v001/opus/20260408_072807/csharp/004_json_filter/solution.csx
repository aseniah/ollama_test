using System.Text.Json;
using System.Text.Json.Nodes;

var json = File.ReadAllText("input/data.json");
var array = JsonNode.Parse(json)!.AsArray();

var filtered = array
    .Where(item => item!["active"]!.GetValue<bool>() && item["age"]!.GetValue<int>() >= 30)
    .OrderBy(item => item!["name"]!.GetValue<string>())
    .ToList();

var result = new JsonArray(filtered.Select(item => JsonNode.Parse(item!.ToJsonString())).ToArray());

Console.WriteLine(result.ToJsonString());
