using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Linq;
using System.Collections.Generic;

string filePath = "input/data.json";

if (!File.Exists(filePath))
{
    Environment.Exit(1);
}

string jsonString = File.ReadAllText(filePath);
var root = JsonNode.Parse(jsonString).AsArray();

var filtered = root
    .Select(node => node.AsObject())
    .Where(obj => 
        obj["active"]?.GetValue<bool>() == true && 
        obj["age"]?.GetValue<int>() >= 30)
    .OrderBy(obj => obj["name"]?.GetValue<string>())
    .ToList();

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(filtered, options));