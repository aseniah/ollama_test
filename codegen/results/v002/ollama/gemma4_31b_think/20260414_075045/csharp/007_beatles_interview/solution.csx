using System;
using System.IO;
using System.Collections.Generic;
using System.Globalization;
using System.Text.Json;

// Reference date for age calculation
DateTime referenceDate = new DateTime(2025, 7, 1);

// Path to the input file
string inputFilePath = "input/input.csv";

if (!File.Exists(inputFilePath))
{
    return;
}

// Read all lines from the CSV file
string[] lines = File.ReadAllLines(inputFilePath);
if (lines.Length <= 1)
{
    Console.WriteLine("[]");
    return;
}

var results = new List<object>();

// Skip the header row and process each line
for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    string[] cols = line.Split(',');
    if (cols.Length < 7) continue;

    // Parsing Name
    string fullName = cols[0];
    string[] nameParts = fullName.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);
    string firstName = nameParts.Length > 0 ? nameParts[0] : "";
    string lastName = nameParts.Length > 1 ? nameParts[nameParts.Length - 1] : "";

    // Parsing Birthday
    DateTime birthday = DateTime.ParseExact(cols[1], "M/d/yyyy", CultureInfo.InvariantCulture);

    // Determining the date for age calculation: Date of death or July 1, 2025
    DateTime ageCalculationDate = referenceDate;
    if (cols[2] != "null" && !string.IsNullOrWhiteSpace(cols[2]))
    {
        ageCalculationDate = DateTime.ParseExact(cols[2], "M/d/yyyy", CultureInfo.InvariantCulture);
    }

    // Calculating Age
    int age = ageCalculationDate.Year - birthday.Year;
    if (ageCalculationDate.Month < birthday.Month || (ageCalculationDate.Month == birthday.Month && ageCalculationDate.Day < birthday.Day))
    {
        age--;
    }

    // Parsing Relatives
    var relatives = new List<object>();
    string[] relationshipTypes = { "Father", "Mother", "Brother", "Sister" };
    for (int j = 0; j < 4; j++)
    {
        string relativeName = cols[j + 3];
        if (relativeName != "null" && !string.IsNullOrWhiteSpace(relativeName))
        {
            string[] relParts = relativeName.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);
            string relFirst = relParts.Length > 0 ? relParts[0] : "";
            string relLast = relParts.Length > 1 ? relParts[relParts.Length - 1] : "";

            relatives.Add(new
            {
                FirstName = relFirst,
                LastName = relLast,
                Relationship = relationshipTypes[j]
            });
        }
    }

    // Create person object
    results.Add(new
    {
        FirstName = firstName,
        LastName = lastName,
        Birthday = birthday.ToString("yyyy-MM-dd"),
        Age = age,
        Relatives = relatives
    });
}

// Output the final JSON array to stdout
var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(results, options));