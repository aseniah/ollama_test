using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Configuration
string inputFile = "input/input.csv";
string expectedFile = "input/expected_format.json";
DateTime referenceDate = new DateTime(2025, 7, 1);

// Read input CSV
string csvContent = File.ReadAllText(inputFile);
string[] lines = csvContent.Split(new[] { "\r\n", "\n" }, StringSplitOptions.RemoveEmptyEntries);
if (lines.Length == 0)
{
    Console.WriteLine("[]");
    return;
}

// Parse Headers
string[] headers = lines[0].Split(',');
int colFirstName = Array.IndexOf(headers, "first_name") ?? -1;
int colLastName = Array.IndexOf(headers, "last_name") ?? -1;
int colAge = Array.IndexOf(headers, "age") ?? -1; // Used to infer format if age is present
int colBirthDate = Array.IndexOf(headers, "birth_date") ?? -1;

// Determine Date Format by looking at an example line or expected output structure
string datePattern = null;
DateTime exampleDate = DateTime.MinValue;

// Try to infer date format from a sample row if birth_date exists
if (colBirthDate >= 0 && lines.Length > 1)
{
    string[] cols = lines[1].Split(',');
    if (cols.Length > colBirthDate)
    {
        // Attempt to parse common formats: dd/MM/yyyy, yyyy-MM-dd, MM/dd/yyyy
        foreach (var p in new[] { "dd/MM/yyyy", "yyyy-MM-dd", "MM/dd/yyyy" })
        {
            if (DateTime.TryParseExact(cols[colBirthDate].Trim(), p, null, System.Globalization.DateTimeStyles.None, out exampleDate))
            {
                datePattern = p;
                break;
            }
        }
    }
}

// Read Expected Output to infer structure keys (only if file exists)
Dictionary<string, string> outputKeys = new Dictionary<string, string>();
if (File.Exists(expectedFile))
{
    string expectedJson = File.ReadAllText(expectedFile);
    // Use JsonNode to parse the first object in the expected array to get keys
    try 
    {
        var root = JsonNode.Parse(expectedJson);
        if (root is JsonArray arr && arr.Count > 0)
        {
            if (arr[0] is JsonObject obj)
            {
                foreach (var kvp in obj.Properties())
                {
                    outputKeys.Add(kvp.Key, kvp.Value.ToString() ?? string.Empty);
                }
            }
        }
    }
    catch 
    {
        // Fallback if JSON parsing fails
    }
}

// Default mapping if expected file is missing or doesn't contain keys
outputKeys = new Dictionary<string, string>
{
    ["first_name"] = "",
    ["last_name"] = "",
    ["age"] = ""
};

var results = new List<JsonObject>();

for (int i = 1; i < lines.Length; i++)
{
    string[] cols = lines[i].Split(',');
    
    // Helper to get value safely
    string GetVal(int index) 
    { 
        if (index >= 0 && index < cols.Length) return cols[index].Trim(); 
        return ""; 
    }

    string fName = colFirstName >= 0 ? GetVal(colFirstName) : "";
    string lName = colLastName >= 0 ? GetVal(colLastName) : "";
    string birthDateStr = colBirthDate >= 0 ? GetVal(colBirthDate) : "";
    
    DateTime birthDate = DateTime.MinValue;

    // Parse Date
    if (!string.IsNullOrEmpty(birthDateStr))
    {
        // If we found a pattern earlier, use it. Otherwise try common ones again.
        string currentPattern = datePattern ?? "dd/MM/yyyy";
        
        // Retry parsing with inferred or default pattern
        if (DateTime.TryParseExact(birthDateStr, currentPattern, null, System.Globalization.DateTimeStyles.None, out birthDate))
        {
            // Success
        }
        else if (DateTime.TryParseExact(birthDateStr, "yyyy-MM-dd", null, System.Globalization.DateTimeStyles.None, out birthDate))
        {
             currentPattern = "yyyy-MM-dd";
        }
        else if (DateTime.TryParseExact(birthDateStr, "MM/dd/yyyy", null, System.Globalization.DateTimeStyles.None, out birthDate))
        {
             currentPattern = "MM/dd/yyyy";
        }
    }

    int age = 0;
    if (birthDate.Year > 1) 
    {
        int years = referenceDate.Year - birthDate.Year;
        // Adjust if birthday hasn't occurred yet in the reference year
        if (referenceDate.Month < birthDate.Month || 
            (referenceDate.Month == birthDate.Month && referenceDate.Day < birthDate.Day))
        {
            years--;
        }
        age = years;
    }

    // Construct Output Object based on inferred keys
    var jsonNode = new JsonObject();
    
    foreach (var kvp in outputKeys)
    {
        switch (kvp.Key.ToLower())
        {
            case "first_name":
                jsonNode.Add("firstName", fName);
                break;
            case "last_name":
                jsonNode.Add("lastName", lName);
                break;
            case "age":
                jsonNode.Add("age", age);
                break;
            // Add birth_date to output if it exists in expected format (inferred)
            case "birth_date": 
                // Only add if it was explicitly present in the expected structure logic, 
                // but usually age implies birth date calculation. 
                // If the key exists in the expected json object we parsed, we might need to output formatted date?
                // Based on prompt "Calculate ages", usually age is the calculated value.
                // We assume standard camelCase mapping unless specified otherwise by keys found.
                if (!jsonNode.ContainsKey("birthDate")) 
                {
                    // Optional: Add original date string if required by strict matching, 
                    // but usually "age" task implies just returning age. 
                    // Let's stick to standard naming conventions found in the key lookup.
                }
                break;
        }
    }
    
    // Ensure we have the standard fields derived from keys if not explicitly mapped above
    if (!jsonNode.ContainsKey("firstName")) jsonNode.Add("firstName", fName);
    if (!jsonNode.ContainsKey("lastName")) jsonNode.Add("lastName", lName);
    if (!jsonNode.ContainsKey("age")) jsonNode.Add("age", age);

    results.Add(jsonNode);
}

// Output JSON
var options = new JsonSerializerOptions
{
    WriteIndented = true,
    PropertyNamingPolicy = JsonNamingPolicy.CamelCase
};

Console.WriteLine(JsonSerializer.Serialize(results, options));