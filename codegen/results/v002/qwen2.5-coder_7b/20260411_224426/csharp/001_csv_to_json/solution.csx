using System;
using System.Collections.Generic;

var lines = File.ReadAllLines("input/data.csv");
var header = lines[0].Split(',');
var data = new List<Dictionary<string, object>>();

for (int i = 1; i < lines.Length; i++)
{
    var row = lines[i].Split(',');
    var obj = new Dictionary<string, object>
    {
        { "Name", row[0] },
        { "Age", int.Parse(row[1]) },
        { "Email", row[2] },
        { "Score", float.Parse(row[3]) }
    };
    data.Add(obj);
}

var json = JsonSerializer.Serialize(data, new JsonSerializerOptions { WriteIndented = true });
Console.WriteLine(json);