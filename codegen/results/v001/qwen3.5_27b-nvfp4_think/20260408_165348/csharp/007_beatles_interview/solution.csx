#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Globalization;

// Read CSV file
string csvPath = "input/input.csv";
string[] lines = File.ReadAllLines(csvPath);

if (lines.Length < 2)
{
    Console.WriteLine("[]");
    return;
}

// Parse header row
string[] headers = lines[0].Split(',');

// Reference date for age calculation
DateTime referenceDate = new DateTime(2025, 7, 1);

// Expected output keys to include in order
List<JsonNode> results = new List<JsonNode>();

for (int i = 1; i < lines.Length; i++)
{
    string[] values = lines[i].Split(',');
    
    // Create dictionary from header/value pairs
    Dictionary<string, string> row = new();
    for (int j = 0; j < headers.Length && j < values.Length; j++)
    {
        row[headers[j].Trim()] = values[j]?.Trim() ?? "";
    }
    
    var obj = new JsonObject();
    
    // Standard fields to include in output
    string nameValue = row.TryGetValue("name", out string n) ? n : 
                       (row.TryGetValue("Name", out string N) ? N : "");
    string dobValue = row.TryGetValue("dob", out string d) ? d : 
                      (row.TryGetValue("DOB", out string D) ? D : "") : "";
    string birthDateValue = row.TryGetValue("birth_date", out string b) ? b :
                            (row.TryGetValue("BirthDate", out string B) ? B : "") : "";
    
    // Add name field
    if (!string.IsNullOrEmpty(nameValue))
        obj["name"] = nameValue;
    
    // Find birth date value from any variation of the column name
    string birthDateString = dobValue;
    if (string.IsNullOrEmpty(birthDateString) && !string.IsNullOrEmpty(birthDateValue))
        birthDateString = birthDateValue;
    
    // Parse and calculate age if birth date exists
    DateTime? birthDate = null;
    if (!string.IsNullOrEmpty(birthDateString))
    {
        birthDate = DateTime.TryParseExact(birthDateString, "yyyy-MM-dd", 
            CultureInfo.InvariantCulture, DateTimeStyles.None, out var parsed)
            ? parsed :
            (DateTime.TryParse(birthDateString, out var parsed2) ? parsed2 : null);
    }
    
    // Calculate age as of July 1, 2025 if birth date was found
    if (birthDate.HasValue)
    {
        int age = referenceDate.Year - birthDate.Value.Year;
        if (referenceDate.Month < birthDate.Value.Month || 
            (referenceDate.Month == birthDate.Value.Month && 
             referenceDate.Day < birthDate.Value.Day))
        {
            age--;
        }
        obj["age"] = age;
    }
    
    // Add other fields that aren't handled specially
    foreach (var kvp in row)
    {
        string keyLower = kvp.Key.ToLower();
        if (!keyLower.Equals("name", StringComparison.OrdinalIgnoreCase) && 
            !keyLower.Contains("dob") && 
            !keyLower.Contains("birth"))
        {
            obj[kvp.Key] = kvp.Value;
        }
    }
    
    results.Add(obj);
}

// Output JSON array to stdout
var options = new JsonSerializerOptions { WriteIndented = false };
Console.WriteLine(JsonSerializer.Serialize(results, options));