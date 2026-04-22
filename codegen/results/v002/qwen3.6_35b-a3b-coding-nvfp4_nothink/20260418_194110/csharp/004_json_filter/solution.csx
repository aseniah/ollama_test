using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string json = File.ReadAllText("input/data.json");
JsonArray array = JsonNode.Parse(json)!.AsArray();

var results = new List<JsonObject>();

foreach (var item in array)
{
    if (item == null) continue;
    
    JsonObject obj = item.AsObject();
    
    if (obj.TryGetValue("active", out JsonNode? activeNode) &&
        obj.TryGetValue("age", out JsonNode? ageNode) &&
        obj.TryGetValue("name", out JsonNode? nameNode))
    {
        bool active = activeNode!.GetValue<bool>();
        int age = ageNode!.GetValue<int>();
        string name = nameNode!.GetValue<string>();
        
        if (active && age >= 30)
        {
            results.Add(obj);
        }
    }
}

results.Sort((a, b) =>
{
    string nameA = a["name"]!.GetValue<string>();
    string nameB = b["name"]!.GetValue<string>();
    return string.Compare(nameA, nameB, StringComparison.OrdinalIgnoreCase);
});

var output = new JsonArray();
foreach (var obj in results)
{
    output.Add(obj.Clone());
}

Console.WriteLine(output.ToJsonString());