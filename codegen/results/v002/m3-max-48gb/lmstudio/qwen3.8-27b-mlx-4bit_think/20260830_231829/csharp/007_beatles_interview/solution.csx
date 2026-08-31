using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/input.csv");
var header = lines[0].Split(',');
var rows = lines.Skip(1).Select(l => l.Split(',')).ToList();

var referenceDate = new DateTime(2025, 7, 1);

var people = new List<JsonObject>();

foreach (var row in rows)
{
    var name = row[0];
    var nameParts = name.Split(' ');
    var firstName = nameParts[0];
    var lastName = nameParts[nameParts.Length - 1];

    // Parse birthday
    var birthdayStr = row[1];
    var birthdayParts = birthdayStr.Split('/');
    var birthMonth = int.Parse(birthdayParts[0]);
    var birthDay = int.Parse(birthdayParts[1]);
    var birthYear = int.Parse(birthdayParts[2]);
    var birthday = new DateTime(birthYear, birthMonth, birthDay);

    // Parse died date
    var diedStr = row[2].Trim();
    DateTime? died = null;
    if (diedStr != "null" && diedStr.Length > 0)
    {
        var diedParts = diedStr.Split('/');
        var diedMonth = int.Parse(diedParts[0]);
        var diedDay = int.Parse(diedParts[1]);
        var diedYear = int.Parse(diedParts[2]);
        died = new DateTime(diedYear, diedMonth, diedDay);
    }

    // Calculate age
    int age;
    if (died != null)
    {
        var deathDate = died.Value;
        age = deathDate.Year - birthday.Year;
        if (deathDate.Month < birthday.Month || (deathDate.Month == birthday.Month && deathDate.Day < birthday.Day))
            age--;
    }
    else
    {
        age = referenceDate.Year - birthday.Year;
        if (referenceDate.Month < birthday.Month || (referenceDate.Month == birthday.Month && referenceDate.Day < birthday.Day))
            age--;
    }

    var birthdayFormatted = $"{birthYear:0000}-{birthMonth:00}-{birthDay:00}";

    // Build relatives
    var relatives = new JsonArray();
    var relCols = new[] { "Father", "Mother", "Brother", "Sister" };
    var relValues = new[] { row[3], row[4], row[5], row[6] };

    for (int i = 0; i < relCols.Length; i++)
    {
        var val = relValues[i].Trim();
        if (val != "null" && val.Length > 0)
        {
            var relParts = val.Split(' ');
            var relFirst = relParts[0];
            var relLast = relParts[relParts.Length - 1];
            relatives.Add(new JsonObject
            {
                ["FirstName"] = relFirst,
                ["LastName"] = relLast,
                ["Relationship"] = relCols[i]
            });
        }
    }

    var person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayFormatted,
        ["Age"] = age,
        ["Relatives"] = relatives
    };

    people.Add(person);
}

var array = new JsonArray();
foreach (var p in people)
    array.Add(p);

var options = new JsonSerializerOptions { WriteIndented = true };
Console.Write(JsonSerializer.Serialize(array, options));