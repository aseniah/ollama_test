using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var jsonData = File.ReadAllText("input/data.json");
var jsonArray = JsonNode.Parse(jsonData)!.AsArray();

var filtered = jsonArray
    .Where(item => item["active"]?.GetValue<bool>() == true && item["age"]?.GetValue<int>() >= 30)
    .OrderBy(item => item["name"]?.GetValue<string>())
    .ToArray();

Console.WriteLine(JsonSerializer.Serialize(filtered));