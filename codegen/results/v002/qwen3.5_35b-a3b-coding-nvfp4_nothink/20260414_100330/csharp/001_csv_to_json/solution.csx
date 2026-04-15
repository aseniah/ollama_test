using System;
using System.IO;
using System.Text.Json;

// Read the CSV file
var lines = File.ReadAllLines("input/data.csv");

// Skip the header
var dataLines = lines.Skip(1);

var results = new System.Text.Json.Nodes.JsonArray();

foreach (var line in dataLines)
{
    var parts = line.Split(',');
    if (parts.Length < 4) continue;

    var name = parts[0];
    var age = int.Parse(parts[1]);
    var email = parts[2];
    var score = float.Parse(parts[3]);

    var obj = new System.Text.Json.Nodes.JsonObject();
    obj["Name"] = name;
    obj["Age"] = age;
    obj["Email"] = email;
    obj["Score"] = score;

    results.Add(obj);
}

// Output JSON to stdout
Console.WriteLine(JsonSerializer.Serialize(results, new JsonSerializerOptions { WriteIndented = false }));