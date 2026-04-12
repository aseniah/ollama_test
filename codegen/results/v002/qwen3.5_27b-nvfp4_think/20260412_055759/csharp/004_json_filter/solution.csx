using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string content = File.ReadAllText("input/data.json");
JsonArray data = (JsonArray)JsonNode.Parse(content);

var filtered = data
    .Cast<JsonObject>()
    .Where(obj => obj["active"]?.GetValue<bool>() == true && obj["age"]?.GetValue<int>() >= 30)
    .OrderBy(obj => obj["name"]!.GetValue<string>());

JsonArray result = new JsonArray();
foreach (var item in filtered)
{
    result.Add(item);
}

Console.WriteLine(result.ToJsonString());