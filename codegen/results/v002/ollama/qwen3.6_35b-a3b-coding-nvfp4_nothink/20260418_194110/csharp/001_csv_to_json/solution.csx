using System;
using System.Collections.Generic;
using System.Text.Json;

var lines = File.ReadAllLines("input/data.csv");
var result = new List<Dictionary<string, object>>();

for (int i = 1; i < lines.Length; i++)
{
    if (string.IsNullOrWhiteSpace(lines[i])) continue;
    var parts = lines[i].Split(',');
    if (parts.Length < 4) continue;

    var obj = new Dictionary<string, object>
    {
        { "Name", parts[0].Trim() },
        { "Age", int.Parse(parts[1].Trim()) },
        { "Email", parts[2].Trim() },
        { "Score", float.Parse(parts[3].Trim()) }
    };

    result.Add(obj);
}

var options = new JsonSerializerOptions
{
    WriteIndented = false
};

Console.WriteLine(JsonSerializer.Serialize(result, options));