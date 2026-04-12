#r "System.Text.Json"
using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;

// Read CSV file
string[] lines = File.ReadAllLines("input/data.csv");

// Prepare list of objects to serialize
List<object> data = new();

// Skip header row, process each line
for (int i = 1; i < lines.Length; i++)
{
    string[] parts = lines[i].Split(',');
    
    data.Add(new {
        Name = parts[0],
        Age = int.Parse(parts[1]),
        Email = parts[2],
        Score = float.Parse(parts[3])
    });
}

// Output JSON array
string json = JsonSerializer.Serialize(data);
Console.Write(json);