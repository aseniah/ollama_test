using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var json = File.ReadAllText("input/data.json");
var array = JsonNode.Parse(json) as JsonArray;

var filtered = array
    .Select(n => (Record)n)
    .Where(r => r["active"].GetValue<bool>() && r["age"].GetValue<int>() >= 30)
    .OrderBy(r => r["name"].GetValue<string>())
    .ToList();

var result = new JsonArray();
foreach (var rec in filtered)
    result.Add(rec);

Console.Write(result.ToJsonString(new JsonSerializerOptions { WriteIndented = true }));