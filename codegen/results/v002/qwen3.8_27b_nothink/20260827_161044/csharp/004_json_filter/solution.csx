using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string content = File.ReadAllText("input/data.json");
JsonArray arr = JsonArray.Parse(content);

var filtered = arr
    .Cast<JsonObject>()
    .Where(o => (bool)o["active"] == true && (int)o["age"] >= 30)
    .OrderBy(o => (string)o["name"])
    .ToList();

JsonArray result = new JsonArray();
foreach (var obj in filtered)
{
    result.Add((JsonNode)obj.DeepClone());
}

Console.Write(result.ToJsonString());