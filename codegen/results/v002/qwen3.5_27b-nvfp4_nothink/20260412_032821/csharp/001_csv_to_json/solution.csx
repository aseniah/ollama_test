#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Serialization;

var lines = File.ReadAllLines("input/data.csv");
var result = new List<object>();

// Skip header line (index 0)
for (int i = 1; i < lines.Length; i++)
{
    var parts = lines[i].Split(',');
    
    if (parts.Length < 4) continue;

    var name = parts[0];
    var age = int.Parse(parts[1]);
    var email = parts[2];
    var score = float.Parse(parts[3]);

    result.Add(new { Name = name, Age = age, Email = email, Score = score });
}

var options = new JsonSerializerOptions 
{ 
    WriteIndented = true,
    PropertyNamingPolicy = JsonNamingPolicy.CamelCase 
};

string json = JsonSerializer.Serialize(result, options);
Console.WriteLine(json);