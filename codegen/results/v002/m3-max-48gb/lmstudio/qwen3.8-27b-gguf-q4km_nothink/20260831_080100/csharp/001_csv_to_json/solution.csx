using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;

var lines = File.ReadAllLines("input/data.csv");

var result = new List<Dictionary<string, object>>();

for (int i = 1; i < lines.Length; i++)
{
    var parts = lines[i].Split(',');
    if (parts.Length < 4) continue;
    
    var obj = new Dictionary<string, object>
    {
        { "Name", parts[0] },
        { "Age", int.Parse(parts[1]) },
        { "Email", parts[2] },
        { "Score", float.Parse(parts[3]) }
    };
    result.Add(obj);
}

var jsonOptions = new JsonSerializerOptions
{
    WriteIndented = false
};

var json = JsonSerializer.Serialize(result, jsonOptions);
Console.WriteLine(json);