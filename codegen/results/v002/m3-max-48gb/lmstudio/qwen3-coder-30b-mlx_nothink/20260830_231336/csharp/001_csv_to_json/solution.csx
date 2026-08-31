using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
var jsonArray = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var values = lines[i].Split(',');
    
    var obj = new JsonObject
    {
        ["Name"] = values[0],
        ["Age"] = int.Parse(values[1]),
        ["Email"] = values[2],
        ["Score"] = float.Parse(values[3])
    };
    
    jsonArray.Add(obj);
}

Console.WriteLine(jsonArray.ToJsonString());