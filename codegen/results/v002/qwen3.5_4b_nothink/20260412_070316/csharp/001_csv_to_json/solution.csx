using System;
using System.IO;
using System.Text.Json.Nodes;

var csvPath = "input/data.csv";
var lines = File.ReadAllLines(csvPath);
if (lines.Length <= 1)
{
    Console.WriteLine("[]");
    return;
}

var objects = new JsonNode[lines.Length - 1];

for (int i = 1; i < lines.Length; i++)
{
    var row = lines[i].Split(',');
    var name = row[0].Trim();
    var age = int.Parse(row[1].Trim());
    var email = row[2].Trim();
    var score = float.Parse(row[3].Trim());

    var obj = new JsonNode();
    obj["Name"] = name;
    obj["Age"] = age;
    obj["Email"] = email;
    obj["Score"] = score;
    objects[i - 1] = obj;
}

Console.WriteLine($"[{string.Join(", ", objects.Select(o => o.ToJsonString()).ToString())}]");