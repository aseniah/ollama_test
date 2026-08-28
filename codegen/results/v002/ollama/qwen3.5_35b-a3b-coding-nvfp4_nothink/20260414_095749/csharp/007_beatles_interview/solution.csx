using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

var csvPath = "input/input.csv";
var lines = File.ReadAllLines(csvPath);

var referenceDate = new DateTime(2025, 7, 1);

var people = new List<JsonObject>();

// Skip header
for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i].Trim();
    if (string.IsNullOrWhiteSpace(line)) continue;

    var parts = line.Split(',');
    if (parts.Length < 8) continue;

    var fullName = parts[0];
    var birthdayStr = parts[1];
    var diedStr = parts[2];
    var fatherStr = parts[3];
    var motherStr = parts[4];
    var brotherStr = parts[5];
    var sisterStr = parts[6];

    // Parse Name
    var nameParts = fullName.Split(' ');
    var firstName = nameParts[0];
    var lastName = nameParts.Length > 1 ? string.Join(" ", nameParts, 1, nameParts.Length - 1) : "";

    // Parse Birthday
    var birthday = DateTime.Parse(birthdayStr);

    // Calculate Age as of July 1, 2025
    int age = referenceDate.Year - birthday.Year;
    if (referenceDate.Month < birthday.Month || (referenceDate.Month == birthday.Month && referenceDate.Day < birthday.Day))
    {
        age--;
    }

    var person = new JsonObject();
    person["FirstName"] = firstName;
    person["LastName"] = lastName;
    person["Birthday"] = birthday.ToString("yyyy-MM-dd");
    person["Age"] = age;

    var relatives = new JsonArray();

    // Father
    if (!string.IsNullOrWhiteSpace(fatherStr) && fatherStr != "null")
    {
        var fParts = fatherStr.Split(' ');
        var fFirstName = fParts[0];
        var fLastName = fParts.Length > 1 ? string.Join(" ", fParts, 1, fParts.Length - 1) : "";
        var relative = new JsonObject();
        relative["FirstName"] = fFirstName;
        relative["LastName"] = fLastName;
        relative["Relationship"] = "Father";
        relatives.Add(relative);
    }

    // Mother
    if (!string.IsNullOrWhiteSpace(motherStr) && motherStr != "null")
    {
        var mParts = motherStr.Split(' ');
        var mFirstName = mParts[0];
        var mLastName = mParts.Length > 1 ? string.Join(" ", mParts, 1, mParts.Length - 1) : "";
        var relative = new JsonObject();
        relative["FirstName"] = mFirstName;
        relative["LastName"] = mLastName;
        relative["Relationship"] = "Mother";
        relatives.Add(relative);
    }

    // Brother
    if (!string.IsNullOrWhiteSpace(brotherStr) && brotherStr != "null")
    {
        var bParts = brotherStr.Split(' ');
        var bFirstName = bParts[0];
        var bLastName = bParts.Length > 1 ? string.Join(" ", bParts, 1, bParts.Length - 1) : "";
        var relative = new JsonObject();
        relative["FirstName"] = bFirstName;
        relative["LastName"] = bLastName;
        relative["Relationship"] = "Brother";
        relatives.Add(relative);
    }

    // Sister
    if (!string.IsNullOrWhiteSpace(sisterStr) && sisterStr != "null")
    {
        var sParts = sisterStr.Split(' ');
        var sFirstName = sParts[0];
        var sLastName = sParts.Length > 1 ? string.Join(" ", sParts, 1, sParts.Length - 1) : "";
        var relative = new JsonObject();
        relative["FirstName"] = sFirstName;
        relative["LastName"] = sLastName;
        relative["Relationship"] = "Sister";
        relatives.Add(relative);
    }

    person["Relatives"] = relatives;
    people.Add(person);
}

Console.WriteLine(JsonSerializer.Serialize(people, new JsonSerializerOptions
{
    WriteIndented = true
}));