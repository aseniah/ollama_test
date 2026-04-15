using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

var csvLines = File.ReadAllLines("input/input.csv");
var jsonLines = File.ReadAllLines("input/expected_format.json"); // Just to ensure file exists, not used for parsing

var referenceDate = new DateTime(2025, 7, 1);
var people = new List<JsonObject>();

for (int i = 1; i < csvLines.Length; i++)
{
    var line = csvLines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    var parts = line.Split(',');
    if (parts.Length < 7) continue;

    var fullName = parts[0];
    var birthdayStr = parts[1];
    var diedStr = parts[2]; // Not used for output but parsed implicitly
    var fatherStr = parts[3];
    var motherStr = parts[4];
    var brotherStr = parts[5];
    var sisterStr = parts[6];

    // Parse Name: "FirstName MiddleName LastName" -> "FirstName" "LastName"
    // Example: "John Winston Lennon" -> First: "John", Last: "Lennon"
    var nameParts = fullName.Split(' ');
    var firstName = nameParts[0];
    var lastName = nameParts[nameParts.Length - 1];

    // Parse Birthday: "M/D/YYYY" -> "YYYY-MM-DD"
    var birthdayParts = birthdayStr.Split('/');
    var year = int.Parse(birthdayParts[2]);
    var month = int.Parse(birthdayParts[0]);
    var day = int.Parse(birthdayParts[1]);
    var birthday = new DateTime(year, month, day);

    // Calculate Age as of July 1, 2025
    int age = referenceDate.Year - birthday.Year;
    if (new DateTime(referenceDate.Year, referenceDate.Month, referenceDate.Day) < birthday)
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
        relatives.Add(new JsonObject
        {
            ["FirstName"] = fParts[0],
            ["LastName"] = fParts[fParts.Length - 1],
            ["Relationship"] = "Father"
        });
    }

    // Mother
    if (!string.IsNullOrEmpty(motherStr) && motherStr != "null")
    {
        var mParts = motherStr.Split(' ');
        relatives.Add(new JsonObject
        {
            ["FirstName"] = mParts[0],
            ["LastName"] = mParts[mParts.Length - 1],
            ["Relationship"] = "Mother"
        });
    }

    // Brother
    if (!string.IsNullOrEmpty(brotherStr) && brotherStr != "null")
    {
        var bParts = brotherStr.Split(' ');
        relatives.Add(new JsonObject
        {
            ["FirstName"] = bParts[0],
            ["LastName"] = bParts[bParts.Length - 1],
            ["Relationship"] = "Brother"
        });
    }

    // Sister
    if (!string.IsNullOrEmpty(sisterStr) && sisterStr != "null")
    {
        var sParts = sisterStr.Split(' ');
        relatives.Add(new JsonObject
        {
            ["FirstName"] = sParts[0],
            ["LastName"] = sParts[sParts.Length - 1],
            ["Relationship"] = "Sister"
        });
    }

    person["Relatives"] = relatives;
    people.Add(person);
}

Console.WriteLine(JsonSerializer.Serialize(people, new JsonSerializerOptions { WriteIndented = true }));