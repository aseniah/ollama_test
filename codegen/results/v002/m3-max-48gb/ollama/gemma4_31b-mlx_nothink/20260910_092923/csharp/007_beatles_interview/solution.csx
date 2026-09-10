using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;

// Target date for age calculation
DateTime targetDate = new DateTime(2025, 7, 1);

if (!File.Exists("input/input.csv"))
{
    return;
}

string[] lines = File.ReadAllLines("input/input.csv");
if (lines.Length == 0) return;

// Extract header to identify columns
string[] headers = lines[0].Split(',');
var dataRows = lines.Skip(1);

var resultList = new List<object>();

foreach (var row in dataRows)
{
    if (string.IsNullOrWhiteSpace(row)) continue;
    
    string[] cols = row.Split(',');
    
    // Map CSV columns to local variables
    string fullName = cols[0];
    string birthdayStr = cols[1];
    string diedStr = cols[2];
    string father = cols[3];
    string mother = cols[4];
    string brother = cols[5];
    string sister = cols[6];

    // Parse Names (taking first and last)
    string[] nameParts = fullName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    string firstName = nameParts[0];
    string lastName = nameParts.Length > 1 ? nameParts[^1] : "";

    // Parse Birthday (Expected format M/D/YYYY)
    DateTime birthday = DateTime.ParseExact(birthdayStr, "M/d/yyyy", CultureInfo.InvariantCulture);
    
    // Age calculation as of July 1, 2025
    // If person died, age is calculated at date of death, otherwise at target date
    DateTime deathDate = DateTime.MinValue;
    if (!string.IsNullOrEmpty(diedStr) && diedStr != "null")
    {
        deathDate = DateTime.ParseExact(diedStr, "M/d/yyyy", CultureInfo.InvariantCulture);
    }

    DateTime calculationDate = (deathDate != DateTime.MinValue && deathDate < targetDate) 
        ? deathDate 
        : targetDate;

    int age = calculationDate.Year - birthday.Year;
    if (birthday.Date > calculationDate.AddYears(-age)) age--;

    // Process Relatives
    var relatives = new List<object>();
    string[] familyCols = { father, mother, brother, sister };
    string[] relationships = { "Father", "Mother", "Brother", "Sister" };

    for (int i = 0; i < familyCols.Length; i++)
    {
        string relName = familyCols[i];
        if (!string.IsNullOrEmpty(relName) && relName != "null")
        {
            string[] relParts = relName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
            relyName = relParts[0];
            string relLast = relParts.Length > 1 ? relParts[^1] : "";
            
            relatives.Add(new 
            { 
                FirstName = relParts[0], 
                LastName = relLast, 
                Relationship = relationships[i] 
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