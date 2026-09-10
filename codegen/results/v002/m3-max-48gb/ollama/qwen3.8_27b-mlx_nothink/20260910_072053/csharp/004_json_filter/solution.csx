using System.Text.Json;
using System.Text.Json.Nodes;

var json = File.ReadAllText("input/data.json");
var array = JsonNode.Parse(json) as JsonArray;

var result = new JsonArray();

foreach (var item in array)
{
    var obj = item as JsonObject;
    if (obj == null) continue;
    
    var active = obj["active"]?.GetValue<bool>() ?? false;
    var age = obj["age"]?.GetValue<int>() ?? 0;
    
    if (active && age >= 30)
    {
        result.Add(obj.DeepClone());
    }
}

// Sort by name ascending
var sorted = result.OrderBy(n => (n as JsonObject)["name"]?.GetValue<string>() ?? "").ToArray();

var output = new JsonArray();
foreach (var item in sorted)
{
    output.Add(item);
}

Console.Write(output.ToJsonString());