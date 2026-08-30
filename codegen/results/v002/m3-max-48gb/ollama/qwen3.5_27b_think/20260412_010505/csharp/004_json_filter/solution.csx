#r "System.Text.Json"
using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string jsonContent = File.ReadAllText("input/data.json");
JsonNode? root = JsonNode.Parse(jsonContent);

var filtered = root!.AsArray()
    .Where(node => node is JsonObject person &&
                   (person["active"]?.GetValue<bool>() ?? false) &&
                   (person["age"]?.GetValue<int>() ?? 0) >= 30)
    .OrderBy(p => (p["name"]?.GetValue<string>() ?? ""));

var result = new JsonArray(filtered.ToArray());
Console.Write(JsonSerializer.Serialize(result));