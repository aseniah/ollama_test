using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Linq;
using System.Collections.Generic;

if (!File.Exists("input/data.json"))
{
    return;
}

string jsonString = File.ReadAllText("input/data.json");
var data = JsonNode.Parse(jsonString)?.AsArray();

if (data == null)
{
    Console.WriteLine("[]");
    return;
}

var filteredAndSorted = data
    .Where(node => 
        node?["active"]?.GetValue<bool>() == true && 
        node?["age"]?.GetValue<int>() >= 30)
    .OrderBy(node => node?["name"]?.GetValue<string>())
    .ToList();

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(filteredAndSorted, options));