using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Globalization;

// Target date for age calculation
DateTime referenceDate = new DateTime(2025, 7, 1);

if (!File.Exists("input/input.csv"))
{
    return;
}

string[] lines = File.ReadAllLines("input/input.csv");
if (lines.Length <= 1) 
{
    Console.WriteLine("[]");
    return;
}

// Skip header
var dataLines = lines.Skip(1);
var results = new List<object>();

foreach (var line in dataLines)
{
    if (string.IsNullOrWhiteSpace(line)) continue;

    // Split CSV. Format: Name, Birthday, Died, Father, Mother, Brother, Sister
    string[] cols = line.Split(',');
    if (cols.Length < 7) continue;

    // Parse Name into First and Last
    string fullName = cols[0];
    string[] nameParts = fullName.Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts.Length > 1 ? nameParts[nameParts.Length - 1] : "";

    // Parse Birthday (Expected format M/d/yyyy based on input.csv)
    DateTime birthday;
    if (!DateTime.TryParseExact(cols[1], "M/d/yyyy", CultureInfo.InvariantCulture, DateTimeStyles.None, out birthday))
    {
        birthday = DateTime.MinValue;
    }

    // Calculate Age
    // If Died is present, age is calculated as of date of death. 
    // Otherwise, age is calculated as of July 1, 2025.
    DateTime ageEndDate = referenceDate;
    if (!string.IsNullOrEmpty(cols[2]) && cols[2] != "null")
    {
        if (DateTime.TryParseExact(cols[2], "M/d/yyyy", CultureInfo.InvariantCulture, DateTimeStyles.None, out DateTime deathDate))
        {
            ageEndDate = deathDate;
        }
    }

    int age = ageEndDate.Year - birthday.Year;
    if (birthday.Date > ageEndDate.AddYears(-age).Date) age--; // Adjust if birthday hasn't occurred yet in that year

    // Process Relatives
    var relatives = new List<object>();
    string[] relativeCols = { cols[3], cols[4], cols[5], cols[6] };
    string[] relationshipTypes = { "Father", "Mother", "Brother", "Sister" };

    for (int i = 0; i < relativeCols.Length; i++)
    {
        string relName = relativeCols[i];
        if (!string.IsNullOrEmpty(relName) && relName != "null")
        {
            string[] relParts = relName.Split(' ');
            relatives.Add(new
            {
                FirstName = relParts[0],
                LastName = relParts.Length > 1 ? relParts[relParts.Length - 1] : "",
                Relationship = relationshipTypes[i]
            });
        }
    }

    results.Add(new
    {
        FirstName = firstName,
        LastName = lastName,
        Birthday = birthday.ToString("yyyy-MM-dd"),
        Age = age,
        Relatives = relatives
    });
}

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(results, options));