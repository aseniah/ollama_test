using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
var header = lines[0].Split(',');
var jsonArr = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var fields = lines[i].Split(',');
    if (fields.Length < 4) continue;
    
    var obj = new JsonObject
    {
        ["Name"] = fields[0].Trim(),
        ["Age"] = int.Parse(fields[1].Trim()),
        ["Email"] = fields[2].Trim(),
        ["Score"] = float.Parse(fields[3].Trim())
    };
    jsonArr.Add(obj);
}

Console.Write(jsonArr.ToJsonString());