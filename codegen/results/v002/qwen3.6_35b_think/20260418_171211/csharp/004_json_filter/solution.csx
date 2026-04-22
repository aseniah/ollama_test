using System;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.IO;

var json = JsonNode.Parse(File.ReadAllText("input/data.json"));
var result = json.AsArray()
    .OfType<JsonObject>()
    .Where(obj => 
    {
        var active = obj["active"]?.GetValue<bool>();
        var age = obj["age"]?.GetValue<int>();
        return active == true && age >= 30;
    })
    .OrderBy(obj => obj["name"]?.GetValue<string>())
    .ToArray();

Console.WriteLine(JsonSerializer.Serialize(result));