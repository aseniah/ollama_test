using System;
using System.Collections.Generic;
using System.Text.Json;
using System.IO;

var lines = File.ReadAllLines("input/data.csv");
var results = new List<Dictionary<string, object>>();

// Skip header
for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i].Trim();
    if (string.IsNullOrWhiteSpace(line)) continue;
    
    var parts = line.Split(',');
    if (parts.Length >= 4)
    {
        var name = parts[0].Trim();
        var age = int.Parse(parts[1].Trim());
        var email = parts[2].Trim();
        var score = float.Parse(parts[3].Trim());
        
        results.Add(new Dictionary<string, object> 
        {
            { "Name", name },
            { "Age", age },
            { "Email", email },
            { "Score", score }
        });
    }
}

var json = JsonSerializer.Serialize(results);
Console.WriteLine(json);