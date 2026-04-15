using System;
using System.Collections.Generic;
using System.Globalization;
using System.Text.Json;
using System.Text.Json.Nodes;

// Reference date for age calculation
var referenceDate = new DateOnly(2025, 7, 1);

// Read input CSV
var lines = System.IO.File.ReadAllLines("input/input.csv");
var header = lines[0];
var dataLines = lines.AsEnumerable().Skip(1);

var result = new JsonArray();

foreach (var line in dataLines)
{
    if (string.IsNullOrWhiteSpace(line)) continue;
    
    var parts = line.Split(',');
    if (parts.Length < 8) continue;

    var fullName = parts[0].Trim();
    var nameParts = fullName.Split(' ');
    var firstName = nameParts[0];
    var lastName = nameParts.Length > 1 ? string.Join(" ", nameParts, 1, nameParts.Length - 1) : "";

    var birthdayStr = parts[1].Trim();
    var diedStr = parts[2].Trim();
    var fatherStr = parts[3].Trim();
    var motherStr = parts[4].Trim();
    var brotherStr = parts[5].Trim();
    var sisterStr = parts[6].Trim();

    // Parse birthday (DD/MM/YYYY format based on example)
    var birthday = DateOnly.ParseExact(birthdayStr, "d/M/yyyy", CultureInfo.InvariantCulture);
    
    // Calculate age
    int age = referenceDate.Year - birthday.Year;
    if (referenceDate < new DateOnly(referenceDate.Year, birthday.Month, birthday.Day))
    {
        age--;
    }

    var person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthday.ToString("yyyy-MM-dd"),
        ["Age"] = age
    };

    var relatives = new JsonArray();

    // Father
    if (!string.IsNullOrEmpty(fatherStr) && fatherStr != "null")
    {
        var fParts = fatherStr.Split(' ');
        var fFirst = fParts[0];
        var fLast = fParts.Length > 1 ? string.Join(" ", fParts, 1, fParts.Length - 1) : "";
        relatives.Add(new JsonObject
        {
            ["FirstName"] = fFirst,
            ["LastName"] = fLast,
            ["Relationship"] = "Father"
        });
    }

    // Mother
    if (!string.IsNullOrEmpty(motherStr) && motherStr != "null")
    {
        var mParts = motherStr.Split(' ');
        var mFirst = mParts[0];
        var mLast = mParts.Length > 1 ? string.Join(" ", mParts, 1, mParts.Length - 1) : "";
        relatives.Add(new JsonObject
        {
            ["FirstName"] = mFirst,
            ["LastName"] = mLast,
            ["Relationship"] = "Mother"
        });
    }

    // Brother
    if (!string.IsNullOrEmpty(brotherStr) && brotherStr != "null")
    {
        var bParts = brotherStr.Split(' ');
        var bFirst = bParts[0];
        var bLast = bParts.Length > 1 ? string.Join(" ", bParts, 1, bParts.Length - 1) : "";
        relatives.Add(new JsonObject
        {
            ["FirstName"] = bFirst,
            ["LastName"] = bLast,
            ["Relationship"] = "Brother"
        });
    }

    // Sister
    if (!string.IsNullOrEmpty(sisterStr) && sisterStr != "null")
    {
        var sParts = sisterStr.Split(' ');
        var sFirst = sParts[0];
        var sLast = sParts.Length > 1 ? string.Join(" ", sParts, 1, sParts.Length - 1) : "";
        relatives.Add(new JsonObject
        {
            ["FirstName"] = sFirst,
            ["LastName"] = sLast,
            ["Relationship"] = "Sister"
        });
    }

    person["Relatives"] = relatives;
    result.Add(person);
}

// Output JSON to stdout
Console.WriteLine(result.ToJsonString());