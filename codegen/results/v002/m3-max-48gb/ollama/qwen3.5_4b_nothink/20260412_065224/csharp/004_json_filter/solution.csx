using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var data = JsonNode.Parse(File.ReadAllText("input/data.json")) as JsonArray;
if (data == null) throw new InvalidOperationException("Failed to read JSON data.");

var filtered = data
    .Cast<JsonElement>()
    .Where(element => element.GetProperty("active").GetValue<bool>())
    .Where(element => element.GetProperty("age").GetValue<int>() >= 30)
    .OrderBy(element => element.GetProperty("name").GetValue<string>())
    .ToList();

var result = filtered.Select(element => new {
    name = element.GetProperty("name").GetString(),
    age = element.GetProperty("age").GetInt32(),
    active = element.GetProperty("active").GetValue<bool>(),
    score = element.GetProperty("score").GetDouble()
}).ToList();

Console.WriteLine(JsonSerializer.Serialize(result));