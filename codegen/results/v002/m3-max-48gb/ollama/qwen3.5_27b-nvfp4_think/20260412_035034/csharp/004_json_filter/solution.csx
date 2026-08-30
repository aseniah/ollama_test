using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

var jsonContent = File.ReadAllText("input/data.json");
var root = JsonNode.Parse(jsonContent);

if (root is JsonArray array) {
    var filtered = array
        .OfType<JsonObject>()
        .Where(node => node["active"]?.GetValue<bool>() == true && node["age"]?.GetValue<int>() >= 30)
        .OrderBy(node => node["name"]?.GetValue<string>())
        .ToList();

    var resultArray = new JsonArray();
    foreach (var item in filtered) {
        resultArray.Add(item);
    }

    Console.WriteLine(resultArray.ToJsonString());
}