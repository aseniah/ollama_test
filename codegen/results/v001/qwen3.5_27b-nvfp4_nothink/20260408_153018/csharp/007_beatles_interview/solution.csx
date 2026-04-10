#r "System.Text.Json"
#r "System.Data" // Optional fallback if needed, but we will implement CSV parsing manually to be safe across versions

using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Configuration
var inputPath = "input/input.csv";
var expectedPath = "input/expected_format.json";
var targetDate = new DateTime(2025, 7, 1);

// Read the expected output to infer structure (optional safety check)
string? expectedContent = null;
if (File.Exists(expectedPath))
{
    expectedContent = File.ReadAllText(expectedPath);
}

// Parse CSV manually to avoid dependency on specific Data sets for header parsing
using var stream = new FileStream(inputPath, FileMode.Open, FileAccess.Read);
using var reader = new StreamReader(stream);

var lines = new List<string>();
while (reader.EndOfStream == false)
{
    var line = reader.ReadLine();
    if (line != null && !string.IsNullOrWhiteSpace(line))
    {
        lines.Add(line);
    }
}

if (lines.Count < 2)
{
    Console.WriteLine("[]");
    return;
}

var headerLine = lines[0];
// Robust CSV parsing to handle quotes and commas within values
var headers = new List<string>();
var buffer = new System.Text.StringBuilder();
bool inQuotes = false;

foreach (char c in headerLine)
{
    if (c == '"')
    {
        inQuotes = !inQuotes;
    }
    else if (c == ',' && !inQuotes)
    {
        headers.Add(buffer.ToString().Trim().Trim('"'));
        buffer.Clear();
    }
    else
    {
        buffer.Append(c);
    }
}
headers.Add(buffer.ToString().Trim().Trim('"')); // Add last column

var records = new List<JsonNode>();
var referenceDateStr = "2025-07-01";

// Process data lines
for (int i = 1; i < lines.Count; i++)
{
    var line = lines[i];
    var values = new List<string>();
    buffer.Clear();
    inQuotes = false;

    foreach (char c in line)
    {
        if (c == '"')
        {
            inQuotes = !inQuotes;
        }
        else if (c == ',' && !inQuotes)
        {
            values.Add(buffer.ToString().Trim().Trim('"'));
            buffer.Clear();
        }
        else
        {
            buffer.Append(c);
        }
    }
    values.Add(buffer.ToString().Trim().Trim('"')); // Add last column

    var json = new JsonObject();

    for (int j = 0; j < headers.Count && j < values.Count; j++)
    {
        var key = headers[j];
        var val = values[j];
        
        // Handle Date and Age calculation logic here if the format requires it
        // We try to infer if a column is a DOB or Name based on content or expected format structure
        
        object? parsedVal = val;

        // Heuristic: If the key is "DOB", "Date of Birth", "Birth Date", parse date and calculate age
        // Or if the expected output shows an 'Age' field derived from a date column.
        // Since we read expected_format, we can try to match keys.
        
        // Basic type inference for generic CSV -> JSON
        if (string.IsNullOrEmpty(val))
        {
            parsedVal = null;
        }
        else if (int.TryParse(val, out var intVal))
        {
            parsedVal = intVal;
        }
        else if (long.TryParse(val, out var longVal))
        {
            parsedVal = longVal;
        }
        else if (double.TryParse(val, NumberStyles.Any, CultureInfo.InvariantCulture, out var doubleVal))
        {
            parsedVal = doubleVal;
        }
        else if (bool.TryParse(val, out var boolVal))
        {
            parsedVal = boolVal;
        }
        else if (DateTime.TryParse(val, out var dateVal))
        {
            // If we find a date column that might need transformation to Age
            // Check keys for "Age" in the expected output headers if available
            if (expectedContent != null)
            {
                try 
                {
                    var expectedRoot = JsonNode.Parse(expectedContent);
                    if (expectedRoot is JsonArray arr && arr.Count > 0)
                    {
                        var firstItem = arr[0];
                        // If the expected JSON has an "Age" property, and our current key looks like a DOB
                        if (firstItem != null && firstItem.AsObject().ContainsKey("Age"))
                        {
                            if (key.ToLowerInvariant().Contains("dob") || 
                                key.ToLowerInvariant().Contains("birth") ||
                                key.ToLowerInvariant().Contains("date"))
                            {
                                // Calculate Age as of July 1, 2025
                                var age = targetDate.Year - dateVal.Year;
                                if (targetDate.Month < dateVal.Month || 
                                    (targetDate.Month == dateVal.Month && targetDate.Day < dateVal.Day))
                                {
                                    age--;
                                }
                                
                                // If the output format expects "Age" as an integer, we set it
                                // But we need to map the column name. 
                                // Usually CSV has DOB, JSON expects Age. Or CSV has DOB and JSON keeps DOB + adds Age?
                                // The prompt says "produce a JSON array... use expected_format to infer transformation".
                                // If expected format has "Age", we should probably output "Age" instead of "DOB" or alongside it?
                                // Safest bet: If expected format has 'Age' and current is DOB, transform.
                                // However, if the CSV already has 'Age', keep it.
                                
                                // Let's look at the standard assumption: 
                                // Input has Date -> Output needs Age integer.
                                json.Add("Age", age);
                                continue; // Skip adding the raw date under its original name if we are transforming DOB->Age
                            }
                        }
                    }
                }
                catch { /* Ignore parsing errors of expected file */ }
            }
            
            // If no transformation detected, keep as ISO date string or raw value?
            // Standard JSON serialization of DateTime usually requires conversion to string.
            parsedVal = dateVal.ToString("yyyy-MM-dd", CultureInfo.InvariantCulture);
        }

        // Add to JSON object
        if (parsedVal != null)
        {
            json.Add(key, parsedVal is string s ? new JsonValue(s) : 
                     parsedVal is int n ? new JsonValue(n) : 
                     parsedVal is double d ? new JsonValue(d) : 
                     parsedVal is bool b ? new JsonValue(b) : 
                     new JsonValue(parsedVal.ToString()!));
        }
    }

    records.Add(json);
}

var options = new JsonSerializerOptions
{
    WriteIndented = true,
    Encoder = System.Text.Encodings.Web.JavaScriptEncoder.UnsafeRelaxedJsonEscaping
};

Console.WriteLine(JsonSerializer.Serialize(records, options));