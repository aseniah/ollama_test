using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var jsonStr = File.ReadAllText("input/data.json");
var array = JsonNode.Parse(jsonStr)!.AsArray();

var filtered = new List<JsonObject>();
foreach (var node in array)
{
    if (node is JsonObject obj)
    {
        bool active = obj["active"]?.GetValue<bool>() ?? false;
        int age = obj["age"]?.GetValue<int>() ?? 0;
        if (active && age >= 30)
        {
            filtered.Add(obj);
        }
    }
}

filtered.Sort((a, b) => string.Compare(a["name"]!.GetValue<string>(), b["name"]!.GetValue<string>(), StringComparison.Ordinal));

var result = new JsonArray();
foreach (var obj in filtered)
{
    result.Add(obj.Clone());
}

Console.WriteLine(JsonSerializer.Serialize(result, new JsonSerializerOptions { WriteIndented = false }));