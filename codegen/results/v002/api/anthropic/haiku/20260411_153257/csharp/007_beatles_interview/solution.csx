using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var referenceDate = new DateTime(2025, 7, 1);

int CalcAge(DateTime birthday, DateTime? died)
{
    var endDate = died.HasValue && died.Value < referenceDate ? died.Value : referenceDate;
    int age = endDate.Year - birthday.Year;
    if (birthday > endDate.AddYears(-age)) age--;
    return age;
}

var lines = File.ReadAllLines("input/input.csv");
var headers = lines[0].Split(',');

var result = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var cols = lines[i].Split(',');
    var fullName = cols[0].Trim();
    var nameParts = fullName.Split(' ');
    var firstName = nameParts[0];
    var lastName = nameParts[nameParts.Length - 1];

    var birthday = DateTime.Parse(cols[1].Trim());
    var diedStr = cols[2].Trim();
    DateTime? died = diedStr == "null" ? null : DateTime.Parse(diedStr);

    int age = CalcAge(birthday, died);

    var relatives = new JsonArray();
    string[] relNames = { cols[3].Trim(), cols[4].Trim(), cols[5].Trim(), cols[6].Trim() };
    string[] relTypes = { "Father", "Mother", "Brother", "Sister" };

    for (int r = 0; r < relNames.Length; r++)
    {
        if (relNames[r] == "null") continue;
        var parts = relNames[r].Split(' ');
        var relFirst = parts[0];
        var relLast = parts[parts.Length - 1];
        relatives.Add(new JsonObject
        {
            ["FirstName"] = relFirst,
            ["LastName"] = relLast,
            ["Relationship"] = relTypes[r]
        });
    }

    result.Add(new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthday.ToString("yyyy-MM-dd"),
        ["Age"] = age,
        ["Relatives"] = relatives
    });
}

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(result.ToJsonString(options));
