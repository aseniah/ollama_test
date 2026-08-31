using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var json = File.ReadAllText("input/data.json");
var nodes = JsonNode.Parse(json).AsArray();

var filtered = nodes
    .Select(n => n.AsObject())
    .Where(o => (bool)o["active"] && (int)o["age"] >= 30)
    .OrderBy(o => (string)o["name"])
    .ToList();

var result = new JsonArray();
foreach (var obj in filtered)
{
    result.Add(obj);
}

Console.WriteLine(result.ToJsonString(new JsonSerializerOptions { WriteIndented = true }));