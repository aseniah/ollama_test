using System;
using System.IO;
using System.Text.Json;

string[] lines = File.ReadAllLines("input/data.csv");
JsonArray jsonArray = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    string[] parts = lines[i].Split(',');
    if (parts.Length == 4)
    {
        JsonObject jsonObject = new JsonObject
        {
            ["Name"] = new JsonValue(parts[0]),
            ["Age"] = new JsonValue(int.Parse(parts[1])),
            ["Email"] = new JsonValue(parts[2]),
            ["Score"] = new JsonValue(float.Parse(parts[3]))
        };
        jsonArray.Add(jsonObject);
    }
}

Console.WriteLine(jsonArray.ToJsonString());