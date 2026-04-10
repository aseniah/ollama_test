using System;
using System.Collections.Generic;
using System.IO;
using System.Globalization;
using System.Text.Json;

// Define the reference date for age calculation
DateTime referenceDate = new DateTime(2025, 7, 1);

string inputPath = "input/input.csv";
string expectedFormatPath = "input/expected_format.json";

if (!File.Exists(inputPath))
{
    // If input doesn't exist, we cannot proceed. 
    // However, the prompt implies the files exist for the task.
    return;
}

var results = new List<Dictionary<string, object>>();

// Read CSV
// Assuming standard CSV format: Header row exists.
// Since no specific CSV schema was provided, we infer from the context 
// that we need to map CSV columns to JSON keys.
string[] lines = File.ReadAllLines(inputPath);
if (lines.Length <= 1) return;

string[] headers = lines[0].Split(',');
// Clean headers (remove quotes/whitespace)
for (int i = 0; i < headers.Length; i++) headers[i] = headers[i].Trim('\"', ' ');

for (int i = 1; i < lines.Length; i++)
{
    if (string.IsNullOrWhiteSpace(lines[i])) continue;
    
    string[] values = lines[i].Split(',');
    var record = new Dictionary<string, object>();

    for (int j = 0; j < headers.Length; j++)
    {
        string key = headers[j];
        string val = j < values.Length ? values[j].Trim('\"', ' ') : "";

        // Logic for Age calculation based on requirements
        // We look for a column that looks like a date (e.g., "BirthDate", "dob", "date_of_birth")
        if (key.Contains("date", StringComparison.OrdinalIgnoreCase) || 
            key.Contains("birth", StringComparison.OrdinalIgnoreCase) || 
            key.Contains("dob", StringComparison.OrdinalIgnoreCase))
        {
            if (DateTime.TryParse(val, out DateTime birthDate))
            {
                int age = referenceDate.Year - birthDate.Year;
                if (birthDate.Date > referenceDate.AddYears(-age)) age--;
                
                // If the JSON expects the age as a specific key, we'd map it.
                // Since we don't see the JSON, we'll provide both the date and the calculated age.
                record[key] = val;
                record["age"] = age;
                continue;
            }
        }

        // Try to parse numbers
        if (double.TryParse(val, NumberStyles.Any, CultureInfo.InvariantCulture, out double num))
        {
            record[key] = num;
        }
        else
        {
            record[key] = val;
        }
    }
    results.Add(record);
}

// Output the JSON array to stdout
var options = new JsonSerializerOptions 
{ 
    WriteIndented = true 
};

string jsonOutput = JsonSerializer.Serialize(results, options);
Console.Write(jsonOutput);