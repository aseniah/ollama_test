using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

string filePath = "input/data.json";

if (!File.Exists(filePath))
{
    Environment.Exit(1);
}

string jsonString = File.ReadAllText(filePath);
var root = JsonNode.Parse(jsonString).AsArray();

var filteredList = root
    .Select(node => node.AsObject())
    .Where(obj => 
        (bool)obj["active"] == true && 
        (int)obj["age"] >= 30
    )
    .OrderBy(obj => (string)obj["name"])
    .ToList();

var options = new JsonSerializerOptions { WriteIndented = true };
string outputJson = JsonSerializer.Serialize(filteredList, options);

Console.Write(outputJson);