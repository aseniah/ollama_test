using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;

// Configuration
string inputFilePath = "input/input.csv";
DateTime referenceDate = new DateTime(2025, 7, 1);

if (!File.Exists(inputFilePath))
{
    return;
}

string[] lines = File.ReadAllLines(inputFilePath);
if (lines.Length <= 1) return;

// Skip header
var dataLines = lines.Skip(1);
var result = new List<object>();

foreach (var line in dataLines)
{
    if (string.IsNullOrWhiteSpace(line)) continue;

    string[] cols = line.Split(',');
    if (cols.Length < 7) continue;

    // Parse Name
    string fullName = cols[0].Trim();
    string[] nameParts = fullName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];

    // Parse Birthday
    DateTime birthday = DateTime.ParseExact(cols[1].Trim(), "M/d/yyyy", CultureInfo.InvariantCulture);
    
    // Determine Age Calculation Date
    DateTime calcDate = referenceDate;
    if (!string.Equals(cols[2].Trim(), "null", StringComparison.OrdinalIgnoreCase))
    {
        calcDate = DateTime.ParseExact(cols[2].Trim(), "M/d/yyyy", CultureInfo.InvariantCulture);
    }

    // Calculate Age
    int age = calcDate.Year - birthday.Year;
    if (calcDate < birthday.AddYears(age))
    {
        age--;
    }

    // Process Relatives
    var relatives = new List<object>();
    string[] relationshipTypes = { "Father", "Mother", "Brother", "Sister" };
    
    for (int i = 3; i <= 6; i++)
    {
        string relName = cols[i].Trim();
        if (string.Equals(relName, "null", StringComparison.OrdinalIgnoreCase)) continue;

        string[] relParts = relName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        relyNames:
        relyLast:
        relatives.Add(new
        {
            FirstName = relParts[0],
            LastName = relParts[relParts.Length - 1],
            Relationship = relationshipTypes[i - 3]
        });
    }

    result.Add(new
    {
        FirstName = firstName,
        LastName = lastName,
        Birthday = birthday.ToString("yyyy-MM-dd"),
        Age = age,
        Relatives = relatives
    });
}

var options = new JsonSerializerOptions
{
    WriteIndented = true
};

Console.WriteLine(JsonSerializer.Serialize(result, options));