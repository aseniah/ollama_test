using System;
using System.Text.Json;
using System.Text.Json.Nodes;

var json = File.ReadAllText("input/data.json");
var root = JsonNode.Parse(json);

if (root is JsonArray array)
{
    var filtered = array
        .Where(node => node is JsonObject obj)
        .Select(node => (JsonObject)node)
        .Where(obj =>
        {
            if (!obj.TryGetValue("active", out var activeNode) || activeNode is not JsonValue activeVal)
                return false;
            if (!obj.TryGetValue("age", out var ageNode) || ageNode is not JsonValue ageVal)
                return false;
            
            return activeVal.GetValue<bool>() && ageVal.GetValue<int>() >= 30;
        })
        .OrderBy(obj => obj["name"]?.GetValue<string>() ?? "")
        .Select(obj => JsonNode.Parse(JsonSerializer.Serialize(obj)));

    var result = JsonArray.Create(filtered.ToArray());
    Console.WriteLine(JsonSerializer.Serialize(result));
}