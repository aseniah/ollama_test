using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var jsonContent = File.ReadAllText("input/data.json");
var jsonArray = JsonArray.Parse(jsonContent);

var filtered = new JsonArray();

foreach (var item in jsonArray)
{
    var obj = (JsonObject)item;
    
    var name = obj["name"]?.GetValue<string>();
    var age = obj["age"]?.GetValue<int>();
    var active = obj["active"]?.GetValue<bool>();
    
    if (active == true && age.HasValue && age.Value >= 30)
    {
        filtered.Add(obj);
    }
}

// Sort by name ascending
var sorted = filtered
    .Select(j => (JsonObject)j)
    .OrderBy(o => o["name"]?.GetValue<string>())
    .ToList();

var resultArray = new JsonArray();
foreach (var obj in sorted)
{
    resultArray.Add(obj);
}

var options = new JsonSerializerOptions
{
    WriteIndented = false
};

Console.Write(resultArray.ToJsonString(options));