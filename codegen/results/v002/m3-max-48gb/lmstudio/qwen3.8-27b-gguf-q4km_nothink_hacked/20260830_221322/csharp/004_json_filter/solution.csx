using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string inputPath = "input/data.json";
string jsonContent = File.ReadAllText(inputPath);

var rootNode = JsonNode.Parse(jsonContent);
var arrayNode = rootNode as JsonArray;

var result = new JsonArray();

foreach (var item in arrayNode)
{
    var obj = item as JsonObject;
    if (obj == null) continue;
    
    var active = obj["active"]?.GetValue<bool>();
    var age = obj["age"]?.GetValue<int>();
    
    if (active == true && age.HasValue && age.Value >= 30)
    {
        result.Add(obj.DeepClone());
    }
}

// Sort by name ascending
var sorted = result.OrderBy(o => ((JsonObject)o)["name"]?.GetValue<string>() ?? "").ToList();

var outputArray = new JsonArray();
foreach (var obj in sorted)
{
    outputArray.Add(obj);
}

Console.WriteLine(JsonSerializer.Serialize(outputArray, new JsonSerializerOptions { WriteIndented = true }));