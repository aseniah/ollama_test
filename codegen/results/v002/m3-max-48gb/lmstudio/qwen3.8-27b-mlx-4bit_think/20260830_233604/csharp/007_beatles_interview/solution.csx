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
var rows = lines.Skip(1).Where(l => !string.IsNullOrWhiteSpace(l)).Select(l => l.Split(',')).ToList();

var referenceDate = new DateTime(2025, 7, 1);

var result = new JsonArray();

foreach (var row in rows)
{
    var nameParts = row[0].Trim().Split(' ', StringSplitOptions.RemoveEmptyEntries);
    var firstName = nameParts[0];
    var lastName = nameParts[nameParts.Length - 1];

    // Parse birthday: M/D/YYYY
    var bd = DateTime.ParseExact(row[1].Trim(), "M/d/yyyy", CultureInfo.InvariantCulture);
    var birthdayStr = bd.ToString("yyyy-MM-dd");

    // Calculate age
    int age;
    var diedStr = row[3].Trim();
    if (diedStr != "null" && !string.IsNullOrEmpty(diedStr))
    {
        var diedDate = DateTime.ParseExact(diedStr, "M/d/yyyy", CultureInfo.InvariantCulture);
        age = diedDate.Year - bd.Year;
        if (diedDate.Month < bd.Month || (diedDate.Month == bd.Month && diedDate.Day < bd.Day))
            age--;
    }
    else
    {
        age = referenceDate.Year - bd.Year;
        if (referenceDate.Month < bd.Month || (referenceDate.Month == bd.Month && referenceDate.Day < bd.Day))
            age--;
    }

    // Build relatives
    var relatives = new JsonArray();
    var relColumns = new[] { 3, 4, 5, 6 };
    var relNames = new[] { "Father", "Mother", "Brother", "Sister" };

    for (int i = 0; i < relColumns.Length; i++)
    {
        var val = row[relColumns[i]].Trim();
        if (val != "null" && !string.IsNullOrEmpty(val))
        {
            var relParts = val.Split(' ', StringSplitOptions.RemoveEmptyEntries);
            var rel = new JsonObject
            {
                ["FirstName"] = relParts[0],
                ["LastName"] = relParts[relParts.Length - 1],
                ["Relationship"] = relNames[i]
            };
            relatives.Add(rel);
        }
    }

    var person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayStr,
        ["Age"] = age,
        ["Relatives"] = relatives
    };

    result.Add(person);
}

var options = new JsonSerializerOptions { WriteIndented = true };
Console.Write(result.ToJsonString(options));