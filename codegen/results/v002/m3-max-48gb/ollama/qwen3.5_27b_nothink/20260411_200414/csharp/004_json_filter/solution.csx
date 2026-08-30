using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var dataPath = "input/data.json";
var jsonContent = File.ReadAllText(dataPath);
var jsonArray = JsonNode.Parse(jsonContent)!.AsArray();

var filtered = new System.Collections.Generic.List<JsonObject>();

foreach (var node in jsonArray)
{
    var obj = node!.AsObject();
    var name = obj["name"]!.GetValue<string>();
    var age = obj["age"]!.GetValue<int>();
    var active = obj["active"]!.GetValue<bool>();

    if (active && age >= 30)
    {
        filtered.Add(obj);
    }
}

filtered.Sort((a, b) => string.Compare(a["name"]!.GetValue<string>(), b["name"]!.GetValue<string>(), StringComparison.Ordinal));

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(filtered, options));