#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Configuration
var inputFile = "input/input.csv";
var expectedFile = "input/expected_format.json";
var cutoffDate = new DateTime(2025, 7, 1);

// Helper: Parse CSV (basic implementation assuming no quoted commas)
string[] ReadLines(string path)
{
    if (!File.Exists(path)) return Array.Empty<string>();
    var lines = File.ReadAllLines(path).Skip(1).Where(l => !string.IsNullOrWhiteSpace(l)).ToArray();
    return lines;
}

// Helper: Parse a single CSV line
string[] SplitRow(string line)
{
    // Simple split by comma. If the data has complex quoting, this might need adjustment.
    return line.Split(',');
}

// Load and parse expected format to infer structure
JsonNode? expectedExample = null;
if (File.Exists(expectedFile))
{
    var jsonStr = File.ReadAllText(expectedFile);
    expectedExample = JsonNode.Parse(jsonStr);
}

var outputList = new List<JsonNode>();

// Read CSV lines
var csvLines = ReadLines(inputFile);

if (csvLines.Length == 0)
{
    Console.WriteLine("[]");
    return; // Exit
}

// Analyze header to guess column mapping based on expected JSON keys
var headers = SplitRow(csvLines[0]);
var colIndexByName = new Dictionary<string, int>(StringComparer.OrdinalIgnoreCase);
for (int i = 0; i < headers.Length; i++)
{
    var trimmed = headers[i].Trim();
    if (!colIndexByName.ContainsKey(trimmed))
    {
        colIndexByName[trimmed] = i;
    }
}

// Determine output keys from expected example if available
string[] targetKeys = { "id", "name", "age" }; // Default fallback
if (expectedExample != null)
{
    if (expectedExample is JsonArray arr && arr.Count > 0)
    {
        if (arr[0] is JsonObject firstObj)
        {
            targetKeys = firstObj.Keys.Where(k => !string.IsNullOrEmpty(k)).ToArray();
        }
    }
}

// Process rows
foreach (var line in csvLines.Skip(1))
{
    var values = SplitRow(line);
    var obj = new JsonObject();

    foreach (var key in targetKeys)
    {
        string mappedKey = key;
        
        // Heuristic mapping: if expected key matches a header name (case-insensitive), use it.
        // Common mappings: "Name" -> name, "Birthdate" -> age calculation, etc.
        
        if (key == "age")
        {
            // Look for date-like headers
            string? dateCol = null;
            foreach (var h in colIndexByName.Keys)
            {
                if (h.Contains("birth", StringComparison.OrdinalIgnoreCase) || 
                    h.Contains("dob", StringComparison.OrdinalIgnoreCase) || 
                    h.Contains("date", StringComparison.OrdinalIgnoreCase))
                {
                    dateCol = h;
                    break;
                }
            }

            if (!string.IsNullOrEmpty(dateCol) && colIndexByName.TryGetValue(dateCol, out int idx))
            {
                if (idx < values.Length)
                {
                    if (DateTime.TryParse(values[idx].Trim(), out var birthDate))
                    {
                        // Calculate age as of July 1, 2025
                        var age = cutoffDate.Year - birthDate.Year;
                        if (cutoffDate < birthDate.AddYears(age))
                        {
                            age--;
                        }
                        obj[key] = age;
                    }
                }
            }
        }
        else if (key == "id" || key == "name")
        {
            // Try exact header match or common aliases
            string? colName = key;
            
            // Fallback search if exact match fails
            if (!colIndexByName.ContainsKey(key))
            {
                foreach (var h in colIndexByName.Keys)
                {
                    if (key == "id" && (h.Contains("id", StringComparison.OrdinalIgnoreCase) || h.Contains("no", StringComparison.OrdinalIgnoreCase)))
                        colName = h;
                    if (key == "name" && h.Contains("name", StringComparison.OrdinalIgnoreCase))
                        colName = h;
                    
                    if (!string.IsNullOrEmpty(colName) && colIndexByName.ContainsKey(colName)) break;
                }
            }

            if (colIndexByName.TryGetValue(colName!, out int idx))
            {
                if (idx < values.Length)
                {
                    var val = values[idx].Trim();
                    // Special handling for ID: try to parse as int, else keep string
                    if (key == "id" && !string.IsNullOrEmpty(val) && int.TryParse(val, out int iVal))
                    {
                        obj[key] = iVal;
                    }
                    else
                    {
                        obj[key] = val; // Default to string
                    }
                }
            }
        }
    }

    outputList.Add(obj);
}

// Output JSON
var options = new JsonSerializerOptions
{
    WriteIndented = true,
    PropertyNameCaseInsensitive = false
};

Console.WriteLine(JsonSerializer.Serialize(outputList, options));