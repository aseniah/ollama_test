using System;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Linq;
using System.Collections.Generic;

var fileContent = File.ReadAllText("input/data.json");
var root = JsonNode.Parse(fileContent);

if (root == null) {
    throw new JsonException("Failed to parse input data");
}

var dataArray = (JsonArray)root;
var filtered = new List<Dictionary<string, JsonNode>>();

foreach (var item in dataArray) {
    var obj = (JsonObject)item;
    var active = obj["active"]?.GetValue<bool>() ?? false;
    var age = obj["age"]?.GetValue<int>() ?? 0;

    if (active && age >= 30) {
        filtered.Add(obj);
    }
}

var sorted = filtered.OrderBy(o => o["name"]?.GetValue<string>() ?? "").ToList();
var result = new List<Dictionary<string, JsonNode>>(sorted.Count);

foreach (var record in sorted) {
    result.Add(record);
}

Console.WriteLine(JsonSerializer.Serialize(result));