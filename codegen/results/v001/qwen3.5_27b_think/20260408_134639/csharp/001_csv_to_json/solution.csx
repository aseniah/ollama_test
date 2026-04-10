#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

string filePath = "input/data.csv";

if (!File.Exists(filePath))
{
    Console.WriteLine("[]");
    Environment.Exit(1);
}

string[] lines = File.ReadAllLines(filePath);

if (lines.Length < 2)
{
    Console.WriteLine("[]");
    Environment.Exit(0);
}

var records = new List<object>();

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    
    // Handle quoted values and commas in fields
    var fields = ParseCSVLine(line);
    
    if (fields.Count < 4)
        continue;
    
    if (int.TryParse(fields[1], out int age) &&
        float.TryParse(fields[3], out float score))
    {
        var record = new Dictionary<string, object>
        {
            { "Name", fields[0] },
            { "Age", age },
            { "Email", fields[2] },
            { "Score", score }
        };
        records.Add(record);
    }
}

var options = new JsonSerializerOptions
{
    WriteIndented = false,
    PropertyNamingPolicy = JsonNamingPolicy.CamelCase
};

Console.WriteLine(JsonSerializer.Serialize(records, options));

static List<string> ParseCSVLine(string line)
{
    var fields = new List<string>();
    var current = new System.Text.StringBuilder();
    bool inQuotes = false;
    
    foreach (char c in line)
    {
        if (c == '"')
        {
            inQuotes = !inQuotes;
        }
        else if (c == ',' && !inQuotes)
        {
            fields.Add(current.ToString().Trim('"'));
            current.Clear();
        }
        else
        {
            current.Append(c);
        }
    }
    
    fields.Add(current.ToString().Trim('"'));
    return fields;
}