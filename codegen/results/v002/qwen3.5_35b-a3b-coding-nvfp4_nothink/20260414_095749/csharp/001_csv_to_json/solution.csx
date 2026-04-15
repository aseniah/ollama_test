using System;
using System.Text.Json;
using System.Text.Json.Nodes;

string[] lines = File.ReadAllLines("input/data.csv");
var jsonArray = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    string[] parts = lines[i].Split(',');
    if (parts.Length < 4) continue;

    var obj = new JsonObject
    {
        ["Name"] = parts[0],
        ["Age"] = JsonValue.Create(int.Parse(parts[1])),
        ["Email"] = parts[2],
        ["Score"] = JsonValue.Create(float.Parse(parts[3]))
    };

    jsonArray.Add(obj);
}

Console.WriteLine(JsonSerializer.Serialize(jsonArray));