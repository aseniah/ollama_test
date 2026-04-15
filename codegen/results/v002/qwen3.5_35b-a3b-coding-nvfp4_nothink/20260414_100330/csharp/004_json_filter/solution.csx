using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string json = File.ReadAllText("input/data.json");
JsonArray arr = JsonArray.Parse(json);

var filtered = new List<JsonObject>();

foreach (var item in arr)
{
    if (item is JsonObject obj)
    {
        if (obj.TryGetValue("active", out JsonNode? activeNode) &&
            obj.TryGetValue("age", out JsonNode? ageNode))
        {
            bool isActive = activeNode is JsonValue val && val.GetValue<bool>();
            int age = ageNode is JsonValue ageVal && ageVal.TryGetValue<int>(out int a) ? a : -1;

            if (isActive && age >= 30)
            {
                filtered.Add(obj);
            }
        }
    }
}

var sorted = filtered.OrderBy(o => o["name"]!.GetValue<string>()).ToList();

var result = new JsonObject
{
    { "result", new JsonArray(sorted) }
};

// Actually, the requirement says output a JSON array, not an object with a "result" key.
// Let's output just the array directly.
var outputArray = new JsonArray();
foreach (var obj in sorted)
{
    outputArray.Add(obj);
}

Console.WriteLine(JsonSerializer.Serialize(outputArray));