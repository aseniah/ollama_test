using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Globalization;

// Constants for date calculation
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

// CSV Header: Name,Birthday,Died,Father,Mother,Brother,Sister
var records = lines.Skip(1).Select(line => line.Split(',')).ToList();
var resultList = new List<object>();

foreach (var cols in records)
{
    if (cols.Length < 7) continue;

    // Parse Name
    string fullName = cols[0];
    string firstName = fullName.Split(' ').First();
    string lastName = fullName.Split(' ').Last();

    // Parse Birthday
    DateTime birthday = DateTime.ParseExact(cols[1], "M/d/yyyy", CultureInfo.InvariantCulture);
    
    // Calculate Age as of July 1, 2025
    int age = referenceDate.Year - birthday.Year;
    if (birthday.Date > referenceDate.AddYears(-age)) age--;

    // Handle Relatives
    var relatives = new List<object>();
    string[] relationTypes = { "Father", "Mother", "Brother", "Sister" };
    
    for (int i = 3; i <= 6; i++)
    {
        string relValue = cols[i];
        if (!string.IsNullOrEmpty(relValue) && relValue != "null")
        {
            string[] relParts = relValue.Split(' ');
            relatives.Add(new
            {
                FirstName = relParts[0],
                LastName = relParts.Length > 1 ? relParts[1] : "",
                Relationship = relationTypes[i - 3]
            });
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