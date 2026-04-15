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
JsonArray data = JsonNode.Parse(jsonString)!.AsArray();

var filteredAndSorted = data
    .Select(node => node!.AsObject())
    .Where(obj => 
        (bool)obj["active"]! && 
        (int)obj["age"]! >= 30)
    .OrderBy(obj => (string)obj["name"]!)
    .ToList();

var options = new JsonSerializerOptions { WriteIndented = true };
string result = JsonSerializer.Serialize(filteredAndSorted, options);

Console.WriteLine(result);