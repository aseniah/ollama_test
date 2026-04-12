using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var jsonContent = File.ReadAllText("input/data.json");
var jsonArray = JsonNode.Parse(jsonContent)!.AsArray();

var filteredAndSorted = jsonArray
    .Where(node => node["active"]?.GetValue<bool>() == true && node["age"]?.GetValue<int>() >= 30)
    .OrderBy(node => node["name"]?.GetValue<string>())
    .ToArray();

Console.WriteLine(JsonSerializer.Serialize(filteredAndSorted));