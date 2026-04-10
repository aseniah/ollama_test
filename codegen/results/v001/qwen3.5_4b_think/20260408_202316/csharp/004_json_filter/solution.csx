using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Collections.Generic;
using System.Linq;

string filePath = "input/data.json";

if (!File.Exists(filePath))
{
    throw new FileNotFoundException($"File {filePath} not found");
}

string content = File.ReadAllText(filePath);
JsonNode data = JsonNode.Parse(content);

var filtered = new List<JsonNode>();

if (data is JsonArray dataArray)
{
    foreach (JsonNode item in dataArray)
    {
        if (item is JsonObject obj)
        {
            bool active = obj["active"]?.GetValue<bool>() ?? false;
            int age = obj["age"]?.GetValue<int>() ?? 0;

            if (active && age >= 30)
            {
                filtered.Add(item);
            }
        }
    }
}

var filteredSorted = filtered.OrderBy(x => x["name"]?.GetValue<string>() ?? "").ToList();

JsonSerializerOptions options = new JsonSerializerOptions
{
    WriteIndented = false
};

Console.Write(JsonSerializer.Serialize(filteredSorted, options));