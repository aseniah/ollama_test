using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var json = File.ReadAllText("input/data.json");
var arr = JsonArray.Parse(json);

var filtered = new List<JsonObject>();
foreach (var item in arr)
{
    var obj = (JsonObject)item;
    var active = obj["active"]?.GetValue<bool>() ?? false;
    var age = obj["age"]?.GetValue<int>() ?? 0;
    if (active && age >= 30)
    {
        filtered.Add(obj);
    }
}

var sorted = filtered.OrderBy(o => (string)o["name"]).ToList();

var output = new JsonArray();
foreach (var obj in sorted)
{
    output.Add(obj.DeepClone());
}

Console.WriteLine(output.ToJsonString(new JsonSerializerOptions { WriteIndented = false }));