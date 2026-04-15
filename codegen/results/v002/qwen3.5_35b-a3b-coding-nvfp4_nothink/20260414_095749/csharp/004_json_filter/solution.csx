using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string inputFile = "input/data.json";
string jsonContent = File.ReadAllText(inputFile);
JsonArray jsonArray = JsonNode.Parse(jsonContent)!.AsArray();

var filtered = jsonArray
    .Where(node =>
    {
        if (node is not JsonObject obj) return false;
        if (!obj.TryGetValue("active", out JsonNode? activeNode) || activeNode is not JsonValue activeValue) return false;
        if (!obj.TryGetValue("age", out JsonNode? ageNode) || ageNode is not JsonValue ageValue) return false;
        if (!obj.TryGetValue("name", out JsonNode? nameNode) || nameNode is not JsonValue nameValue) return false;

        bool isActive = activeValue.GetValue<bool>();
        int age = ageValue.GetValue<int>();
        string name = nameValue.GetValue<string>();

        return isActive && age >= 30;
    })
    .OrderBy(node => node!.AsObject()["name"]!.GetValue<string>())
    .ToList();

var resultArray = new JsonArray();
foreach (var node in filtered)
{
    resultArray.Add(node);
}

Console.WriteLine(JsonSerializer.Serialize(resultArray));