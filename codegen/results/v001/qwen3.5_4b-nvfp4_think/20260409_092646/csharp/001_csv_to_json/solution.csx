using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;

var lines = File.ReadAllLines("input/data.csv");
var jsonOutput = new StringBuilder();
var rows = new List<string>();

// Skip header (index 0)
for (var i = 1; i < lines.Length; i++)
{
    var line = lines[i].Trim();
    if (!string.IsNullOrEmpty(line))
    {
        try
        {
            // Parse CSV fields. 
            // Assuming format: Name,Age,Email,Score
            var parts = line.Split(',');
            if (parts.Length < 4) continue; // Skip malformed rows

            var name = parts[0].Trim();
            int age = 0;
            if (!int.TryParse(parts[1], out age)) age = 0;
            var email = parts[2].Trim();
            float score = 0.0f;
            if (!float.TryParse(parts[3], out score)) score = 0.0f;

            // Escape quotes for Name and Email to ensure valid JSON strings
            string jsonName = name.Replace("\"", "\\\"");
            string jsonEmail = email.Replace("\"", "\\\"");

            rows.Add($@"{{""Name"":""{jsonName}""}, ""Age"":{age}, ""Email"":""{jsonEmail}"", ""Score"":{score}}");
        }
        catch (Exception)
        {
            continue;
        }
    }
}

jsonOutput.Append("[\n  ");
for (var i = 0; i < rows.Count; i++)
{
    if (i > 0) jsonOutput.AppendLine(",");
    jsonOutput.AppendLine(rows[i]);
}
jsonOutput.AppendLine("]");

// Output to stdout
Console.WriteLine(jsonOutput.ToString());