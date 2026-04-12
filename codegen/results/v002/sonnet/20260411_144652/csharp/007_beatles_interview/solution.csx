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
    var nameParts = fullName.Split(' ');
    var firstName = nameParts[0];
    var lastName = nameParts[nameParts.Length - 1];

    var birthdayRaw = fields[1].Trim();
    var birthday = DateTime.Parse(birthdayRaw);
    var birthdayFormatted = birthday.ToString("yyyy-MM-dd");

    var diedRaw = fields[2].Trim();
    DateTime? diedDate = diedRaw == "null" ? null : DateTime.Parse(diedRaw);

    var ageAsOf = diedDate.HasValue && diedDate.Value < referenceDate ? diedDate.Value : referenceDate;
    int age = ageAsOf.Year - birthday.Year;
    if (birthday.Month > ageAsOf.Month || (birthday.Month == ageAsOf.Month && birthday.Day > ageAsOf.Day))
        age--;

    var relatives = new JsonArray();
    var relationshipColumns = new[] { "Father", "Mother", "Brother", "Sister" };
    foreach (var rel in relationshipColumns)
    {
        int colIndex = Array.IndexOf(headers, rel);
        if (colIndex < 0 || colIndex >= fields.Length) continue;
        var relValue = fields[colIndex].Trim();
        if (relValue == "null" || string.IsNullOrEmpty(relValue)) continue;

        var relParts = relValue.Split(' ');
        var relFirst = relParts[0];
        var relLast = relParts[relParts.Length - 1];

        var relObj = new JsonObject
        {
            ["FirstName"] = relFirst,
            ["LastName"] = relLast,
            ["Relationship"] = rel
        };
        relatives.Add(relObj);
    }

    var personObj = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayFormatted,
        ["Age"] = age,
        ["Relatives"] = relatives
    };

    result.Add(personObj);
}

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(result.ToJsonString(options));
