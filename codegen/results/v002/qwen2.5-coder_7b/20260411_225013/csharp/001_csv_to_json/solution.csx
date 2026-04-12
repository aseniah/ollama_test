using System;
using System.Collections.Generic;
using System.Linq;

var lines = File.ReadAllLines("input/data.csv");
var header = lines[0].Split(',');
var data = lines.Skip(1).Select(line => line.Split(',')).ToList();

var jsonArray = new List<Dictionary<string, object>>();

foreach (var row in data)
{
    var jsonItem = new Dictionary<string, object>
    {
        { "Name", row[0] },
        { "Age", int.Parse(row[1]) },
        { "Email", row[2] },
        { "Score", float.Parse(row[3]) }
    };

    jsonArray.Add(jsonItem);
}

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(jsonArray, options));