using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var referenceDate = new DateTime(2025, 7, 1);

var lines = File.ReadAllLines("input/input.csv");
var headers = lines[0].Split(',');

var result = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var fields = lines[i].Split(',');
    var fullName = fields[0].Trim();
    var birthdayRaw = fields[1].Trim();
    var diedRaw = fields[2].Trim();
    var father = fields[3].Trim();
    var mother = fields[4].Trim();
    var brother = fields[5].Trim();
    var sister = fields[6].Trim();

    var nameParts = fullName.Split(' ');
    var firstName = nameParts[0];
    var lastName = nameParts[nameParts.Length - 1];

    var birthday = DateTime.Parse(birthdayRaw);
    var birthdayFormatted = birthday.ToString("yyyy-MM-dd");

    DateTime ageDate = (diedRaw != "null" && !string.IsNullOrEmpty(diedRaw))
        ? DateTime.Parse(diedRaw)
        : referenceDate;

    int age = ageDate.Year - birthday.Year;
    if (birthday.Month > ageDate.Month || (birthday.Month == ageDate.Month && birthday.Day > ageDate.Day))
        age--;

    var relatives = new JsonArray();

    void AddRelative(string fullRelName, string relationship)
    {
        if (fullRelName == "null" || string.IsNullOrEmpty(fullRelName)) return;
        var parts = fullRelName.Split(' ');
        var relFirst = parts[0];
        var relLast = parts[parts.Length - 1];
        var rel = new JsonObject
        {
            ["FirstName"] = relFirst,
            ["LastName"] = relLast,
            ["Relationship"] = relationship
        };
        relatives.Add(rel);
    }

    AddRelative(father, "Father");
    AddRelative(mother, "Mother");
    AddRelative(brother, "Brother");
    AddRelative(sister, "Sister");

    var person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayFormatted,
        ["Age"] = age,
        ["Relatives"] = relatives
    };

    result.Add(person);
}

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(result.ToJsonString(options));
