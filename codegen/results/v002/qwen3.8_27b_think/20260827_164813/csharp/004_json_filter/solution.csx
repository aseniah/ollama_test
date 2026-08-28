using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string json = File.ReadAllText("input/data.json");
var arr = JsonArray.Parse(json);

var filtered = arr
    .Select(o => (JsonNode)o)
    .Where(o => o["active"]!.GetValue<bool>() && o["age"]!.GetValue<int>() >= 30)
    .OrderBy(o => o["name"]!.GetValue<string>(), StringComparer.Ordinal)
    .ToList();

var result = new JsonArray();
foreach (var item in filtered)
{
    result.Add(item.DeepClone());
}

string output = result.ToJsonString(new JsonSerializerOptions { WriteIndented = true });
Console.WriteLine(output);