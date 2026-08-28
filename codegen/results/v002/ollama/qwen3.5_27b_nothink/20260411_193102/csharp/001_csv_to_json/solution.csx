#r "nuget: System.Text.Json"
using System;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Collections.Generic;
using System.IO;

var lines = File.ReadAllLines("input/data.csv");
var jsonArray = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var parts = lines[i].Split(',');
    if (parts.Length < 4) continue;

    var row = new JsonObject();
    row["Name"] = parts[0];
    row["Age"] = int.Parse(parts[1]);
    row["Email"] = parts[2];
    row["Score"] = float.Parse(parts[3]);
    
    jsonArray.Add(row);
}

Console.WriteLine(jsonArray.ToJsonString());