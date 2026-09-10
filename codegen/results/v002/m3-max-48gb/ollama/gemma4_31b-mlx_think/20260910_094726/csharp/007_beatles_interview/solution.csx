using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;

// Define the reference date for age calculation
DateTime targetDate = new DateTime(2025, 7, 1);

// Read CSV file
string[] lines = File.ReadAllLines("input/input.csv");
if (lines.Length == 0)
{
    Console.WriteLine("[]");
    return;
}

// Extract header and data
string[] header = lines[0].Split(',');
var dataLines = lines.Skip(1);

var result = new List<object>();

foreach (var line in dataLines)
{
    if (string.IsNullOrWhiteSpace(line)) continue;

    string[] cols = line.Split(',');
    
    // Map columns to values based on header
    var row = new Dictionary<string, string>();
    for (int i = 0; i < header.Length; i++)
    {
        row[header[i]] = cols[i].Trim();
    }

    // Parse Main Person's Name
    string fullName = row["Name"];
    string[] nameParts = fullName.Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];

    // Parse Birthday
    DateTime birthday = DateTime.ParseExact(row["Birthday"], "M/d/yyyy", CultureInfo.InvariantCulture);
    
    // Determine the reference date for age (Died date or Target Date)
    DateTime refDate = targetDate;
    if (!string.Equals(row["Died"], "null", StringComparison.OrdinalIgnoreCase) && !string.IsNullOrEmpty(row["Died"]))
    {
        refDate = DateTime.ParseExact(row["Died"], "M/d/yyyy", CultureInfo.InvariantCulture);
    }

    // Calculate Age
    int age = refDate.Year - birthday.Year;
    if (refDate < birthday.AddYears(age))
    {
        age--;
    }

    // Parse Relatives
    var relatives = new List<object>();
    string[] relationshipTypes = { "Father", "Mother", "Brother", "Sister" };
    
    foreach (var type in relationshipTypes)
    {
        string relativeName = row[type];
        if (!string.Equals(relativeName, "null", StringComparison.OrdinalIgnoreCase) && !string.IsNullOrEmpty(relativeName))
        {
            string[] relParts = relativeName.Split(' ');
            relatives.Add(new
            {
                FirstName = relParts[0],
                LastName = relParts[relParts.Length - 1],
                Relationship = type
            });
        }
    }

    // Add person to final list
    result.Add(new
    {
        FirstName = firstName,
        LastName = lastName,
        Birthday = birthday.ToString("yyyy-MM-dd"),
        Age = age,
        Relatives = relatives
    });
}

// Output JSON
var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(result, options));