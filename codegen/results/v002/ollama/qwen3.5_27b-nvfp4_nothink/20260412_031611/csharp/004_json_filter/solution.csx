#r "System.Text.Json"
using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var jsonString = File.ReadAllText("input/data.json");
var jsonArray = JsonNode.Parse(jsonString)!.AsArray();

var result = jsonArray
    .Select(node => node!.AsObject())
    .Where(obj =>
        obj["active"]!.GetValue<bool>() &&
        obj["age"]!.GetValue<int>() >= 30)
    .OrderBy(obj => obj["name"]!.GetValue<string>())
    .ToList();

var outputJson = JsonNode.Parse(new System.Text.Json.Utf8StringWriter().Write(result))?.ToString() 
                 ?? result.Select(r => r.ToJsonString()).Aggregate((a, b) => a + "," + b).Insert(0, "[").Append("]");

// Ensure the list is properly wrapped in brackets for output
var options = new JsonSerializerOptions { WriteIndented = false };
Console.Write(result.Count == 0 
    ? "[]" 
    : result.Select(r => r!.ToJsonString(options)).Aggregate((a, b) => a + "," + b).Insert(0, "[").Append("]"));