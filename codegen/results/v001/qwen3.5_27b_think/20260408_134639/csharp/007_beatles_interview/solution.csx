#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read input files
string csvContent = File.ReadAllText("input/input.csv");
string expectedContent = File.ReadAllText("input/expected_format.json");

// Calculate age as of July 1, 2025
DateTime referenceDate = new DateTime(2025, 7, 1);

int CalculateAge(string dobString)
{
    if (string.IsNullOrWhiteSpace(dobString))
        return 0;
    
    DateTime dob = DateTime.Parse(dobString.Trim());
    int age = referenceDate.Year - dob.Year;
    
    if (referenceDate.Date < dob.Date.AddYears(age))
        age--;
    
    return age;
}

// Parse CSV (simple parser, handles basic CSV format)
string[] lines = csvContent.Split(new[] { "\r\n", "\n" }, StringSplitOptions.None);
List<string[]> rows = new List<string[]>();

for (int i = 0; i < lines.Length; i++)
{
    string line = lines[i].Trim();
    if (string.IsNullOrEmpty(line))
        continue;
    
    List<string> fields = new List<string>();
    string current = "";
    bool inQuotes = false;
    
    foreach (char c in line)
    {
        if (c == '"')
        {
            inQuotes = !inQuotes;
        }
        else if (c == ',' && !inQuotes)
        {
            fields.Add(current.Trim('"'));
            current = "";
        }
        else
        {
            current += c;
        }
    }
    fields.Add(current.Trim('"'));
    
    rows.Add(fields.ToArray());
}

if (rows.Count == 0)
{
    Console.WriteLine("[]");
    return;
}

// Identify columns by examining header
string[] header = rows[0];
Dictionary<string, int> headerMap = new Dictionary<string, int>();
for (int i = 0; i < header.Length; i++)
{
    string col = header[i].Trim().ToLower();
    headerMap[col] = i;
}

// Build JSON array
var output = new JsonArray();

for (int i = 1; i < rows.Count; i++)
{
    string[] fields = rows[i];
    var record = new JsonObject();
    
    // Find and extract relevant fields
    foreach (var key in headerMap.Keys)
    {
        int index = headerMap[key];
        if (index < fields.Length)
        {
            string value = fields[index].Trim();
            
            // Map column names to JSON property names
            string jsonKey = key;
            if (key.Contains("firstname") || key.Contains("first_name") || key == "fname")
                jsonKey = "firstName";
            else if (key.Contains("lastname") || key.Contains("last_name") || key == "lname")
                jsonKey = "lastName";
            else if (key.Contains("name"))
                jsonKey = "name";
            else if (key.Contains("dateofbirth") || key.Contains("birthdate") || key.Contains("dob") || key.Contains("birthday"))
                jsonKey = "age";
            else if (key.Contains("email"))
                jsonKey = "email";
            else if (key.Contains("city") || key.Contains("town"))
                jsonKey = "city";
            else if (key.Contains("country"))
                jsonKey = "country";
            
            // Calculate age if this is a date of birth field
            if (jsonKey == "age")
            {
                record[jsonKey] = CalculateAge(value);
            }
            else if (!string.IsNullOrEmpty(value))
            {
                record[jsonKey] = value;
            }
        }
    }
    
    output.Add(record);
}

// Output JSON array
Console.WriteLine(JsonSerializer.Serialize(output, new JsonSerializerOptions { WriteIndented = true }));