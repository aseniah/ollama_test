using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
var header = lines[0].Split(',');
var objects = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;
    
    var values = line.Split(',');
    var obj = new JsonObject();
    
    // Map by header index
    for (int j = 0; j < header.Length; j++)
    {
        var key = header[j].Trim();
        var value = values[j].Trim();
        
        switch (key)
        {
            case "Name":
                obj[key] = value;
                break;
            case "Age":
                obj[key] = int.Parse(value);
                break;
            case "Email":
                obj[key] = value;
                break;
            case "Score":
                obj[key] = float.Parse(value);
                break;
            default:
                obj[key] = value;
                break;
        }
    }
    
    objects.Add(obj);
}

var options = new JsonSerializerOptions
{
    WriteIndented = false,
    Encoder = System.Text.Encodings.Web.JavaScriptEncoder.UnsafeRelaxedJsonEscaping
};

Console.WriteLine(objects.ToJsonString(options));