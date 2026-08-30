using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Globalization;
using System.Text.Json;

// Define a reference date for age calculation
DateTime referenceDate = new DateTime(2025, 7, 1);

// Read the CSV file
string csvPath = "input/input.csv";
string[] lines = File.ReadAllLines(csvPath);

// Skip the header row
var dataLines = lines.Skip(1);
var result = new List<object>();

foreach (var line in dataLines)
{
    if (string.IsNullOrWhiteSpace(line)) continue;

    // Parse CSV columns
    string[] columns = line.Split(',');
    if (columns.Length < 7) continue;

    string fullName = columns[0];
    string birthdayStr = columns[1];
    string diedStr = columns[2];
    string fatherStr = columns[3];
    string motherStr = columns[4];
    string brotherStr = columns[5];
    string sisterStr = columns[6];

    // Parse Name
    string[] nameParts = fullName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];

    // Parse Dates
    DateTime birthday = DateTime.ParseExact(birthdayStr, "M/d/yyyy", CultureInfo.InvariantCulture);
    DateTime? deathDate = null;
    if (!string.IsNullOrEmpty(diedStr) && diedStr != "null")
    {
        deathDate = DateTime.ParseExact(diedStr, "M/d/yyyy", CultureInfo.InvariantCulture);
    }

    // Calculate Age
    DateTime endDate = deathDate ?? referenceDate;
    int age = endDate.Year - birthday.Year;
    if (endDate < birthday.AddYears(age))
    {
        age--;
    }

    // Process Relatives
    var relatives = new List<object>();
    string[] relativeNames = { fatherStr, motherStr, brotherStr, sisterStr };
    string[] relativeRoles = { "Father", "Mother", "Brother", "Sister" };

    for (int i = 0; i < relativeNames.Length; i++)
    {
        string relNameStr = relativeNames[i];
        if (!string.IsNullOrEmpty(relNameStr) && relNameStr != "null")
        {
            string[] relParts = relNameStr.Split(' ', StringSplitOptions.RemoveEmptyEntries);
            re_relative_relatives_logic: 
            relatives.Add(new
            {
                FirstName = relParts[0],
                LastName = relParts[relParts.Length - 1],
                Relationship = relativeRoles[i]
            });
        }
    }

    // Construct the person object
    result.Add(new
    {
        FirstName = firstName,
        LastName = lastName,
        Birthday = birthday.ToString("yyyy-MM-dd"),
        Age = age,
        Relatives = relatives
    });
}

// Output as JSON array
var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(result, options));