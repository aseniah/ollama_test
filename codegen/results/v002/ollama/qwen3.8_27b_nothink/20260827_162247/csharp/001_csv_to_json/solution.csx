using System.Text.Json;
using System.Text.Json.Nodes;

string[] lines = File.ReadAllLines("input/data.csv");

var jsonArray = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;
    
    string[] parts = line.Split(',');
    if (parts.Length < 4) continue;
    
    string name = parts[0].Trim();
    int age = int.Parse(parts[1].Trim());
    string email = parts[2].Trim();
    float score = float.Parse(parts[3].Trim());
    
    var obj = new JsonObject
    {
        ["Name"] = name,
        ["Age"] = age,
        ["Email"] = email,
        ["Score"] = score
    };
    
    jsonArray.Add(obj);
}

var options = new JsonSerializerOptions { WriteIndented = false };
string json = jsonArray.ToJsonString(options);
Console.Write(json);