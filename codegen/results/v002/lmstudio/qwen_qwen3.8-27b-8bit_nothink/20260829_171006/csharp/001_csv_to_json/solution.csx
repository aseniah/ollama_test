using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Text.Json;

string filePath = "input/data.csv";
string[] lines = File.ReadAllLines(filePath);

List<object> results = new List<object>();

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;
    
    string[] parts = line.Split(',');
    
    string name = parts[0].Trim();
    int age = int.Parse(parts[1].Trim());
    string email = parts[2].Trim();
    float score = float.Parse(parts[3].Trim(), CultureInfo.InvariantCulture);
    
    var obj = new Dictionary<string, object>
    {
        { "Name", name },
        { "Age", age },
        { "Email", email },
        { "Score", score }
    };
    
    results.Add(obj);
}

var options = new JsonSerializerOptions
{
    WriteIndented = false
};

string json = JsonSerializer.Serialize(results, options);
Console.Write(json);