#r "netstandard"

using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Linq;

string filePath = "input/data.json";
string jsonContent = File.ReadAllText(filePath);
var root = JsonNode.Parse(jsonContent) as JsonArray;

var filtered = root
    .Select(node => node.AsObject())
    .Where(obj => 
        obj["active"]?.GetValue<bool>() == true && 
        obj["age"]?.GetValue<int>() >= 30)
    .OrderBy(obj => obj["name"]?.GetValue<string>())
    .ToList();

var result = new JsonArray(filtered);
Console.WriteLine(result.ToJsonString(new JsonSerializerOptions { WriteIndented = false }));