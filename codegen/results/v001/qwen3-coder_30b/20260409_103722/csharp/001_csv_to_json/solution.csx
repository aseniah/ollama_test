using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var data = new List<JsonObject>();
var lines = File.ReadAllLines("input/data.csv");

// Skip header row
for (int i = 1; i < lines.Length; i++)
{
    var columns = lines[i].Split(',');
    
    var obj = new JsonObject
    {
        ["Name"] = columns[0].Trim(),
        ["Age"] = int.Parse(columns[1].Trim()),
        ["Email"] = columns[2].Trim(),
        ["Score"] = float.Parse(columns[3].Trim())
    };
    
    data.Add(obj);
}

var array = new JsonArray(data.ToArray());
Console.WriteLine(array.ToJsonString());