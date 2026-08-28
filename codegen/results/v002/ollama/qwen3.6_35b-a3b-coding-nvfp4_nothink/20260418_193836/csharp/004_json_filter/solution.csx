using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var json = File.ReadAllText("input/data.json");
var nodes = JsonNode.Parse(json)!;
var array = nodes.AsArray();

var filtered = new List<JsonObject>();

foreach (var node in array)
{
    if (node == null) continue;
    var obj = node.AsObject();
    
    if (!obj.TryGetValue("active", out var activeNode) || !activeNode!.AsValue().GetBoolean())
        continue;
    
    if (!obj.TryGetValue("age", out var ageNode) || ageNode == null)
        continue;
    
    int age = ageNode.AsValue().GetInt32();
    if (age < 30)
        continue;
    
    filtered.Add(obj);
}

filtered.Sort((a, b) =>
{
    string nameA = a["name"]!.ToString();
    string nameB = b["name"]!.ToString();
    return string.Compare(nameA, nameB, StringComparison.Ordinal);
});

var result = new JsonArray();
foreach (var obj in filtered)
{
    result.Add(obj.Clone());
}

Console.WriteLine(result.ToJsonString());