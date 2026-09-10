using System.Text.Json;
using System.Text.Json.Nodes;

var json = JsonNode.Parse(File.ReadAllText("input/data.json"));
var array = json as JsonArray;

var filtered = array
    .Where(n => n["active"]?.GetValue<bool>() == true && n["age"]?.GetValue<int>() >= 30)
    .OrderBy(n => n["name"]?.GetValue<string>(), StringComparer.Ordinal)
    .Select(n => n)
    .ToList();

var result = new JsonArray();
foreach (var item in filtered)
    result.Add(item);

Console.Write(result.ToJsonString(new JsonSerializerOptions { WriteIndented = true }));