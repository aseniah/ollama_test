using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the CSV file
string csvPath = "input/input.csv";
string expectedPath = "input/expected_format.json";

string csvContent = File.ReadAllText(csvPath);

// Parse CSV - we'll infer column names from first row or assume standard format
// Try to detect common patterns
var lines = csvContent.Trim().Split(new[] { "\r\n", "\n" }, StringSplitOptions.RemoveEmptyEntries);

// If file is empty, output empty array
if (lines.Length == 0)
{
    Console.WriteLine("[]");
    Environment.Exit(0);
}

// First line might be header or data
var headers = lines[0].Split(',').Select(h => h.Trim()).ToArray();

List<string> outputItems = new List<string>();

// Try to infer format: likely "name,age" or "name,year" or similar
// Based on age calculation task, we need to parse dates

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i].Trim();
    if (string.IsNullOrEmpty(line))
        continue;
    
    // Try to parse as CSV
    var parts = line.Split(',');
    
    // Common patterns:
    // 1. First column might be a date (YYYY-MM-DD format)
    // 2. Age is typically derived from a birthdate
    // 3. Need to calculate age as of July 1, 2025
    
    // Try to detect if we have date information in the first column
    if (DateTime.TryParse(parts[0].Trim(), out var startDate) && startDate.Kind != DateTimeKind.Unspecified)
    {
        // Assume this is birthdate, calculate age
        var referenceDate = new DateTime(2025, 7, 1);
        var birthDate = startDate;
        
        int age;
        if (birthDate.Month > referenceDate.Month || 
            (birthDate.Month == referenceDate.Month && birthDate.Day > referenceDate.Day))
        {
            age = referenceDate.Year - birthDate.Year - 1;
        }
        else
        {
            age = referenceDate.Year - birthDate.Year;
        }
        
        outputItems.Add($"{parts[1].Trim()}"); // Assuming second column is name or id
    }
    else if (DateTime.TryParse(parts[0].Trim(), out startDate))
    {
        // Try parsing different date formats
        var birthDate = startDate;
        var referenceDate = new DateTime(2025, 7, 1);
        
        int age;
        if (birthDate.Month > referenceDate.Month || 
            (birthDate.Month == referenceDate.Month && birthDate.Day > referenceDate.Day))
        {
            age = referenceDate.Year - birthDate.Year - 1;
        }
        else
        {
            age = referenceDate.Year - birthDate.Year;
        }
        
        outputItems.Add($"{parts[1].Trim()}");
    }
    else
    {
        // Assume simple name column, try to find date info elsewhere
        // Try to parse other fields as dates
        for (int j = 1; j < parts.Length; j++)
        {
            if (DateTime.TryParse(parts[j].Trim(), out var date))
            {
                var birthDate = date;
                var referenceDate = new DateTime(2025, 7, 1);
                
                int age;
                if (birthDate.Month > referenceDate.Month || 
                    (birthDate.Month == referenceDate.Month && birthDate.Day > referenceDate.Day))
                {
                    age = referenceDate.Year - birthDate.Year - 1;
                }
                else
                {
                    age = referenceDate.Year - birthDate.Year;
                }
                
                outputItems.Add($"{parts[0].Trim()}");
                break;
            }
        }
    }
}

// Sort by age or some other rule based on expected format
// For now, just create JSON array
var json = new JsonObject(new[] { "[]" });

// Create proper JSON array
JsonDocument jsonDoc = JsonDocument.Parse($"[{string.Join(",", outputItems.Select(i => "\"" + i + "\""))}]");
JsonDocument result = jsonDoc;

// Output the result as JSON array
Console.WriteLine(result.ValueKind == JsonValueKind.Array ? result.ToString() : "[[]]");