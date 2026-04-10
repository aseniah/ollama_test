using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Globalization;

// Define the target date for age calculations
DateTime referenceDate = new DateTime(2025, 7, 1);

// File paths
string csvPath = "input/input.csv";
string jsonPath = "input/expected_format.json";

if (!File.Exists(csvPath) || !File.Exists(jsonPath))
{
    return;
}

// Read and parse CSV
// Using a simple split approach as per common requirements for this type of task
var csvLines = File.ReadAllLines(csvPath)
    .Where(line => !string.IsNullOrWhiteSpace(line))
    .Select(line => line.Split(',').Select(cell => cell.Trim()).ToArray())
    .ToList();

if (csvLines.Count == 0)
{
    Console.WriteLine("[]");
    return;
}

// The first row is the header
string[] headers = csvLines[0].Select(h => h.ToLowerInvariant()).ToArray();
// The remaining rows are the data
var dataRows = csvLines.Skip(1).ToList();

// Read the JSON template to infer the output structure and keys
string jsonTemplateText = File.ReadAllText(jsonPath);
JsonNode templateNode = JsonNode.Parse(jsonDTO(jsonTemplateText));
// If the template is an array, we look at the first object to find the structure
JsonObject templateObj = (templateNode is JsonArray arr && arr.Count > 0) 
    ? arr[0].AsObject() 
    : templateNode.AsObject();

var outputList = new List<JsonNode>();

foreach (var row in dataRows)
{
    var jsonRow = new JsonObject();

    foreach (var property in templateObj)
    {
        string key = property.Key;
        string keyLower = key.ToLowerInvariant();

        // RULE 1: Age Calculation
        // If the property name suggests an age, search for a column containing 'dob', 'birth', or 'date'
        if (keyLower.Contains("age"))
        {
            int dobIdx = Array.FindIndex(headers, h => h.Contains("dob") || h.Contains("birth") || h.Contains("date"));
            if (dobIdx != -1 && row.Length > dobIdx)
            {
                if (DateTime.TryParse(row[dobIdx], CultureInfo.InvariantCulture, DateTimeStyles.None, out DateTime dob))
                {
                    int age = referenceDate.Year - dob.Year;
                    // Adjust age if the birthday hasn't occurred yet in the reference year
                    if (dob.Date > referenceDate.AddYears(-age))
                    {
                        age--;
                    }
                    jsonRow[key] = age;
                    continue;
                }
            }
        }

        // RULE 2: Name Concatenation
        // If the key is 'name' or 'fullname', attempt to combine 'first' and 'last' columns
        if (keyLower.Contains("name") && !keyLower.Contains("first") && !keyLower.Contains("last"))
        {
            int firstIdx = Array.FindIndex(headers, h => h.Contains("first"));
            int lastIdx = Array.FindIndex(headers, h => h.Contains("last"));

            if (firstIdx != -1 && lastIdx != -1 && row.Length > Math.Max(firstIdx, lastIdx))
            {
                jsonRow[key] = $"{row[firstIdx]} {row[lastIdx]}".Trim();
                continue;
            }
        }

        // RULE 3: Direct Mapping / Fallback
        // Try to match the JSON key directly to a CSV header (case-insensitive, handling underscores)
        int matchIdx = Array.FindIndex(headers, h => 
            h == keyLower || 
            h.Replace("_", "") == keyLower || 
            h.EndsWith(keyLower)
        );

        if (matchIdx != -1 && row.Length > matchIdx)
        {
            string rawValue = row[matchIdx];

            // Try to preserve data types (bool, double, int) in the JSON output
            if (bool.TryParse(rawValue, out bool boolVal))
            {
                jsonRow[key] = boolVal;
            }
            else if (double.TryParse(rawValue, NumberStyles.Any, CultureInfo.InvariantCulture, out double doubleVal))
            {
                // Check if it's actually an integer to keep JSON clean
                if (doubleVal % 1 == 0)
                    jsonRow[key] = (long)doubleVal;
                else
                    jsonRow[key] = doubleVal;
            }
            else
            {
                jsonRow[key] = rawValue;
            }
        }
    }
    
    outputList.Add(jsonRow);
}

// Final Output to stdout
var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(outputList, options));

// Helper to handle potential array-wrapped templates
string dDTO(string json) => json;