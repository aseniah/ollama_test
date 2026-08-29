using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var json = File.ReadAllText("input/data.json");
var root = JsonNode.Parse(json);
if (root is not JsonArray array)
{
    Console.Write("[]");
    return;
}

var filtered = array
    .Select(node => node as JsonObject)
    .Where(obj => obj != null
        && obj["age"] != null
        && obj["age"].GetValue<int>() >= 30
        && obj["active"] != null
        && obj["active"].GetValue<bool>())
    .OrderBy(obj => obj["name"]?.GetValue<string>() ?? string.Empty, StringComparer.Ordinal)
    .ToArray();

var output = new JsonArray();
foreach (var obj in filtered)
{
    output.Add(obj);
}

Console.Write(output.ToJsonString());