using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Globalization;

// Configuration
string inputPath = "input/input.csv";
DateTime referenceDate = new DateTime(2025, 7, 1);

if (!File.Exists(inputPath))
{
    return;
}

string[] lines = File.ReadAllLines(inputPath);
if (lines.Length == 0) 
{
    Console.WriteLine("[]");
    return;
}

// Extract header to find column indices
string[] headers = lines[0].Split(',');
var dataLines = lines.Skip(1);

var resultList = new List<object>();

foreach (var line in dataLines)
{
    if (string.IsNullOrWhiteSpace(line)) continue;
    
    string[] cols = line.Split(',');
    
    // Map columns based on header names
    string fullName = cols[0];
    string birthdayStr = cols[1];
    string diedStr = cols[2];
    string father = cols[3];
    string mother = cols[4];
    string brother = cols[5];
    string sister = cols[6];

    // Split Name into FirstName and LastName
    // Based on input: "John Winston Lennon" -> "John", "Lennon"
    string[] nameParts = fullName.Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];

    // Parse Birthday (Format: M/d/yyyy)
    DateTime birthday = DateTime.ParseExact(birthdayStr, "M/d/yyyy", CultureInfo.InvariantCulture);
    
    // Calculate Age as of July 1, 2025
    // If they died, age is calculated as of death date. Otherwise, reference date.
    DateTime deathDate = DateTime.MaxValue;
    if (!string.IsNullOrEmpty(diedStr) && diedStr != "null")
    {
        deathDate = DateTime.ParseExact(diedStr, "M/d/yyyy", CultureInfo.InvariantCulture);
    }
    
    DateTime endPoint = deathDate < referenceDate ? deathDate : referenceDate;
    
    int age = endPoint.Year - birthday.Year;
    if (endPoint < birthday.AddYears(age)) age--;

    // Process Relatives
    var relatives = new List<object>();
    string[] familyMembers = { father, mother, brother, sister };
    string[] relationships = { "Father", "Mother", "Brother", "Sister" };

    for (int i = 0; i < familyMembers.Length; i++)
    {
        string member = familyMembers[i];
        if (!string.IsNullOrEmpty(member) && member != "null")
        {
            string[] relNameParts = member.Split(' ');
            relNameParts = relNameParts.Where(p => !string.IsNullOrWhiteSpace(p)).ToArray();
            
            relyObj = new
            {
                FirstName = relNameParts[0],
                LastName = relNameParts.Length > 1 ? relNameParts[relNameParts.Length - 1] : "",
                Relationship = relationships[i]
            };
            relatives.Add(relyObj);
        }
    }

    resultList.Add(new
    {
        FirstName = firstName,
        LastName = lastName,
        Birthday = birthday.ToString("yyyy-MM-dd"),
        Age = age,
        Relatives = relatives
    });
}

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(resultList, options));