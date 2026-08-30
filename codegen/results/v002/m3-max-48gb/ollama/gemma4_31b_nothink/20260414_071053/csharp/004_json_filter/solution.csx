using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

if (!File.Exists("input/data.json"))
{
    Environment.Exit(1);
}

string jsonString = File.ReadAllText("input/data.json");
JsonArray data = JsonNode.Parse(jsonString)!.AsArray();

var filteredAndSorted = data
    .Select(node => node!.AsObject())
    .Where(obj => 
        (bool)obj["active"]! && 
        (int)obj["age"]! >= 30)
    .OrderBy(obj => (string)obj["name"]!)
    .ToList();

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(filteredAndSorted, options));