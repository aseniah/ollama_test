#r "System.Text.Json"

using System;
using System.IO;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

var csvContent = File.ReadAllText("input/input.csv");
var lines = csvContent.Split(new[] { '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

if (lines.Length == 0)
{
    Console.WriteLine("[]");
    return;
}

var headers = lines[0].Split(',');
var result = new List<object>();
var referenceDate = new DateTime(2025, 7, 1);

for (int i = 1; i < lines.Length; i++)
{
    if (string.IsNullOrWhiteSpace(lines[i])) continue;
    
    var values = lines[i].Split(',');
    var record = new JsonObject();
    
    for (int j = 0; j < headers.Length && j < values.Length; j++)
    {
        var key = headers[j].Trim().ToLowerInvariant();
        var value = values[j].Trim();
        
        // Check if this is a date field and calculate age
        if (key.Contains("birth") && key.Contains("date"))
        {
            if (DateTime.TryParse(value, out var birthDate))
            {
                int age = referenceDate.Year - birthDate.Year;
                if (referenceDate.DayOfYear < birthDate.DayOfYear)
                {
                    age--;
                }
                record["age"] = age;
            }
        }
        
        // Try to parse numeric values
        if (int.TryParse(value, out var intValue))
        {
            record[key] = intValue;
        }
        else if (double.TryParse(value, out var doubleValue))
        {
            record[key] = doubleValue;
        }
        else if (bool.TryParse(value, out var boolValue))
        {
            record[key] = boolValue;
        }
        else
        {
            record[key] = value;
        }
    }
    
    result.Add(record);
}

var options = new JsonSerializerOptions
{
    WriteIndented = true,
    PropertyNamingPolicy = JsonNamingPolicy.CamelCase
};

Console.WriteLine(JsonSerializer.Serialize(result, options));