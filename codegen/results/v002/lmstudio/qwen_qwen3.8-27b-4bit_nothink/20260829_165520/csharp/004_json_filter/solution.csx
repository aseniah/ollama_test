using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var jsonText = File.ReadAllText("input/data.json");
var root = JsonNode.Parse(jsonText);

if (root == null)
{
    Console.WriteLine("[]");
    return;
}

var array = root as JsonArray;
if (array == null)
{
    Console.WriteLine("[]");
    return;
}

var filtered = new List<JsonNode>();
foreach (var item in array)
{
    if (item is JsonObject obj)
    {
        var active = obj["active"]?.GetValue<bool>() ?? false;
        var age = obj["age"]?.GetValue<int>() ?? 0;
        
        if (active && age >= 30)
        {
            filtered.Add(obj);
        }
    }
}

// Sort by name ascending
var sorted = filtered.OrderBy(x => 
{
    var name = (x as JsonObject)?["name"]?.GetValue<string>() ?? "";
    return name;
}).ToList();

var resultArray = new JsonArray();
foreach (var item in sorted)
{
    resultArray.Add(item);
}

var options = new JsonSerializerOptions
{
    WriteIndented = true
};

var output = JsonSerializer.Serialize(resultArray, options);
Console.WriteLine(output);