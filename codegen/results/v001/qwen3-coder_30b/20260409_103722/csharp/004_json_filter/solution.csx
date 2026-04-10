using System.Text.Json;
using System.Text.Json.Nodes;
using System.Linq;

var json = File.ReadAllText("input/data.json");
var array = JsonNode.Parse(json).AsArray();

var filtered = array
    .Where(item => item["active"]?.AsBoolean() == true && item["age"]?.AsInt() >= 30)
    .OrderBy(item => item["name"]?.ToString())
    .ToArray();

Console.WriteLine(JsonSerializer.Serialize(filtered));