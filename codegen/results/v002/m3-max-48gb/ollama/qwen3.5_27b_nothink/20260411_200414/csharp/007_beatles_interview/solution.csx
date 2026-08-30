using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var csvPath = "input/input.csv";
var lines = File.ReadAllLines(csvPath);

var result = new List<object>();
var cutoffDate = new DateTime(2025, 7, 1);

// Skip header
for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    // CSV parsing: Name,Birthday,Died,Father,Mother,Brother,Sister
    // The Name field might contain spaces, but no commas in this specific dataset.
    // However, to be robust against potential edge cases in a real CSV, we would usually use a parser.
    // Given the constraint "parse with string.Split(',')", we split by comma.
    // We assume the name does not contain a comma based on the input provided.
    var parts = line.Split(',');
    
    if (parts.Length < 7) continue;

    string fullName = parts[0].Trim();
    string birthdayRaw = parts[1].Trim();
    // parts[2] is Died (ignored for calculation, but present in CSV)
    string fatherName = parts[3].Trim();
    string motherName = parts[4].Trim();
    string brotherName = parts[5].Trim();
    string sisterName = parts[6].Trim();

    // Split full name into First and Last
    var nameParts = fullName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    string firstName = nameParts[0];
    string lastName = nameParts.Length > 1 ? nameParts[nameParts.Length - 1] : "";

    // Parse Birthday (Format: M/D/YYYY)
    // We parse the components manually to avoid culture issues with M/D/YYYY vs D/M/YYYY
    var dateParts = birthdayRaw.Split('/');
    if (dateParts.Length != 3) continue;
    
    int month = int.Parse(dateParts[0]);
    int day = int.Parse(dateParts[1]);
    int year = int.Parse(dateParts[2]);
    
    var birthDate = new DateTime(year, month, day);

    // Calculate Age as of July 1, 2025
    int age = cutoffDate.Year - birthDate.Year;
    if (cutoffDate.DayOfYear < birthDate.DayOfYear)
    {
        age--;
    }

    // Construct Relatives list
    var relatives = new List<object>();

    if (fatherName != "null")
    {
        var fParts = fatherName.Split(' ');
        relatives.Add(new {
            FirstName = fParts[0],
            LastName = fParts.Length > 1 ? fParts[fParts.Length - 1] : "",
            Relationship = "Father"
        });
    }

    if (motherName != "null")
    {
        var mParts = motherName.Split(' ');
        relatives.Add(new {
            FirstName = mParts[0],
            LastName = mParts.Length > 1 ? mParts[mParts.Length - 1] : "",
            Relationship = "Mother"
        });
    }

    if (brotherName != "null")
    {
        var bParts = brotherName.Split(' ');
        relatives.Add(new {
            FirstName = bParts[0],
            LastName = bParts.Length > 1 ? bParts[bParts.Length - 1] : "",
            Relationship = "Brother"
        });
    }

    if (sisterName != "null")
    {
        var sParts = sisterName.Split(' ');
        relatives.Add(new {
            FirstName = sParts[0],
            LastName = sParts.Length > 1 ? sParts[sParts.Length - 1] : "",
            Relationship = "Sister"
        });
    }

    // Create the record object matching the expected JSON structure
    var record = new {
        FirstName = firstName,
        LastName = lastName,
        Birthday = $"{year:0000}-{month:00}-{day:00}",
        Age = age,
        Relatives = relatives
    };

    result.Add(record);
}

// Output JSON
var options = new JsonSerializerOptions
{
    WriteIndented = true,
    PropertyNamingPolicy = JsonNamingPolicy.CamelCase,
};

Console.WriteLine(JsonSerializer.Serialize(result, options));