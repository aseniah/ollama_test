using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Globalization;

// Define the path to the input CSV file
string csvPath = "input/input.csv";

// Check if the input file exists; if not, exit silently as per requirements
if (!File.Exists(csvPath))
{
    return;
    // Note: The prompt implies the file exists for the logic to run.
}

// The target date for age calculation: July 1, 2025
DateTime targetDate = new DateTime(2025, 7, 1);

// Read all lines from the CSV file
string[] lines = File.ReadAllLines(csvPath);

// If the file is empty or only contains a header, return an empty array
if (lines.Length < 2)
{
    if (lines.Length == 1)
    {
        Console.Write("[]");
    }
    return;
}

// The first line is the header
string[] headers = lines[0].Split(',').Select(h => h.Trim()).ToArray();
JsonArray jsonArray = new JsonArray();

// Iterate through each data row
for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    // Split the CSV line by comma and trim whitespace
    string[] values = line.Split(',').Select(v => v.Trim()).ToArray();
    
    // Map the current row to a dictionary for easier lookups
    var rowMap = new Dictionary<string, string>(StringComparer.OrdinalIgnoreCase);
    for (int j = 0; j < headers.Length; j++)
    {
        if (j < values.Length)
        {
            rowMap[headers[j]] = values[j];
        }
    }

    JsonObject jsonObject = new JsonObject();
    // Keep track of which columns from the original CSV have been "consumed" by transformations
    HashSet<string> consumedKeys = new HashSet<string>(StringComparer.OrdinalIgnoreCase);

    // --- Transformation 1: Name Construction ---
    // Check for direct 'name' or 'full_name' columns
    string fullName = null;
    if (rowMap.TryGetValue("name", out var n)) fullName = n;
    else if (rowMap.TryGetValue("fullname", out var fnm)) fullName = fnm;
    else if (rowMap.TryGetValue("full_name", out var fnm2)) fullName.Equals(fnm2); 
    else
    {
        // Attempt to construct name from First and Last name components
        string firstName = null;
        string lastName = null;

        string[] fnKeys = { "firstname", "first_name", "fname" };
        string[] lnKeys = { "lastname", "last_name", "lname" };

        foreach (var k in fnKeys) if (rowMap.TryGetValue(k, out var v)) { firstName = v; break; }
        foreach (var k in lnKeys) if (rowMap.TryGetValue(k, out var v)) { lastName = v; break; }

        if (firstName != null && lastName != null)
        {
            fullName = $"{firstName} {lastName}";
        }
    }

    if (fullName != null)
    {
        jsonObject["name"] = JsonValue.Create(fullName);
        // Mark the source keys as consumed so we don't duplicate them in the final JSON
        foreach (var key in rowMap.Keys)
        {
            if (key.Contains("name", StringComparison.OrdinalIgnoreCase) || 
                key.Contains("first", StringComparison.OrdinalIgnoreCase) || 
                key.Contains("last", StringComparison.OrdinalIgnoreCase) ||
                key.Equals("fname", StringComparison.OrdinalIgnoreCase) ||
                key.Equals("lname", StringComparison.OrdinalIgnoreCase))
            {
                consumedKeys.Add(key);
            }
        }
    }

    // --- Transformation 2: Age Calculation ---
    // Identify a date column (BirthDate, DOB, etc.)
    string dateKey = rowMap.Keys.FirstOrDefault(k =>
        k.Contains("birth", StringComparison.OrdinalIgnoreCase) ||
        k.Contains("dob", StringComparison.OrdinalIgnoreCase) ||
        k.Contains("date", StringComparison.OrdinalIgnoreCase));

    if (dateKey != null && DateTime.TryParse(rowMap[dateKey], out DateTime dob))
    {
        // Calculate age as of July 1, 2025
        int age = targetDate.Year - dob.Year;
        if (dob.Date > targetDate.AddYears(-age))
        {
            age--;
        }
        jsonObject["age"] = JsonValue.Create(age);
        consumedKeys.Add(dateKey);
    }

    // --- Transformation 3: Include Remaining Columns ---
    // Add any columns that weren't transformed, using lowercase keys
    foreach (var kvp in rowMap)
    {
        if (!consumedKeys.Contains(kvp.Key))
        {
            string key = kvp.Key.ToLowerInvariant();
            
            // Avoid creating duplicate 'name' or 'age' keys if they exist in raw format
            if (key != "name" && key != "age")
            {
                // Try to parse as a number for cleaner JSON output
                if (double.TryParse(kvp.Value, NumberStyles.Any, CultureInfo.InvariantCulture, out double num))
                {
                    jsonObject[key] = JsonValue.Create(num);
                }
                else
                {
                    jsonObject[key] = JsonValue.Create(kvp.Value);
                }
            }
        }
    }

    jsonArray.Add(jsonObject);
}

// Output the final JSON array to stdout
Console.Write(jsonArray.ToJsonString());