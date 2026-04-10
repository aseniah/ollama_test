using System.Text.Json;
using System.Text.Json.Nodes;

var json = System.IO.File.ReadAllText("input/data.json");
var array = JsonNode.Parse(json).AsArray();

var filtered = array
    .Where(item => item["active"].GetValue<bool>() && item["age"].GetValue<int>() >= 30)
    .OrderBy(item => item["name"].GetValue<string>())
    .ToList();

var result = new JsonArray();
foreach (var item in filtered)
{
    result.Add(item);
}

Console.WriteLine(JsonSerializer.Serialize(result));
