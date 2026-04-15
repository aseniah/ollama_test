using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Globalization;

// Reference date for age calculation
DateTime referenceDate = new DateTime(2025, 7, 1);

if (!File.Exists("input/input.csv"))
{
    Environment.Exit(1);
}

string[] lines = File.ReadAllLines("input/input.csv");
if (lines.Length <= 1)
{
    Console.WriteLine("[]");
    return;
}

// Skip header
var dataLines = lines.Skip(1);
var resultList = new List<object>();

foreach (var line in dataLines)
{
    if (string.IsNullOrWhiteSpace(line)) continue;

    string[] cols = line.Split(',');
    if (cols.Length < 7) continue;

    // Split full name into First and Last
    string fullName = cols[0];
    string[] nameParts = fullName.Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts.Length > 1 ? nameParts[nameParts.Length - 1] : "";

    // Parse Birthday (Expected format MM/dd/yyyy based on CSV)
    DateTime birthday;
    DateTime.TryParseExact(cols[1], "M/d/yyyy", CultureInfo.InvariantCulture, DateTimeStyles.None, out birthday);
    string birthdayStr = birthday.ToString("yyyy-MM-dd");

    // Calculate Age
    // If Died date is present, age is at time of death. Otherwise, as of referenceDate.
    DateTime endReference = referenceDate;
    if (!string.IsNullOrEmpty(cols[2]) && cols[2] != "null")
    {
        if (DateTime.TryParseExact(cols[2], "M/d/yyyy", CultureInfo.InvariantCulture, DateTimeStyles.None, out DateTime deathDate))
        {
            endReference = deathDate;
        }
    }

    int age = endReference.Year - birthday.Year;
    if (endReference < birthday.AddYears(age)) age--;

    // Process Relatives
    var relatives = new List<object>();
    string[] relativeCols = { cols[3], cols[4], cols[5], cols[6] };
    string[] labels = { "Father", "Mother", "Brother", "Sister" };

    for (int i = 0; i < relativeCols.Length; i++)
    {
        string relName = relativeCols[i];
        if (string.IsNullOrEmpty(relName) || relName == "null") continue;

        string[] relParts = relName.Split(' ');
        reatives.Add(new
        {
            FirstName = relParts[0],
            LastName = relParts.Length > 1 ? relParts[relParts.Length - 1] : "",
            Relationship = labels[i]
        });
    }

    resultList.Add(new
    {
        FirstName = firstName,
        LastName = lastName,
        Birthday = birthdayStr,
        Age = age,
        Relatives = relatives
    });
}

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(resultList, options));