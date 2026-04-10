using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var referenceDate = new DateOnly(2025, 7, 1);

int CalculateAge(DateOnly birthDate, DateOnly asOf)
{
    int age = asOf.Year - birthDate.Year;
    if (birthDate.Month > asOf.Month || (birthDate.Month == asOf.Month && birthDate.Day > asOf.Day))
        age--;
    return age;
}

var lines = File.ReadAllLines("input/input.csv");
var headers = lines[0].Split(',');

// Find column indices
int nameIdx = Array.IndexOf(headers, "Name");
int birthdayIdx = Array.IndexOf(headers, "Birthday");
int diedIdx = Array.IndexOf(headers, "Died");
int fatherIdx = Array.IndexOf(headers, "Father");
int motherIdx = Array.IndexOf(headers, "Mother");
int brotherIdx = Array.IndexOf(headers, "Brother");
int sisterIdx = Array.IndexOf(headers, "Sister");

var result = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    if (string.IsNullOrWhiteSpace(lines[i])) continue;

    var fields = lines[i].Split(',');

    string fullName = fields[nameIdx].Trim();
    var nameParts = fullName.Split(' ');
    string lastName = nameParts[nameParts.Length - 1];
    string firstName = string.Join(" ", nameParts[0..^1]);

    var birthday = DateOnly.Parse(fields[birthdayIdx].Trim());

    string diedStr = fields[diedIdx].Trim();
    int age;
    if (string.IsNullOrEmpty(diedStr) || diedStr == "null")
    {
        age = CalculateAge(birthday, referenceDate);
    }
    else
    {
        var diedDate = DateOnly.Parse(diedStr);
        age = CalculateAge(birthday, diedDate);
    }

    var relatives = new JsonArray();

    void AddRelative(string relField, string relationship)
    {
        string val = fields[relField].Trim();
        if (string.IsNullOrEmpty(val) || val == "null") return;

        var parts = val.Split(' ');
        string relLast = parts[parts.Length - 1];
        string relFirst = string.Join(" ", parts[0..^1]);

        var rel = new JsonObject
        {
            ["FirstName"] = relFirst,
            ["LastName"] = relLast,
            ["Relationship"] = relationship
        };
        relatives.Add(rel);
    }

    AddRelative(fatherIdx, "Father");
    AddRelative(motherIdx, "Mother");
    AddRelative(brotherIdx, "Brother");
    AddRelative(sisterIdx, "Sister");

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

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(result.ToJsonString(options));
