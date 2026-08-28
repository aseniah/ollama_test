using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;

string[] lines = File.ReadAllLines("input/data.csv");
string header = lines[0];
string[] dataLines = lines.Skip(1).ToArray();

var result = new List<object>();

foreach (string line in dataLines)
{
    if (string.IsNullOrWhiteSpace(line))
        continue;
    
    string[] parts = line.Split(',');
    if (parts.Length != 4)
        continue;
    
    string name = parts[0].Trim();
    int age = int.Parse(parts[1].Trim());
    string email = parts[2].Trim();
    float score = float.Parse(parts[3].Trim());
    
    result.Add(new { Name = name, Age = age, Email = email, Score = score });
}

var options = new JsonSerializerOptions
{
    WriteIndented = true,
    DefaultIgnoreCondition = System.Text.Json.Serialization.JsonIgnoreCondition.Never
};

string json = JsonSerializer.Serialize(result, options);
Console.WriteLine(json);