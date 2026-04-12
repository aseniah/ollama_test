using System;
using System.IO;

var lines = File.ReadAllLines("input/data.csv");
var header = lines[0].Split(',');
var data = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var row = lines[i].Split(',');
    var jsonObject = new JsonObject
    {
        { "Name", row[0] },
        { "Age", int.Parse(row[1]) },
        { "Email", row[2] },
        { "Score", float.Parse(row[3]) }
    };
    data.Add(jsonObject);
}

Console.WriteLine(data.ToString());