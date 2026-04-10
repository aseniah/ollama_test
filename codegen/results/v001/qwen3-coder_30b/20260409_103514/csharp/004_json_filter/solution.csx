using System.Text.Json;
using System.Text.Json.Nodes;
using System.Linq;

var json = File.ReadAllText("input/data.json");
var array = JsonNode.Parse(json).AsArray();

var filtered = array
    .Where(obj => obj["active"]?.GetValue<bool>() == true && obj["age"]?.GetValue<int>() >= 30)
    .OrderBy(obj => obj["name"]?.GetValue<string>())
    .ToArray();

Console.WriteLine(JsonSerializer.Serialize(filtered));