using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var json = File.ReadAllText("input/data.json");
var arr = JsonNode.Parse(json) as JsonArray;

var filtered = arr
    .Select(n => (Record)n.AsObject())
    .Where(r => (bool)r["active"].GetValue<bool>() && r["age"].GetValue<int>() >= 30)
    .OrderBy(r => r["name"].GetValue<string>(), StringComparer.Ordinal)
    .Select(r => r.AsObject())
    .ToList();

var result = new JsonArray();
foreach (var r in filtered)
    result.Add(r);

Console.Write(result.ToJsonString(new JsonSerializerOptions { WriteIndented = false }));