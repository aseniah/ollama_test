using System;
using System.IO;
using System.Collections.Generic;
using System.Text.Json;
using System.Globalization;

var filePath = "input/data.csv";
var lines = File.ReadAllLines(filePath);

var results = new List<object>();

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;
    
    var parts = line.Split(',');
    if (parts.Length >= 4)
    {
        var person = new
        {
            Name = parts[0].Trim(),
            Age = int.Parse(parts[1].Trim(), CultureInfo.InvariantCulture),
            Email = parts[2].Trim(),
            Score = float.Parse(parts[3].Trim(), CultureInfo.InvariantCulture)
        };
        results.Add(person);
    }
}

var options = new JsonSerializerOptions
{
    PropertyNamingPolicy = null
};

Console.WriteLine(JsonSerializer.Serialize(results, options));