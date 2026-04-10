using System;
using System.IO;
using System.Collections.Generic;
using System.Globalization;
using System.Text.Json;
using System.Linq;

// Target date for age calculation
DateTime referenceDate = new DateTime(2025, 7, 1);

string inputPath = "input/input.csv";
string expectedFormatPath = "input/expected_format.json";

if (!File.Exists(inputPath))
{
    // If input doesn't exist, we can't proceed. 
    // Requirements state "Do not output anything other than the JSON array".
    return;
}

// Read the CSV file
var lines = File.ReadAllLines(inputPath);
if (lines.Length <= 1)
{
    Console.WriteLine("[]");
    return;
}

// Parse CSV headers
var headers = lines[0].Split(',').Select(h => h.Trim().Replace("\"", "")).ToList();
var dataRows = lines.Skip(1).Where(l => !string.IsNullOrWhiteSpace(l));

var results = new List<Dictionary<string, object>>();

foreach (var row in dataRows)
{
    var values = row.Split(',').Select(v => v.Trim().Replace("\"", "")).ToList();
    var record = new Dictionary<string, object>();

    for (int i = 0; i < headers.Count; i++)
    {
        if (i >= values.Count) break;

        string key = headers[i];
        string val = values[i];

        // Logic for transformation:
        // 1. Handle Date of Birth to calculate Age
        // 2. Convert numeric strings to numbers if possible
        // 3. Ensure standard casing/naming if inferred from JSON
        
        // Check if this column looks like a Date (e.g., 'dob', 'birthdate')
        if (key.ToLower().Contains("dob") || key.ToLower().Contains("birth"))
        {
            if (DateTime.TryParse(val, out DateTime dob))
            {
                int age = referenceDate.Year - dob.Year;
                if (dob.Date > referenceDate.AddYears(-age)) age--;
                record["age"] = age;
                // We include the original dob as a string if needed, 
                // but typically transformation implies replacing or adding.
                // Based on "infer rules", we provide the calculated age.
                continue; 
            }
        }

        // Try parsing numbers
        if (double.TryParse(val, NumberStyles.Any, CultureInfo.InvariantCulture, out double numVal))
        {
            // Check if it's an integer to avoid unnecessary decimals
            if (numVal % 1 == 0)
                record[key] = (int)numVal;
            else
                record[key] = numVal;
        }
        else
        {
            record[key] = val;
        }
    }
    
    results.Add(record);
}

// Note: Since I cannot see the content of 'expected_format.json', 
// the logic above implements the most common CSV-to-JSON transformation 
// patterns (type conversion and the specific age calculation requested).

var options = new JsonSerializerOptions
{
    WriteIndented = true
};

Console.WriteLine(JsonSerializer.Serialize(results, options));