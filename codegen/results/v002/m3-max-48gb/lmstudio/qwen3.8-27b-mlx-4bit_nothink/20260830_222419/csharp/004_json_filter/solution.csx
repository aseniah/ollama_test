using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var jsonText = File.ReadAllText("input/data.json");
var arrayNode = JsonNode.Parse(jsonText);
var array = arrayNode as JsonArray;

var filtered = array
    .Where(x => (bool)(x["active"].GetValue<bool>()) && x["age"].GetValue<int>() >= 30)
    .OrderBy(x => (string)x["name"])
    .ToList();

var resultArray = new JsonArray();
foreach (var item in filtered)
{
    resultArray.Add(item.DeepClone());
}

Console.WriteLine(resultArray.ToJsonString(new JsonSerializerOptions { WriteIndented = false }));