using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var refDate = new DateOnly(2025, 7, 1);
var lines = File.ReadAllLines("input/input.csv");
var headers = lines[0].Split(',');

int ColIdx(string name) => Array.IndexOf(headers, name);

int nameIdx      = ColIdx("Name");
int birthdayIdx  = ColIdx("Birthday");
int diedIdx      = ColIdx("Died");
int fatherIdx    = ColIdx("Father");
int motherIdx    = ColIdx("Mother");
int brotherIdx   = ColIdx("Brother");
int sisterIdx    = ColIdx("Sister");

static (string First, string Last) SplitName(string fullName)
{
    var parts = fullName.Trim().Split(' ');
    var last = parts[^1];
    var first = string.Join(" ", parts[..^1]);
    return (first, last);
}

static int CalcAge(DateOnly birth, DateOnly asOf)
{
    int age = asOf.Year - birth.Year;
    if (asOf < birth.AddYears(age)) age--;
    return age;
}

static DateOnly ParseDate(string s)
{
    // M/d/yyyy
    var parts = s.Split('/');
    return new DateOnly(int.Parse(parts[2]), int.Parse(parts[0]), int.Parse(parts[1]));
}

var result = new JsonArray();

foreach (var line in lines[1..])
{
    if (string.IsNullOrWhiteSpace(line)) continue;
    var cols = line.Split(',');

    var (firstName, lastName) = SplitName(cols[nameIdx]);
    var birthday = ParseDate(cols[birthdayIdx]);
    var diedStr = cols[diedIdx].Trim();
    DateOnly? died = diedStr == "null" || string.IsNullOrEmpty(diedStr) ? null : ParseDate(diedStr);

    DateOnly ageAsOf = died.HasValue ? died.Value : refDate;
    int age = CalcAge(birthday, ageAsOf);

    var relatives = new JsonArray();

    void AddRelative(string relName, string relationship)
    {
        if (string.IsNullOrEmpty(relName) || relName.Trim() == "null") return;
        var (rFirst, rLast) = SplitName(relName);
        var rel = new JsonObject
        {
            ["FirstName"] = rFirst,
            ["LastName"] = rLast,
            ["Relationship"] = relationship
        };
        relatives.Add(rel);
    }

    AddRelative(cols[fatherIdx].Trim(), "Father");
    AddRelative(cols[motherIdx].Trim(), "Mother");
    AddRelative(cols[brotherIdx].Trim(), "Brother");
    AddRelative(cols[sisterIdx].Trim(), "Sister");

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

var opts = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(result.ToJsonString(opts));
