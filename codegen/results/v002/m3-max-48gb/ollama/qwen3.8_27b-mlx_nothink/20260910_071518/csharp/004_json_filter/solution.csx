using System.Text.Json;
using System.Text.Json.Nodes;

var input = File.ReadAllText("input/data.json");
var array = JsonNode.Parse(input) as JsonArray;

var filtered = new JsonArray();

foreach (var item in array)
{
    var obj = item as JsonObject;
    var active = obj["active"]?.GetValue<bool>() ?? false;
    var age = obj["age"]?.GetValue<int>() ?? 0;
    
    if (active && age >= 30)
    {
        filtered.Add(obj.DeepClone());
    }
}

// Sort by name ascending
var sorted = filtered.Select(n => n as JsonObject).OrderBy(o => o["name"]?.GetValue<string>(), StringComparer.Ordinal).ToList();

var result = new JsonArray();
foreach (var item in sorted)
{
    result.Add(item);
}

Console.Write(result.ToJsonString(new JsonSerializerOptions { WriteIndented = false }));