#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var fileName = "input/data.json";
var jsonContent = File.ReadAllText(fileName);
var parsed = JsonArray.Parse(jsonContent);

var filtered = parsed
    .Where(obj =>
    {
        var activeNode = obj["active"];
        var ageNode = obj["age"];
        
        if (activeNode == null || ageNode == null) return false;
        
        bool isActive = activeNode.GetValue<bool>();
        int age = ageNode.GetValue<int>();
        
        return isActive && age >= 30;
    })
    .OrderBy(obj => obj["name"]?.ToString() ?? "")
    .ToArray();

var outputJson = new JsonArray(filtered);
Console.WriteLine(outputJson.ToJsonString());