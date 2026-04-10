using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Globalization;
using System.Collections.Generic;

// Define file paths based on requirements
string csvPath = "input/input.csv";
string jsonPath = "input/expected_format.json";

// Reference date for age calculation
DateTime referenceDate = new DateTime(2025, 7, 1);

// Ensure files exist before proceeding
if (!File.Exists(csvPath) || !File.Exists(jsonPath))
{
    return;
}

// Read CSV content
string[] csvLines = File.ReadAllLines(csvPath);
if (csvLines.Length == 0)
{
    Console.Write("[]");
    return;
}

// Extract headers from the first row of CSV
string[] headers = csvLines[0].Split(',').Select(h => h.Trim()).ToArray();

// Read the expected format JSON to infer the target schema/keys
string jsonTemplate = File.ReadAllText(jsonPath);
var templateArray = JsonNode.Parse(jsonTemplate).AsArray();

if (templateArray.Count == 0)
{
    Console.Write("[]");
    return;
}

// Identify the properties required in the output by looking at the first object in the template
var targetProperties = templateArray[0].AsObject().Select(p => p.Key).ToList();
var jsonOutput = new JsonArray();

// Iterate through each data row in the CSV (skipping the header)
for (int i = 1; i < csvLines.Length; i++)
{
    string currentLine = csvLines[i];
    if (string.IsNullOrWhiteSpace(currentLine)) continue;

    // Split line by comma (basic CSV parsing)
    string[] values = currentLine.Split(',').Select(v => v.Trim()).ToArray();
    var rowObj = new JsonObject();

    foreach (var prop in targetProperties)
    {
        // Rule 1: If the JSON key is 'age', search for a date-related column in CSV
        if (prop.Equals("age", StringComparison.OrdinalIgnoreCase))
        {
            int dateColIdx = -1;
            for (int j = 0; j < headers.Length; j++)
            {
                string h = headers[j].ToLower();
                if (h.Contains("birth") || h.Contains("date") || h.Contains("dob"))
                {
                    dateColIdx = j;
                    break;
                }
            }

            if (dateColIdx != -1 && dateColIdx < values.Length)
            {
                if (DateTime.TryParse(values[dateColIdx], CultureInfo.InvariantCulture, DateTimeStyles.None, out DateTime dob))
                {
                    // Calculate age as of July 1, 2025
                    int age = referenceDate.Year - dob.Year;
                    if (referenceDate.Month < dob.Month || (referenceDate.Month == dob.Month && referenceDate.Day < dob.Day))
                    {
                        age--;
                    }
                    rowObj[prop] = age;
                }
            }
        }
        else
        {
            // Rule 2: Map other keys by finding a matching or containing header name
            int matchIdx = -1;
            for (int j = 0; j < headers.Length; j++)
            {
                string h = headers[j].ToLower();
                string p = prop.ToLower();
                // Match if header is exact, or header is a substring of property, or property is a substring of header
                if (h == p || h.Contains(p) || p.Contains(h))
                {
                    matchIdx = j;
                    break;
                }
            }

            if (matchIdx != -1 && matchIdx < values.Length)
            {
                rowObj[prop] = values[matchIdx];
            }
        }
    }
    jsonOutput.Add(rowObj);
}

// Output the final JSON array to stdout
Console.WriteLine(JsonSerializer.Serialize(jsonOutput, new JsonSerializerOptions { WriteIndented = true }));