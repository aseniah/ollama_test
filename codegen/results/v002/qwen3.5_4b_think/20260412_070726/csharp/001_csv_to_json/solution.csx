using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

string path = "input/data.csv";
var lines = File.ReadAllLines(path);

var results = new List<Dictionary<string, object>>();

// Skip header line (index 0)
for (int i = 1; i < lines.Length; i++)
{
    string[] parts = lines[i].Split(',');
    
    var obj = new Dictionary<string, object>
    {
        { "Name", parts[0] },
        { "Age", int.Parse(parts[1]) },
        { "Email", parts[2] },
        { "Score", float.Parse(parts[3]) }
    };
    
    results.Add(obj);
}

// Serialize to JSON string
string json = JsonSerializer.Serialize(results, new JsonSerializerOptions { WriteIndented = false });

// Output to stdout
System.Console.WriteLine(json);