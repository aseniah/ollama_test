using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var filePath = "input/data.json";
if (!File.Exists(filePath))
{
    Console.WriteLine("[]");
    return;
}

var jsonContent = File.ReadAllText(filePath);
var root = JsonNode.Parse(jsonContent);

if (root is not JsonArray array)
{
    Console.WriteLine("[]");
    return;
}

var filteredList = new List<JsonElement>();

foreach (var item in array)
{
    if (item.TryGetPropertyValue("active", out var activeNode) && activeNode.GetValueKind() == JsonValueKind.True)
    {
        if (item.TryGetPropertyValue("age", out var ageNode) && ageNode.GetValueKind() == JsonValueKind.Number)
        {
            if (int.TryParse(ageNode.ToString(), out int age) && age >= 30)
            {
                filteredList.Add(item);
            }
        }
    }
}

var sortedList = filteredList.OrderBy(x => x.GetProperty("name").GetString());

var options = new JsonSerializerOptions
{
    WriteIndented = false,
    PropertyNamingPolicy = JsonNamingPolicy.CamelCase
};

Console.WriteLine(JsonSerializer.Serialize(sortedList, options));