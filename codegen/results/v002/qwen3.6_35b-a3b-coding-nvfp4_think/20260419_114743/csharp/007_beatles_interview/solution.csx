using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/input.csv").Skip(1).ToList();
var result = new JsonArray();
DateTime refDate = new DateTime(2025, 7, 1);

foreach (var line in lines)
{
    if (string.IsNullOrWhiteSpace(line)) continue;
    var cols = line.Split(',');
    string name = cols[0].Trim();
    string birthdayStr = cols[1].Trim();
    string father = cols[3].Trim();
    string mother = cols[4].Trim();
    string brother = cols[5].Trim();
    string sister = cols[6].Trim();

    var nameParts = name.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    string firstName = nameParts[0];
    string lastName = nameParts[^1];

    DateTime birthday = DateTime.ParseExact(birthdayStr, "M/d/yyyy", CultureInfo.InvariantCulture);
    int age = refDate.Year - birthday.Year;
    if (refDate < birthday.AddYears(age)) age--;

    var relatives = new JsonArray();
    if (father != "null")
    {
        var fParts = father.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        relatives.Add(new JsonObject { ["FirstName"] = fParts[0], ["LastName"] = fParts[^1], ["Relationship"] = "Father" });
    }
    if (mother != "null")
    {
        var mParts = mother.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        relatives.Add(new JsonObject { ["FirstName"] = mParts[0], ["LastName"] = mParts[^1], ["Relationship"] = "Mother" });
    }
    if (brother != "null")
    {
        var bParts = brother.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        relatives.Add(new JsonObject { ["FirstName"] = bParts[0], ["LastName"] = bParts[^1], ["Relationship"] = "Brother" });
    }
    if (sister != "null")
    {
        var sParts = sister.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        relatives.Add(new JsonObject { ["FirstName"] = sParts[0], ["LastName"] = sParts[^1], ["Relationship"] = "Sister" });
    }

    var person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthday.ToString("yyyy-MM-dd"),
        ["Age"] = age,
        ["Relatives"] = relatives
    };

    result.Add(person);
}

Console.WriteLine(result.ToJsonString(new JsonSerializerOptions { WriteIndented = true }));