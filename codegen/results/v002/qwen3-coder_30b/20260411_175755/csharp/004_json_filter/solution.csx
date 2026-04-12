using System.Text.Json;
using System.Text.Json.Nodes;
using System.Linq;

var jsonContent = File.ReadAllText("input/data.json");
var jsonArray = JsonNode.Parse(jsonContent).AsArray();

var filtered = jsonArray
    .Where(obj => obj["active"]?.GetValue<bool>() == true && obj["age"]?.GetValue<int>() >= 30)
    .OrderBy(obj => obj["name"]?.GetValue<string>())
    .ToArray();

Console.WriteLine(JsonSerializer.Serialize(filtered));