using System;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Linq;

var data = File.ReadAllText("input/data.json");
var records = JsonNode.Parse(data) as JsonArray;

var filtered = records
    ?.Select(r => {
        var name = r["name"]?.ToString();
        var age = r["age"]?.GetValue<int>();
        var active = r["active"]?.GetValue<bool>();
        var score = r["score"]?.GetValue<double>();
        return new {
            name = name,
            age = age,
            active = active,
            score = score
        };
    })
    .Where(r => r != null && r.active == true && r.age >= 30)
    .OrderBy(r => r.name)
    .Select(r => new {
        name = r.name,
        age = r.age,
        active = r.active,
        score = r.score
    })
    .ToList();

Console.WriteLine(JsonSerializer.Serialize(filtered, new JsonSerializerOptions {
    WriteIndented = false,
    PropertyNamingPolicy = JsonNamingPolicy.CamelCase
}))