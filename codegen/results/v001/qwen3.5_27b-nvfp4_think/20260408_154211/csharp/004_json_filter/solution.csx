#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var jsonContent = File.ReadAllText("input/data.json");
var jsonArray = JsonNode.Parse(jsonContent)!.AsArray();

var filtered = jsonArray
    .Select(node => node!.AsObject())
    .Where(obj => obj!["active"]!.GetValue<bool>() && obj["age"]!.GetValue<int>() >= 30)
    .OrderBy(obj => obj["name"]!.ToString()!)
    .ToList();

var outputJson = JsonSerializer.Serialize(filtered, new JsonSerializerOptions { WriteIndented = false });
Console.Write(outputJson);