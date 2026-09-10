using System;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Collections.Generic;

var lines = File.ReadAllLines("input/data.csv");

var array = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var fields = lines[i].Split(',');
    var obj = new JsonObject
    {
        ["Name"] = fields[0],
        ["Age"] = int.Parse(fields[1]),
        ["Email"] = fields[2],
        ["Score"] = double.Parse(fields[3])
    };
    array.Add(obj);
}

Console.Write(array.ToJsonString());