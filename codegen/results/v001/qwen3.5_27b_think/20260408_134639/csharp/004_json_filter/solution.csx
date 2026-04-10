using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var json = File.ReadAllText("input/data.json");
var array = JsonNode.Parse(json).AsArray();

var filtered = array
    .Where(node =>
    {
        var obj = node.AsObject();
        var active = obj["active"]?.GetValue<bool>() ?? false;
        var age = obj["age"]?.GetValue<int>() ?? 0;
        return active && age >= 30;
    })
    .OrderBy(node => node.AsObject()["name"]?.GetValue<string>())
    .ToArray();

Console.Write(JsonSerializer.Serialize(filtered));