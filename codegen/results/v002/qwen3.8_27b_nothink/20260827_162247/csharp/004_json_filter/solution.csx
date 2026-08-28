using System.Text.Json;
using System.Text.Json.Nodes;

string json = File.ReadAllText("input/data.json");
JsonNode root = JsonNode.Parse(json);

if (root is not JsonArray array)
{
    Console.WriteLine("[]");
    return;
}

var filtered = new JsonArray();
foreach (JsonNode item in array)
{
    if (item is not JsonObject obj) continue;
    
    bool active = obj["active"]?.GetValue<bool>() ?? false;
    int age = obj["age"]?.GetValue<int>() ?? 0;
    
    if (active && age >= 30)
    {
        filtered.Add(obj.DeepClone());
    }
}

// Sort by name ascending
var sorted = filtered
    .Where(x => x is JsonObject)
    .OrderBy(x => (x as JsonObject)["name"]?.GetValue<string>() ?? "")
    .Select(x => x.DeepClone())
    .ToList();

var resultArray = new JsonArray();
foreach (var item in sorted)
{
    resultArray.Add(item);
}

Console.WriteLine(resultArray.ToJsonString(new JsonSerializerOptions { WriteIndented = false }));