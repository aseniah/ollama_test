using System.Text.Json;
using System.Text.Json.Nodes;

var json = File.ReadAllText("input/data.json");
var arr = JsonNode.Parse(json).AsArray();

var filtered = new JsonArray();

foreach (var item in arr)
{
    var obj = item.AsObject();
    bool active = (bool)obj["active"]!;
    int age = (int)obj["age"]!;

    if (active && age >= 30)
    {
        filtered.Add((JsonNode)item.DeepClone());
    }
}

// Sort by name ascending
var sorted = filtered.OrderByDescending(x => (string)((JsonObject)x)["name"]!)
                     .OrderBy(x => (string)((JsonObject)x)["name"]!)
                     .ToArray();

var sortedArr = new JsonArray();
foreach (var item in sorted)
{
    sortedArr.Add((JsonNode)item.DeepClone());
}

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(sortedArr.ToJsonString(options));