using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Globalization;

string[] lines = File.ReadAllLines("input/data.csv");
JsonArray result = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;
    
    string[] parts = line.Split(',');
    if (parts.Length >= 4)
    {
        JsonObject obj = new JsonObject();
        obj["Name"] = parts[0].Trim();
        obj["Age"] = int.Parse(parts[1].Trim());
        obj["Email"] = parts[2].Trim();
        obj["Score"] = float.Parse(parts[3].Trim(), CultureInfo.InvariantCulture);
        result.Add(obj);
    }
}

Console.WriteLine(JsonSerializer.Serialize(result));