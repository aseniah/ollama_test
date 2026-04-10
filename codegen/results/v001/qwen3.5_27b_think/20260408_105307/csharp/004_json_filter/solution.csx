using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var inputFile = "input/data.json";
var jsonData = File.ReadAllText(inputFile);
var root = JsonNode.Parse(jsonData)!;
var array = root.AsArray();

var filtered = array
    .Where(x => x["active"]!.GetValue<bool>() && x["age"]!.GetValue<int>() >= 30)
    .OrderBy(x => x["name"]!.GetValue<string>())
    .ToArray();

var jsonArray = new JsonArray();
foreach (var item in filtered)
{
    jsonArray.Add(item);
}

var options = new JsonSerializerOptions { WriteIndented = false };
Console.Write(jsonArray.ToJsonString(options));