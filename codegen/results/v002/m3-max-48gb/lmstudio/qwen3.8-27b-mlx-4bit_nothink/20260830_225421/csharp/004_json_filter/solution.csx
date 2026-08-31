using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var jsonText = File.ReadAllText("input/data.json");
var root = JsonNode.Parse(jsonText);
var array = (JsonArray)root;

var records = new List<JsonObject>();
foreach (var item in array)
{
    var obj = (JsonObject)item;
    bool active = (bool)obj["active"]!.Value<bool>();
    int age = (int)obj["age"]!.Value<int>();
    if (active && age >= 30)
    {
        records.Add(obj);
    }
}

records = records.OrderBy(r => (string)r["name"]!.Value<string>()).ToList();

var result = new JsonArray();
foreach (var r in records)
{
    result.Add(r.DeepClone());
}

Console.Write(JsonSerializer.Serialize(result, new JsonSerializerOptions { WriteIndented = true }));