using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var content = File.ReadAllText("input/data.json");
var doc = JsonDocument.Parse(content);
var elements = doc.RootElement.EnumerateArray();

var filtered = elements
    .Where(e => e.GetProperty("active").GetBoolean() && e.GetProperty("age").GetInt32() >= 30)
    .OrderBy(e => e.GetProperty("name").GetString());

var result = new JsonArray();
foreach (var item in filtered)
{
    result.Add(JsonNode.Parse(item.GetRawText()));
}

Console.WriteLine(JsonSerializer.Serialize(result, new JsonSerializerOptions { WriteIndented = false }));