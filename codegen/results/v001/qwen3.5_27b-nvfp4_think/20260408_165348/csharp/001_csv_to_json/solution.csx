#r "System.Text.Json"
using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var filePath = "input/data.csv";
var lines = File.ReadAllLines(filePath);

var result = new List<object>();

// Skip header (first line)
for (int i = 1; i < lines.Length; i++)
{
    var row = lines[i].Split(',');
    
    var obj = new { Name = "", Age = 0, Email = "", Score = 0f };
    
    if (row.Length >= 4)
    {
        obj.Name = row[0];
        obj.Age = int.Parse(row[1]);
        obj.Email = row[2];
        obj.Score = float.Parse(row[3], System.Globalization.CultureInfo.InvariantCulture);
    }
    
    result.Add(obj);
}

var jsonOptions = new JsonWriterOptions { Indented = false };
var jsonContent = JsonSerializer.Serialize(result, jsonOptions);

Console.Write(jsonContent);