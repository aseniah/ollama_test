using System;
using System.Text.Json;

string[] lines = File.ReadAllLines("input/data.csv");
var jsonObjects = new System.Collections.Generic.List<System.Text.Json.Nodes.JsonObject>();

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;
    
    string[] parts = line.Split(',');
    if (parts.Length < 4) continue;
    
    string name = parts[0].Trim();
    int age = int.Parse(parts[1].Trim());
    string email = parts[2].Trim();
    double score = double.Parse(parts[3].Trim());
    
    var obj = new System.Text.Json.Nodes.JsonObject();
    obj["Name"] = name;
    obj["Age"] = age;
    obj["Email"] = email;
    obj["Score"] = score;
    
    jsonObjects.Add(obj);
}

var array = new System.Text.Json.Nodes.JsonArray();
foreach (var obj in jsonObjects)
{
    array.Add(obj);
}

string json = array.ToJsonString(new JsonSerializerOptions { WriteIndented = true });
Console.WriteLine(json);