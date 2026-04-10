#r "System.Text.Json"
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Define the calculation date
DateTime cutoffDate = new DateTime(2025, 7, 1);

// Read the input file
string inputPath = "input/input.csv";
if (!File.Exists(inputPath))
{
    Console.WriteLine("[]");
    return;
}

string[] lines = File.ReadAllLines(inputPath);
if (lines.Length == 0)
{
    Console.WriteLine("[]");
    return;
}

// Parse CSV headers
string[] headers = lines[0].Split(',');
int nameIdx = Array.IndexOf(headers, "Name");
int birthIdx = Array.IndexOf(headers, "Birth Date");
int emailIdx = Array.IndexOf(headers, "Email");

var result = new List<JsonObject>();

// Parse data rows
for (int i = 1; i < lines.Length; i++)
{
    string[] row = lines[i].Split(',');
    if (row.Length < 2) continue; // Skip empty lines

    string name = nameIdx >= 0 && nameIdx < row.Length ? row[nameIdx].Trim() : "";
    string birthStr = birthIdx >= 0 && birthIdx < row.Length ? row[birthIdx].Trim() : "";
    string email = emailIdx >= 0 && emailIdx < row.Length ? row[emailIdx].Trim() : "";

    DateTime birthDate;
    if (!DateTime.TryParse(birthStr, out birthDate))
    {
        continue; // Skip rows with invalid dates
    }

    // Calculate age as of July 1, 2025
    int age = cutoffDate.Year - birthDate.Year;
    if (cutoffDate < birthDate.AddYears(age))
    {
        age--;
    }

    // Create JSON object manually to match expected structure
    var person = new JsonObject();
    person["name"] = name;
    person["age"] = age;
    
    if (!string.IsNullOrEmpty(email))
    {
        person["email"] = email;
    }

    result.Add(person);
}

// Serialize to JSON
var options = new JsonSerializerOptions
{
    WriteIndented = true
};
string jsonOutput = JsonSerializer.Serialize(new JsonArray(result), options);

Console.WriteLine(jsonOutput);