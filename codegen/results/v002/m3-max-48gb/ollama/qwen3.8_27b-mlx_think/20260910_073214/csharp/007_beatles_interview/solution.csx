using System;
using System.Collections.Generic;
using System.Text;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/input.csv");
var header = lines[0].Split(',');

var results = new JsonArray();

var referenceDate = new DateTime(2025, 7, 1);

for (int i = 1; i < lines.Length; i++)
{
    var fields = lines[i].Split(',');
    if (fields.Length < 8) continue;

    var fullName = fields[0].Trim();
    var birthdayStr = fields[1].Trim();
    var diedStr = fields[2].Trim();
    var fatherStr = fields[3].Trim();
    var motherStr = fields[4].Trim();
    var brotherStr = fields[5].Trim();
    var sisterStr = fields[6].Trim();

    // Split full name into First and Last
    var nameParts = fullName.Split(' ');
    var firstName = nameParts[0];
    var lastName = fullName.Substring(firstName.Length + 1);

    // Parse birthday M/D/YYYY
    var bdayParts = birthdayStr.Split('/');
    var bdayMonth = int.Parse(bdayParts[0]);
    var bdayDay = int.Parse(bdayParts[1]);
    var bdayYear = int.Parse(bdayParts[2]);
    var birthday = new DateTime(bdayYear, bdayMonth, bdayDay);

    // Format birthday as YYYY-MM-DD
    var birthdayFormatted = birthday.ToString("yyyy-MM-dd");

    // Calculate age
    int age;
    if (diedStr != "null" && diedStr != "")
    {
        var diedParts = diedStr.Split('/');
        var diedMonth = int.Parse(diedParts[0]);
        var diedDay = int.Parse(diedParts[1]);
        var diedYear = int.Parse(diedParts[2]);
        var diedDate = new DateTime(diedYear, diedMonth, diedDay);
        age = diedDate.Year - birthday.Year;
        if (diedDate < birthday.AddYears(age))
            age--;
    }
    else
    {
        age = referenceDate.Year - birthday.Year;
        if (referenceDate < birthday.AddYears(age))
            age--;
    }

    // Build relatives
    var relatives = new JsonArray();

    void AddRelative(string name, string relationship)
    {
        if (name == "null" || name == "") return;
        var parts = name.Split(' ');
        var rel = new JsonObject
        {
            ["FirstName"] = parts[0],
            ["LastName"] = name.Substring(parts[0].Length + 1),
            ["Relationship"] = relationship
        };
        relatives.Add(rel);
    }

    AddRelative(fatherStr, "Father");
    AddRelative(motherStr, "Mother");
    AddRelative(brotherStr, "Brother");
    AddRelative(sisterStr, "Sister");

    var person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayFormatted,
        ["Age"] = age,
        ["Relatives"] = relatives
    };

    results.Add(person);
}

var options = new JsonSerializerOptions { WriteIndented = true };
var output = results.ToJsonString(options);
Console.Write(output);