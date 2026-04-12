using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

string filePath = "input/data.csv";
if (!File.Exists(filePath))
{
    Console.WriteLine("[]");
    return;
}

var lines = File.ReadAllLines(filePath);
if (lines.Length < 2)
{
    Console.WriteLine("[]");
    return;
}

var result = new List<JsonNode>();

for (int i = 1; i < lines.Length; i++)
{
    var parts = lines[i].Split(',');
    if (parts.Length != 4) continue;

    string name = parts[0];
    int age = int.Parse(parts[1]);
    string email = parts[2];
    float score = float.Parse(parts[3]);

    var node = JsonNode.Parse("{\"Name\":\"\" + name.replace(\"\\\"\", \"\\\\\\\"\") + \"\",\"Age\":" + age + ",\"Email\":\"" + email.Replace("\"", "\\\"") + "\",\"Score\":" + score + "}");
    
    // More robust manual construction to ensure valid JSON without parsing a raw string
    var obj = new JsonObject();
    obj["Name"] = name;
    obj["Age"] = age;
    obj["Email"] = email;
    obj["Score"] = score;

    result.Add(obj);
}

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(result, options));