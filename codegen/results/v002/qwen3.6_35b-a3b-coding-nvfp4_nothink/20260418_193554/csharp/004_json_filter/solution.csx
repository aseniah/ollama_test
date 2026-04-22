using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var jsonText = File.ReadAllText("input/data.json");
var array = JsonNode.Parse(jsonText)!.AsArray();

var results = new List<object>();

foreach (var node in array)
{
    if (node == null) continue;
    
    var obj = node.AsObject();
    
    if (!obj.ContainsKey("active") || !obj.ContainsKey("age") || !obj.ContainsKey("name"))
        continue;
    
    bool active = obj["active"]!.GetValue<bool>();
    int age = obj["age"]!.GetValue<int>();
    string name = obj["name"]!.GetValue<string>();
    
    if (active && age >= 30)
    {
        results.Add(new { name, age, active, score = obj["score"]!.GetValue<double>() });
    }
}

results.Sort((a, b) => string.Compare(a.name, b.name, StringComparison.Ordinal));

var outputArray = new JsonArray();
foreach (var item in results)
{
    var objNode = new JsonObject();
    objNode["name"] = item.name;
    objNode["age"] = item.age;
    objNode["active"] = item.active;
    objNode["score"] = item.score;
    outputArray.Add(objNode);
}

var options = new JsonSerializerOptions
{
    WriteIndented = true
};

Console.WriteLine(outputArray.ToJsonString(options));