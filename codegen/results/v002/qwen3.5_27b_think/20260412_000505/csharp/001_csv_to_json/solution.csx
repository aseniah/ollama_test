#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;

var lines = File.ReadAllLines("input/data.csv");

var records = new List<Dictionary<string, object>>();

for (int i = 1; i < lines.Length; i++)
{
    var parts = lines[i].Split(',');
    var record = new Dictionary<string, object>
    {
        ["Name"] = parts[0],
        ["Age"] = int.Parse(parts[1]),
        ["Email"] = parts[2],
        ["Score"] = float.Parse(parts[3])
    };
    records.Add(record);
}

var jsonOptions = new JsonSerializerOptions
{
    WriteIndented = false,
    PropertyNamingPolicy = JsonNamingPolicy.CamelCase
};

var json = JsonSerializer.Serialize(records, jsonOptions);
Console.Write(json);