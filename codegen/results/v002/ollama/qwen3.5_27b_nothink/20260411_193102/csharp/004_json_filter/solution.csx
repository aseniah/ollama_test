using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var jsonData = File.ReadAllText("input/data.json");
var root = JsonNode.Parse(jsonData)!.AsArray();

var filtered = new List<JsonObject>();

foreach (var node in root)
{
    var obj = node.AsObject();
    var name = obj["name"]?.GetValue<string>();
    var age = obj["age"]?.GetValue<int>();
    var active = obj["active"]?.GetValue<bool>();

    if (name != null && age != null && active != null)
    {
        if (active.Value && age.Value >= 30)
        {
            filtered.Add(obj);
        }
    }
}

var sorted = filtered.OrderBy(x => x["name"]?.GetValue<string>() ?? "").ToList();

var options = new JsonSerializerOptions
{
    WriteIndented = false,
    PropertyNamingPolicy = JsonNamingPolicy.CamelCase
};

var output = JsonSerializer.Serialize(sorted, options);
Console.Write(output);