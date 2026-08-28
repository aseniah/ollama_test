using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var jsonContent = File.ReadAllText("input/data.json");
var json = JsonNode.Parse(jsonContent);

var result = json
    .AsArray()
    .Where(node =>
    {
        var active = node["active"];
        var age = node["age"];
        if (active == null || age == null) return false;
        
        return active.GetValue<bool>() && age.GetValue<int>() >= 30;
    })
    .OrderBy(node => node["name"]?.GetValue<string>() ?? "")
    .ToList();

var output = new JsonArray();
foreach (var item in result)
{
    output.Add(item);
}

Console.WriteLine(output.ToJsonString());