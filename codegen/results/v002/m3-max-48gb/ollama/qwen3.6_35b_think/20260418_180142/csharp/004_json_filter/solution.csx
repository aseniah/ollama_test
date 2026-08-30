using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var data = File.ReadAllText("input/data.json");
var nodes = JsonNode.Parse(data)!.AsArray();

var filtered = nodes
    .OfType<JsonObject>()
    .Where(obj => obj["active"]!.GetValue<bool>() && obj["age"]!.GetValue<int>() >= 30)
    .OrderBy(obj => obj["name"]!.GetValue<string>())
    .ToArray();

var result = new JsonArray(filtered);
Console.WriteLine(result.ToJsonString());