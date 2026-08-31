using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

string input = File.ReadAllText("input/data.json");
JsonNode root = JsonNode.Parse(input);
var array = root as JsonArray;

var filtered = new JsonArray();
foreach (var node in array)
{
    var obj = (JsonObject)node;
    bool active = obj["active"].GetValue<bool>();
    int age = obj["age"].GetValue<int>();
    if (active && age >= 30)
    {
        filtered.Add(obj);
    }
}

// Sort by name ascending
filtered.Sort((a, b) => ((JsonObject)a)["name"].GetValue<string>().CompareTo(((JsonObject)b)["name"].GetValue<string>()));

string output = filtered.ToJsonString();
Console.Write(output);