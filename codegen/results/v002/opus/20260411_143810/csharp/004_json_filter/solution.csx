using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Linq;

var json = File.ReadAllText("input/data.json");
var array = JsonNode.Parse(json)!.AsArray();

var filtered = array
    .Where(n => n!["active"]!.GetValue<bool>() && n!["age"]!.GetValue<int>() >= 30)
    .OrderBy(n => n!["name"]!.GetValue<string>())
    .ToList();

var result = new JsonArray();
foreach (var item in filtered)
    result.Add(JsonNode.Parse(item!.ToJsonString()));

Console.WriteLine(result.ToJsonString());
