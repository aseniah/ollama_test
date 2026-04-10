using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var jsonArray = new JsonArray();

try
{
    var lines = File.ReadAllLines("input/data.csv");
    
    for (int i = 1; i < lines.Length; i++)
    {
        var columns = lines[i].Split(',');
        
        if (columns.Length >= 4)
        {
            var obj = new JsonObject
            {
                ["Name"] = columns[0].Trim(),
                ["Age"] = int.Parse(columns[1].Trim()),
                ["Email"] = columns[2].Trim(),
                ["Score"] = float.Parse(columns[3].Trim())
            };
            
            jsonArray.Add(obj);
        }
    }
    
    Console.WriteLine(jsonArray.ToJsonString());
}
catch (Exception)
{
    // Handle file not found or parsing errors silently
}